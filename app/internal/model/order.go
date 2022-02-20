package model

import g "main/app/global"

type Order struct {
	g.Model
	OrderId     string `db:"order_id"`
	Uid         int
	AllPrice    float64 `db:"all_price"`
	Phone       string
	Name        string
	Address     string
	Zipcode     string
	PayStatus   int         `db:"pay_status"`   // 支付状态： 0 表示未支付     1 已经支付
	PayType     int         `db:"pay_type"`     // 支付类型： 0 alipay    1 wechat
	OrderStatus int         `db:"order_status"` // 订单状态： 0 已下单  1 已付款  2 已配货  3、发货   4、交易成功   5、退货     6、取消
	OrderItem   []OrderItem //`gorm:"foreignkey:OrderId;association_foreignkey:Id"`
}
type OrderItem struct {
	g.Model
	OrderId        string `db:"order_id"`
	Uid            int
	ProductTitle   string  `db:"product_title"`
	ProductId      int     `db:"product_id"`
	ProductImg     string  `db:"product_img"`
	ProductPrice   float64 `db:"product_price"`
	ProductNum     int     `db:"product_num"`
	ProductVersion string  `db:"product_version"`
	ProductColor   string  `db:"product_color"`
}

func (Order) TableName() string {
	return "order"
}

func (OrderItem) TableName() string {
	return "order_item"
}
