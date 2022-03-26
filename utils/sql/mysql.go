package dao

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx/reflectx"
	"reflect"
	"strings"
	"sync"
)

// NameMapper 用于将列名映射到结构字段名.
// 默认情况下使用 strings.ToLower 来进行结构体字段名小写化
var NameMapper = strings.ToLower
var origMapper = reflect.ValueOf(NameMapper)

var mpr *Mapper

// 互斥锁保护mpr.
var mprMu sync.Mutex

// 映射器使用配置的NameMapper返回有效的映射器
func mapper() *Mapper {
	mprMu.Lock()
	defer mprMu.Unlock()

	if mpr == nil {
		mpr = NewMapperFunc("db", NameMapper)
	} else if origMapper != reflect.ValueOf(NameMapper) {
		// if NameMapper has changed, create a new mapper
		mpr = NewMapperFunc("db", NameMapper)
		origMapper = reflect.ValueOf(NameMapper)
	}
	return mpr
}

// isScannable 通过reflect.Type和参数值返回这个参数能否被扫描
// 不能被扫描的参数:
//   * 它不是结构体
//   * 实现了sql.Scanner
//   * 没有公有字段
func isScannable(t reflect.Type) bool {
	if reflect.PtrTo(t).Implements(_scannerInterface) {
		return true
	}
	if t.Kind() != reflect.Struct {
		return true
	}

	// 扫描一个对象选择一个正确的映射器不重要
	// 重要的是关于这个结构体有多少导出字段
	return len(mapper().TypeMap(t).Index) == 0
}

// Queryer 被 Get 和 Select 使用的接口
type Queryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Queryx(query string, args ...interface{}) (*Rows, error)
	QueryRowx(query string, args ...interface{}) *Row
}

// Execer 被 MustExec 使用的接口
type Execer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var _scannerInterface = reflect.TypeOf((*sql.Scanner)(nil)).Elem()
var _valuerInterface = reflect.TypeOf((*driver.Valuer)(nil)).Elem()

// Row 是sql.Row的再实现，用来实现获取sql.Rows.Columns() 的数据, StructScan所需
type Row struct {
	err    error
	rows   *sql.Rows
	Mapper *Mapper
}

// Scan 是sql.Row.Scan的固定实现
func (r *Row) Scan(dest ...interface{}) error {
	if r.err != nil {
		return r.err
	}
	defer r.rows.Close()
	for _, dp := range dest {
		if _, ok := dp.(*sql.RawBytes); ok {
			return errors.New("sql: RawBytes isn't allowed on Row.Scan")
		}
	}

	if !r.rows.Next() {
		if err := r.rows.Err(); err != nil {
			return err
		}
		return sql.ErrNoRows
	}
	err := r.rows.Scan(dest...)
	if err != nil {
		return err
	}
	// 确保查询可以处理到完成, 没有错误
	if err := r.rows.Close(); err != nil {
		return err
	}
	return nil
}

// Columns 返回sql.Rows.Columns(), 或者是Row.Scan()返回的deferred error
func (r *Row) Columns() ([]string, error) {
	if r.err != nil {
		return []string{}, r.err
	}
	return r.rows.Columns()
}

// ColumnTypes 返回sql.Rows.ColumnTypes(), 或者是deferred error
func (r *Row) ColumnTypes() ([]*sql.ColumnType, error) {
	if r.err != nil {
		return []*sql.ColumnType{}, r.err
	}
	return r.rows.ColumnTypes()
}

// Err 返回扫描时遇到的错误
func (r *Row) Err() error {
	return r.err
}

// DB 是对sql.DB, 驱动名, 映射器的包装
type DB struct {
	*sql.DB
	driverName string
	Mapper     *Mapper
}

// NewDb 返回新的DB
func NewDb(db *sql.DB, driverName string) *DB {
	return &DB{DB: db, driverName: driverName, Mapper: mapper()}
}

// DriverName 返回DB的驱动名
func (db *DB) DriverName() string {
	return db.driverName
}

// Open 和sql.Open一样，但是是返回*dao.DB
func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{DB: db, driverName: driverName, Mapper: mapper()}, err
}

