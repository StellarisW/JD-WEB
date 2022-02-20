package frontend

import (
	g "main/app/global"
	"main/app/internal/model"
	dao "main/utils/sql"
	"strings"
)

type sProduct struct{}

var insProduct = sProduct{}

func (s *sProduct) GetCateList(cateId int) (model.ProductCate, []model.ProductCate) {
	var currentProductCate model.ProductCate
	var subProductCate []model.ProductCate

	_ = g.DB.Get(&currentProductCate, "select * from product_cate where id=?", cateId)
	_ = g.DB.Select(&subProductCate, "select * from product_cate where pid=?", currentProductCate.Id)

	return currentProductCate, subProductCate
}

func (s *sProduct) GetProductByCate(ids []int, page, pageSize int) ([]model.Product, int) {
	var products []model.Product
	var count int
	where, args, _ := dao.In("where cate_id in (?) ", ids)
	vArgs := append(args, pageSize)
	vArgs = append(vArgs, (page-1)*pageSize)
	_ = g.DB.Select(&products,
		"select id,title,price,product_img,sub_title from product "+where+"order by sort desc limit ? offset ?", vArgs...)
	//查询product表里面的数量
	_ = g.DB.QueryRow("select count(*) from product "+where, args...).Scan(&count)
	return products, count
}

func (s *sProduct) GetProductById(id int) (product model.Product) {
	_ = g.DB.Get(&product, "select * from product where id=?", id)
	return product
}

func (s *sProduct) GetRelationProduct(product model.Product) (rProduct []model.Product) {
	product.RelationProduct = strings.ReplaceAll(product.RelationProduct, "，", ",")
	relationIds := strings.Split(product.RelationProduct, ",")
	str, args, _ := dao.In("select id,title,price,product_version from product where id in (?)", relationIds)
	_ = g.DB.Select(&rProduct, str, args...)
	return rProduct
}

func (s *sProduct) GetGift(product model.Product) (productGift []model.Product) {
	product.ProductGift = strings.ReplaceAll(product.ProductGift, "，", ",")
	giftIds := strings.Split(product.ProductGift, ",")
	str, args, _ := dao.In("select id,title,price,product_img from product where id in (?)", giftIds)
	_ = g.DB.Select(&productGift, str, args...)
	return productGift
}

func (s *sProduct) GetColor(product model.Product) (productColor []model.ProductColor) {
	product.ProductColor = strings.ReplaceAll(product.ProductColor, "，", ",")
	colorIds := strings.Split(product.ProductColor, ",")
	str, args, _ := dao.In("select * from product_color where id in (?)", colorIds)
	_ = g.DB.Select(&productColor, str, args...)
	return productColor
}

func (s *sProduct) GetFitting(product model.Product) (productFitting []model.Product) {
	product.ProductFitting = strings.ReplaceAll(product.ProductFitting, "，", ",")
	fittingIds := strings.Split(product.ProductFitting, ",")
	str, args, _ := dao.In("select id,title,price,product_img from product where id in (?)", fittingIds)
	_ = g.DB.Select(&productFitting, str, args...)
	return productFitting
}

func (s *sProduct) GetImg(product model.Product) (productImg []model.ProductImage) {
	_ = g.DB.Select(&productImg, "select * from product_image where product_id=?", product.Id)
	return productImg
}

func (s *sProduct) GetAttr(product model.Product) (productAttr []model.ProductAttr) {
	_ = g.DB.Select(&productAttr, "select * from product_attr where product_id=?", product.Id)
	return productAttr
}

func (s *sProduct) GetImgList(productId, colorId int) (imageList []model.ProductImage) {
	_ = g.DB.Select(&imageList,
		"select * from product_image where product_id=? and color_id=?", productId, colorId)
	if len(imageList) == 0 {
		_ = g.DB.Select(&imageList, "select * from product_image where product_id=?", productId)
	}
	return imageList
}

func (s *sProduct) CollectProduct(userId, productId int) bool {
	var user model.ProductCollect
	_ = g.DB.Get(&user, "select * from product_collect where user_id=? and product_id=?", userId, productId)
	if user.Id == 0 {
		_, _ = g.DB.Exec("insert into product_collect(user_id,product_id) values (?,?)",
			userId, productId)
		return true
	} else {
		_, _ = g.DB.Exec("delete from product_collect where user_id=? and product_id=?",
			userId, productId)
		return false
	}
}
