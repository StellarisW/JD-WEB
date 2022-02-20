package model

import g "main/app/global"

type Banner struct {
	g.Model
	Title      string
	BannerType int    `db:"banner_type"`
	BannerImg  string `db:"banner_img"`
	Link       string
	Sort       int
	Status     int
}

func (Banner) TableName() string {
	return "banner"
}