// MapperFunc 设置一个默认配置的映射器
func (db *DB) MapperFunc(mf func(string) string) {
	db.Mapper = NewMapperFunc("db", mf)
}

// Select 可以使所有占位符参数都将替换为提供的参数
func (db *DB) Select(dest interface{}, query string, args ...interface{}) error {
	return Select(db, dest, query, args...)
}

// Get 可以使所有占位符参数都将替换为提供的参数, 如果结果集为空，则返回错误。
func (db *DB) Get(dest interface{}, query string, args ...interface{}) error {
	return Get(db, dest, query, args...)
}

// Queryx 返回*dao.Rows, 使所有占位符参数都将替换为提供的参数
func (db *DB) Queryx(query string, args ...interface{}) (*Rows, error) {
	r, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return &Rows{Rows: r, Mapper: db.Mapper}, err
}

// QueryRowx 返回*dao.Row, 使所有占位符参数都将替换为提供的参数
func (db *DB) QueryRowx(query string, args ...interface{}) *Row {
	rows, err := db.DB.Query(query, args...)
	return &Row{rows: rows, err: err, Mapper: db.Mapper}
}

// MustExec (panic) 运行MustExec, 使所有占位符参数都将替换为提供的参数
func (db *DB) MustExec(query string, args ...interface{}) sql.Result {
	return MustExec(db, query, args...)
}

// Rows 是sql.Rows的包装, 用来缓存在扫描期间大量的反射操作
type Rows struct {
	*sql.Rows
	Mapper *Mapper
	// these fields cache memory use for a rows during iteration w/ structScan
	started bool
	fields  [][]int
	values  []interface{}
}

