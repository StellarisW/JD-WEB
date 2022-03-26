package frontend

import (
	g "main/app/global"
	"main/app/internal/model"
	"strings"
)

type sProduct struct{}

var insProduct = sProduct{}

func (s *sProduct) GetCateList(cateId int) (model.ProductCate, []model.ProductCate) {
	var currentProductCate model.ProductCate
	var subProductCate []model.ProductCate
	g.DB.Where("id=?", cateId).Find(&currentProductCate)
	g.DB.Where("pid=?", currentProductCate.Id).Find(&subProductCate)
	return currentProductCate, subProductCate
}

func (s *sProduct) GetProductByCate(ids []int, page, pageSize int) ([]model.Product, int64) {
	var products []model.Product
	var count int64
	where := "cate_id in (?)"
	g.DB.Where(where, ids).Select("id,title,price,product_img,sub_title").Offset((page - 1) * pageSize).
		Limit(pageSize).Order("sort desc").Find(&products)
	g.DB.Where(where, ids).Table("product").Count(&count)
	return products, count
}

func (s *sProduct) GetProductById(id int) (product model.Product) {
	g.DB.Where("id=?", id).Find(&product)
	return product
}

func (s *sProduct) GetRelationProduct(product model.Product) (rProduct []model.Product) {
	product.RelationProduct = strings.ReplaceAll(product.RelationProduct, "，", ",")
	relationIds := strings.Split(product.RelationProduct, ",")
	g.DB.Where("id in (?)", relationIds).Select("id,title,price,product_version").Find(&rProduct)
	return rProduct
}

func (s *sProduct) GetGift(product model.Product) (productGift []model.Product) {
	product.ProductGift = strings.ReplaceAll(product.ProductGift, "，", ",")
	giftIds := strings.Split(product.ProductGift, ",")
	g.DB.Where("id in (?)", giftIds).Select("id,title,price,product_img").Find(&productGift)
	return productGift
}

func (s *sProduct) GetColor(product model.Product) (productColor []model.ProductColor) {
	product.ProductColor = strings.ReplaceAll(product.ProductColor, "，", ",")
	colorIds := strings.Split(product.ProductColor, ",")
	g.DB.Where("id in (?)", colorIds).Find(&productColor)
	return productColor
}

func (s *sProduct) GetFitting(product model.Product) (productFitting []model.Product) {
	product.ProductFitting = strings.ReplaceAll(product.ProductFitting, "，", ",")
	fittingIds := strings.Split(product.ProductFitting, ",")
	g.DB.Where(" id in (?)", fittingIds).Select("id,title,price,product_img").Find(&productFitting)
	return productFitting
}

func (s *sProduct) GetImg(product model.Product) (productImg []model.ProductImage) {
	g.DB.Where("product_id=?", product.Id).Find(&productImg)
	return productImg
}

func (s *sProduct) GetAttr(product model.Product) (productAttr []model.ProductAttr) {
	g.DB.Where("product_id=?", product.Id).Find(&productAttr)
	return productAttr
}

func (s *sProduct) GetImgList(productId, colorId int) (imageList []model.ProductImage) {
	g.DB.Where("product_id=? AND color_id=?", productId, colorId).Find(&imageList)
	if len(imageList) == 0 {
		g.DB.Where("product_id=?", productId).Find(&imageList)
	}
	return imageList
}

func (s *sProduct) CollectProduct(userId, productId int) bool {
	var user model.ProductCollect
	isExist := g.DB.Where("user_id=? AND product_id=?", userId, productId).First(&user)
	if isExist.RowsAffected == 0 {
		return true
	}
	return false
}
