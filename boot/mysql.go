package boot

import (
	_ "github.com/go-sql-driver/mysql"
	"main/app/global"
	dao "main/utils/sql"
	"time"
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

func MySQLSetup() {
	config := g.Config.Mysql

	db, err := dao.ConnectDB("mysql", config.Dsn())
	if err != nil {
		g.Logger.Fatalf("Initialize MySQL server failed, err: %v\n", err)
	}
	db.SetConnMaxIdleTime(10 * time.Second)  // 最大空闲时间
	db.SetConnMaxLifetime(100 * time.Second) // 最大存活时间
	db.SetMaxIdleConns(10)                   // 最大空闲连接数
	db.SetMaxOpenConns(100)                  // 最大连接数
	err = db.Ping()
	if err != nil {
		g.Logger.Fatalf("Connect to MySQL server failed, err: %v\n", err)
	}
	g.Logger.Info("Initialize MySQL server successfully!")
	g.DB = db
}