// ConnectDB 用来连接数据库, 连接成功时返回ping值
func ConnectDB(driverName, dataSourceName string) (*DB, error) {
	db, err := Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

// Select 使用提供的 Queryer 进行查询, 扫描每一行数据到参数里面,
// 如果切片元素是可以被扫描的, 则返回一行结果集
// 不然 *dao.Rows 将会自动关闭
func Select(q Queryer, dest interface{}, query string, args ...interface{}) error {
	rows, err := q.Queryx(query, args...)
	if err != nil {
		return err
	}
	// if something happens here, we want to make sure the rows are Closed
	defer rows.Close()
	return scanAll(rows, dest, false)
}

// Get 使用提供的 Queryer 进行查询,扫描每一行数据到参数里面,
// 如果切片元素是可以被扫描的, 则返回一行结果集
// 不然 Get 会返回 sql.ErrNoRows
// 如果结果集时空的话会返回错误
func Get(q Queryer, dest interface{}, query string, args ...interface{}) error {
	r := q.QueryRowx(query, args...)
	return r.scanAny(dest, false)
}

// MustExec 如果Exec出现错误会Panic
func MustExec(e Execer, query string, args ...interface{}) sql.Result {
	res, err := e.Exec(query, args...)
	if err != nil {
		panic(err)
	}
	return res
}

func asSliceForIn(i interface{}) (v reflect.Value, ok bool) {
	if i == nil {
		return reflect.Value{}, false
	}

	v = reflect.ValueOf(i)
	t := reflectx.Deref(v.Type())

	// Only expand slices
	if t.Kind() != reflect.Slice {
		return reflect.Value{}, false
	}

	// []byte is a driver.Value type so it should not be expanded
	if t == reflect.TypeOf([]byte{}) {
		return reflect.Value{}, false

	}

	return v, true
}

func appendReflectSlice(args []interface{}, v reflect.Value, vlen int) []interface{} {
	switch val := v.Interface().(type) {
	case []interface{}:
		args = append(args, val...)
	case []int:
		for i := range val {
			args = append(args, val[i])
		}
	case []string:
		for i := range val {
			args = append(args, val[i])
		}
	default:
		for si := 0; si < vlen; si++ {
			args = append(args, v.Index(si).Interface())
		}
	}

	return args
}

// In 批量操作时使用, 它可以扩展args的切片值,返回修改后的查询字符串和可以被数据库执行的新参数列表
func In(query string, args ...interface{}) (string, []interface{}, error) {
	// argMeta 存储reflect.Value, 切片的长度, 非切片参数的值
	type argMeta struct {
		v      reflect.Value
		i      interface{}
		length int
	}

	var flatArgsCount int
	var anySlices bool

	var stackMeta [32]argMeta

	var meta []argMeta
	if len(args) <= len(stackMeta) {
		meta = stackMeta[:len(args)]
	} else {
		meta = make([]argMeta, len(args))
	}

	for i, arg := range args {
		if a, ok := arg.(driver.Valuer); ok {
			var err error
			arg, err = a.Value()
			if err != nil {
				return "", nil, err
			}
		}

		if v, ok := asSliceForIn(arg); ok {
			meta[i].length = v.Len()
			meta[i].v = v

			anySlices = true
			flatArgsCount += meta[i].length

			if meta[i].length == 0 {
				return "", nil, errors.New("empty slice passed to 'in' query")
			}
		} else {
			meta[i].i = arg
			flatArgsCount++
		}
	}

	// 如果不是切片就不会解析
	// 不然如果之后出现一些错误, 那么这些错误将不会被返回
	if !anySlices {
		return query, args, nil
	}

	newArgs := make([]interface{}, 0, flatArgsCount)

	var buf strings.Builder
	buf.Grow(len(query) + len(", ?")*flatArgsCount)

	var arg, offset int

	for i := strings.IndexByte(query[offset:], '?'); i != -1; i = strings.IndexByte(query[offset:], '?') {
		if arg >= len(meta) {
			// if an argument wasn't passed, lets return an error;  this is
			// not actually how database/sql Exec/Query works, but since we are
			// creating an argument list programmatically, we want to be able
			// to catch these programmer errors earlier.
			return "", nil, errors.New("number of bindVars exceeds arguments")
		}

		argMeta := meta[arg]
		arg++

		// 不是切片, 继续
		// ? 要么在下一次扩展切片之前写入, 要么在编写其余查询时在循环之后写入
		if argMeta.length == 0 {
			offset = offset + i + 1
			newArgs = append(newArgs, argMeta.i)
			continue
		}

		// 写入所有数据包括 ?
		buf.WriteString(query[:offset+i+1])

		for si := 1; si < argMeta.length; si++ {
			buf.WriteString(", ?")
		}

		newArgs = appendReflectSlice(newArgs, argMeta.v, argMeta.length)

		// 切片查询并重置偏移量, 这避免了在循环后写入一些数据
		query = query[offset+i+1:]
		offset = 0
	}

	buf.WriteString(query)

	if arg < len(meta) {
		return "", nil, errors.New("number of bindVars less than number arguments")
	}

	return buf.String(), newArgs, nil
}

func (r *Row) scanAny(dest interface{}, structOnly bool) error {
	if r.err != nil {
		return r.err
	}
	if r.rows == nil {
		r.err = sql.ErrNoRows
		return r.err
	}
	defer r.rows.Close()

	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value, to StructScan destination")
	}
	if v.IsNil() {
		return errors.New("nil pointer passed to StructScan destination")
	}

	base := Deref(v.Type())
	scannable := isScannable(base)

	if structOnly && scannable {
		return structOnlyError(base)
	}

	columns, err := r.Columns()
	if err != nil {
		return err
	}

	if scannable && len(columns) > 1 {
		return fmt.Errorf("scannable dest type %s with >1 columns (%d) in result", base.Kind(), len(columns))
	}

	if scannable {
		return r.Scan(dest)
	}

	m := r.Mapper

	fields := m.TraversalsByName(v.Type(), columns)
	// 如果丢失字段
	if f, err := missingFields(fields); err != nil {
		return fmt.Errorf("missing destination name %s in %T", columns[f], dest)
	}
	values := make([]interface{}, len(columns))

	err = fieldsByTraversal(v, fields, values, true)
	if err != nil {
		return err
	}
	// 扫描结构体字段的指针, 将结果添加到结果集
	return r.Scan(values...)
}

// StructScan 对参数单独的 Row
func (r *Row) StructScan(dest interface{}) error {
	return r.scanAny(dest, true)
}

type rowsi interface {
	Close() error
	Columns() ([]string, error)
	Err() error
	Next() bool
	Scan(...interface{}) error
}

