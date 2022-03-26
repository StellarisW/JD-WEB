package frontend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	g "main/app/global"
	"main/app/internal/model"
	"main/utils"
	"strings"
	"time"
)

type sBase struct{}

var insBase = sBase{}

func (s *sBase) InitTopMenu(rs *utils.RedisStore) (topMenu []model.Menu, err error) {
	if hasTopMenu := rs.Get(rs.PreKey+"topMenu", &topMenu); hasTopMenu == true {
		return topMenu, nil
	} else {
		g.DB.Where("status=1 AND position=1").Order("sort desc").Find(&topMenu)
		rs.Set("topMenu", topMenu)
		return topMenu, nil
	}
}

func (s *sBase) InitProductCate(rs *utils.RedisStore) (productCate []model.ProductCate, err error) {
	if hasProductCate := rs.Get(rs.PreKey+"productCate", &productCate); hasProductCate == true {
		return productCate, nil
	} else {
		g.DB.Preload("ProductCateItem",
			func(db *gorm.DB) *gorm.DB {
				return db.Where("product_cate.status=1").
					Order("product_cate.sort DESC")
			}).Where("pid=0 AND status=1").Order("sort desc").
			Find(&productCate)
		rs.Set("productCate", productCate)
		return productCate, nil
	}
}

// mysql耗时:5000ms
func (s *sBase) InitMiddleMenu(rs *utils.RedisStore) (middleMenu []model.Menu, err error) {
	if hasMiddleMenu := rs.Get(rs.PreKey+"middleMenu", &middleMenu); hasMiddleMenu == true {
		return middleMenu, nil
	} else {
		g.DB.Where("status=1 AND position=2").Order("sort desc").Find(&middleMenu)
		g.Logger.Debugf("%v\n", time.Now())
		// 获取关联商品 TODO:获取时间长，待优化
		for i := 0; i < len(middleMenu); i++ {
			middleMenu[i].Relation = strings.ReplaceAll(middleMenu[i].Relation, ". ", ",")
			relation := strings.Split(middleMenu[i].Relation, ",")
			var product []model.Product
			g.DB.Where("id in (?)", relation).Limit(6).Order("sort ASC").
				Select("id,title,product_img,price").Find(&product)
			middleMenu[i].ProductItem = product
			g.Logger.Debugf("index:%v time:%v\n", i, time.Now())
		}
		rs.Set("middleMenu", middleMenu)
		return middleMenu, nil
	}
}

func (s *sBase) IsLogin(c *gin.Context) string {
	//cookie.get
	phone := utils.GetUserPhone(c)
	if len(phone) == 11 {
		str := fmt.Sprintf(`<ul>
			<li class="userinfo">
				<a href="#">%v</a>

				<i class="i"></i>
				<ol>
					<li><a href="/user">个人中心</a></li>

					<li><a href="/user/collect">我的收藏</a></li>

					<li><a href="/logout">退出登录</a></li>
				</ol>

			</li>
		</ul> `, phone)
		return str
	} else {
		str := fmt.Sprintf(`<ul>
			<li><a href="/login" target="_blank">登录</a></li>
			<li>|</li>
			<li><a href="/register_step1" target="_blank" >注册</a></li>
		</ul>`)
		return str
	}
}
