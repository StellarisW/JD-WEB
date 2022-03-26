package model

import (
	g "main/app/global"
)

type Product struct {
	g.Model
	Title           string
	SubTitle        string `db:"sub_title"`
	ProductSn       string `db:"product_sn"`
	CateId          int    `db:"cate_id"`
	ClickCount      int    `db:"click_count"`
	ProductNumber   int    `db:"product_number"`
	Price           float64
	MarketPrice     float64 `db:"market_price"`
	RelationProduct string  `db:"relation_product"`
	ProductAttr     string  `db:"product_attr"`
	ProductVersion  string  `db:"product_version"`
	ProductImg      string  `db:"product_img"`
	ProductGift     string  `db:"product_gift"`
	ProductFitting  string  `db:"product_fitting"`
	ProductColor    string  `db:"product_color"`
	ProductKeywords string  `db:"product_keywords"`
	ProductDesc     string  `db:"product_desc"`
	ProductContent  string  `db:"product_content"`
	IsDelete        int     `db:"is_delete"`
	IsHot           int     `db:"is_hot"`
	IsBest          int     `db:"is_best"`
	IsNew           int     `db:"is_new"`
	ProductTypeId   int     `db:"product_type_id"`
	Sort            int
	Status          int
}

type ProductCate struct {
	g.Model
	Title           string
	CateImg         string `db:"cate_img"`
	Link            string
	Template        string
	Pid             int
	SubTitle        string `db:"sub_title"`
	Keywords        string
	Description     string
	Sort            int
	Status          int
	ProductCateItem []ProductCate //`gorm:"foreignkey:Pid;association_foreignkey:Id"`
}

type ProductAttr struct {
	g.Model
	ProductId       int    `db:"product_id"`
	AttributeCateId int    `db:"attribute_cate_id"`
	AttributeId     int    `db:"attribute_id"`
	AttributeTitle  string `db:"attribute_title"`
	AttributeType   int    `db:"attribute_type"`
	AttributeValue  string `db:"attribute_value"`
	Sort            int
	Status          int
}

type ProductColor struct {
	g.Model
	ColorName  string `db:"color_name"`
	ColorValue string `db:"color_value"`
	Status     int
	Checked    bool `gorm:"-"`
}

type ProductImage struct {
	g.Model
	ProductId int    `json:"product_id" form:"product_id" db:"product_id"`
	ImgUrl    string `json:"img_url" form:"img_url" db:"img_url"`
	ColorId   int    `json:"color_id" form:"color_id" db:"color_id"`
	Sort      int    `json:"sort"`
	Status    int    `json:"status"`
}

type ProductType struct {
	g.Model
	Title       string
	Description string
	Status      int
}

type ProductTypeAttribute struct {
	g.Model
	CateId    int    `json:"cate_id" form:"cate_id" db:"cate_id"`
	Title     string `json:"title" form:"title"`
	AttrType  int    `json:"attr_type" form:"attr_type" db:"attr_type"`
	AttrValue string `json:"attr_value" form:"attr_value" db:"attr_value"`
	Status    int    `json:"status" form:"status"`
	Sort      int    `json:"sort" form:"sort"`
}

type ProductCollect struct {
	g.Model
	UserId    int `json:"user_id" form:"user_id" db:"user_id"`
	ProductId int `json:"product_id" form:"product_id" db:"product_id"`
}

func (Product) TableName() string {
	return "product"
}

func (ProductCate) TableName() string {
	return "product_cate"
}

func (ProductAttr) TableName() string {
	return "product_attr"
}

func (ProductColor) TableName() string {
	return "product_color"
}

func (ProductImage) TableName() string {
	return "product_image"
}

func (ProductType) TableName() string {
	return "product_type"
}

func (ProductTypeAttribute) TableName() string {
	return "product_type_attribute"
}

func (ProductCollect) TableName() string {
	return "product_collect"
}

func GetProductByCategory(cateId int, productType string, limitNum int) ([]Product, error) {
	var product []Product
	var productCate []ProductCate
	_ = g.DB.Select(&productCate,
		"select * from product_cate where pid=?", cateId)
	var tempSlice []int
	if len(productCate) > 0 {
		for i := 0; i < len(productCate); i++ {
			tempSlice = append(tempSlice, productCate[i].Id)
		}
	}
	tempSlice = append(tempSlice, cateId)
	where := "where cate_id in (?)"
	switch productType {
	case "hot":
		where += " AND is_hot=1 "
	case "best":
		where += " AND is_best=1 "
	case "new":
		where += " AND is_new=1 "
	default:
		break
	}

	g.DB.Where(where, tempSlice).Select("id,title,price,product_img,sub_title").Limit(limitNum).Order("sort desc").Find(&product)
	return product, nil
}