// structOnlyError 返回合适的错误类型
func structOnlyError(t reflect.Type) error {
	isStruct := t.Kind() == reflect.Struct
	isScanner := reflect.PtrTo(t).Implements(_scannerInterface)
	if !isStruct {
		return fmt.Errorf("expected %s but got %s", reflect.Struct, t.Kind())
	}
	if isScanner {
		return fmt.Errorf("structscan expects a struct dest but the provided struct type %s implements scanner", t.Name())
	}
	return fmt.Errorf("expected a struct, but struct %s has no exported fields", t.Name())
}

// scanAll 扫描所有行到 dest 中, 所以必须是任何类型的切片
// 如果 dest 切片是结构体, 那么 StructScan 会被每一行使用
// 如果 dest 是其他类型, 那么每一行必须有一列可以扫描到这个类型中去
func scanAll(rows rowsi, dest interface{}, structOnly bool) error {
	var v, vp reflect.Value

	value := reflect.ValueOf(dest)

	if value.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value, to StructScan destination")
	}
	if value.IsNil() {
		return errors.New("nil pointer passed to StructScan destination")
	}
	direct := reflect.Indirect(value)

	slice, err := baseType(value.Type(), reflect.Slice)
	if err != nil {
		return err
	}

	isPtr := slice.Elem().Kind() == reflect.Ptr
	base := Deref(slice.Elem())
	scannable := isScannable(base)

	if structOnly && scannable {
		return structOnlyError(base)
	}

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	// 如果这是基本数据类型, 需要保证它至少有一列数据
	// 不然会返回错误
	if scannable && len(columns) > 1 {
		return fmt.Errorf("non-struct dest type %s with >1 columns (%d)", base.Kind(), len(columns))
	}

	if !scannable {
		var values []interface{}
		var m *Mapper

		switch rows.(type) {
		case *Rows:
			m = rows.(*Rows).Mapper
		default:
			m = mapper()
		}

		fields := m.TraversalsByName(base, columns)
		// 如果丢失字段, 返回错误
		if f, err := missingFields(fields); err != nil {
			return fmt.Errorf("missing destination name %s in %T", columns[f], dest)
		}
		values = make([]interface{}, len(columns))

		for rows.Next() {
			// 创建新的结构体类型(返回 PtrTo)
			vp = reflect.New(base)
			v = reflect.Indirect(vp)

			err = fieldsByTraversal(v, fields, values, true)
			if err != nil {
				return err
			}

			// 扫描结构体字段的指针, 将结果添加到结果集
			err = rows.Scan(values...)
			if err != nil {
				return err
			}

			if isPtr {
				direct.Set(reflect.Append(direct, vp))
			} else {
				direct.Set(reflect.Append(direct, v))
			}
		}
	} else {
		for rows.Next() {
			vp = reflect.New(base)
			err = rows.Scan(vp.Interface())
			if err != nil {
				return err
			}
			// append
			if isPtr {
				direct.Set(reflect.Append(direct, vp))
			} else {
				direct.Set(reflect.Append(direct, reflect.Indirect(vp)))
			}
		}
	}

	return rows.Err()
}

// 判断基本数据类型
func baseType(t reflect.Type, expected reflect.Kind) (reflect.Type, error) {
	t = Deref(t)
	if t.Kind() != expected {
		return nil, fmt.Errorf("expected %s but got %s", expected, t.Kind())
	}
	return t, nil
}

// fieldsByTraversal 遍历字段
func fieldsByTraversal(v reflect.Value, traversals [][]int, values []interface{}, ptrs bool) error {
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Struct {
		return errors.New("argument not a struct")
	}

	for i, traversal := range traversals {
		if len(traversal) == 0 {
			values[i] = new(interface{})
			continue
		}
		f := FieldByIndexes(v, traversal)
		if ptrs {
			values[i] = f.Addr().Interface()
		} else {
			values[i] = f.Interface()
		}
	}
	return nil
}

// 判断是否丢失了字段
func missingFields(transversals [][]int) (field int, err error) {
	for i, t := range transversals {
		if len(t) == 0 {
			return i, errors.New("missing field")
		}
	}
	return 0, nil
}
