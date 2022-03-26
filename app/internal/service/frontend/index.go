package frontend

import (
	g "main/app/global"
	"main/app/internal/model"
	"main/utils"
)

type sIndex struct{}

var insIndex = sIndex{}

func (s *sIndex) GetBanner(rs *utils.RedisStore) ([]model.Banner, error) {
	var banner []model.Banner
	if hasBanner := rs.Get(rs.PreKey+"banner", &banner); hasBanner == true {
		return banner, nil
	}
	g.DB.Where("status=1 AND banner_type=1").Order("sort desc").Find(&banner)
	rs.Set("banner", banner)
	return banner, nil
}

func (s *sIndex) GetPhoneProduct(rs *utils.RedisStore) ([]model.Product, error) {
	var PhoneProduct []model.Product
	if hasPhoneProduct := rs.Get(rs.PreKey+"phone", &PhoneProduct); hasPhoneProduct == true {
		return PhoneProduct, nil
	}
	PhoneProduct, err := model.GetProductByCategory(1, "hot", 8)
	if err != nil {
		return nil, err
	}
	return PhoneProduct, nil
}

func (s *sIndex) GetTvProduct(rs *utils.RedisStore) ([]model.Product, error) {
	var TvProduct []model.Product
	if hasTvProduct := rs.Get(rs.PreKey+"tv", &TvProduct); hasTvProduct == true {
		return TvProduct, nil
	}
	TvProduct, err := model.GetProductByCategory(4, "best", 8)
	if err != nil {
		return nil, err
	}
	return TvProduct, nil
}
