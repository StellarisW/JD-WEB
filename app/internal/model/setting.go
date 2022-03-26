package model

import (
	g "main/app/global"
)

type Setting struct {
	g.Model
	SiteTitle       string `json:"site_title" form:"site_title" db:""`
	SiteLogo        string `json:"site_logo" form:"site_logo" db:"site_logo"`
	SiteKeywords    string `json:"site_keywords" form:"site_keywords" db:""`
	SiteDescription string `json:"site_description" form:"site_description" db:"site_description"`
	NoPicture       string `json:"no_picture" form:"no_picture" db:"no_picture"`
	SiteIcp         string `json:"site_icp" form:"site_icp" db:"site_icp"`
	SiteTel         string `json:"site_tel" form:"site_tel" db:"site_tel"`
	SearchKeywords  string `json:"search_keywords" form:"search_keywords" db:"search_keywords"`
	TongjiCode      string `json:"tongji_code" form:"tongji_code" db:"tongji_code"`
	Appid           string `json:"appid" form:"appid"`
	AppSecret       string `json:"app_secret" form:"app_secret" db:"app_secret"`
	EndPoint        string `json:"end_point" form:"end_point" db:"end_point"`
	BucketName      string `json:"bucket_name" form:"bucket_name" db:"bucket_name"`
	OssStatus       int    `json:"oss_status" form:"oss_status" db:"oss_status"`
}
