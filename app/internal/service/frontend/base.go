package frontend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	g "main/app/global"
	"main/app/internal/model"
	"main/utils"
	dao "main/utils/sql"
	"strings"
	"time"
)

type sBase struct{}

var insBase = sBase{}

func (s *sBase) InitTopMenu(rs *utils.RedisStore) (topMenu []model.Menu, err error) {
	if hasTopMenu := rs.Get(rs.PreKey+"topMenu", &topMenu); hasTopMenu == true {
		return topMenu, nil
	} else {
		err := g.DB.Select(&topMenu,
			"select * from menu where status=1 and position=1 order by sort desc")
		if err != nil {
			return nil, err
		}
		rs.Set("topMenu", topMenu)
		return topMenu, nil
	}
}

func (s *sBase) InitProductCate(rs *utils.RedisStore) (productCate []model.ProductCate, err error) {
	if hasProductCate := rs.Get(rs.PreKey+"productCate", &productCate); hasProductCate == true {
		return productCate, nil
	} else {
		var productCateItem []model.ProductCate
		err := g.DB.Select(&productCateItem,
			"select * from product_cate where status=1 order by sort desc")
		if err != nil {
			return nil, err
		}
		err = g.DB.Select(&productCate,
			"select * from product_cate where pid=0 and status=1 order by sort desc")
		if err != nil {
			return nil, err
		}
		for i, cate := range productCate {
			cate.ProductCateItem = productCateItem
			productCate[i] = cate
		}
		rs.Set("productCate", productCate)
		return productCate, nil
	}
}

// mysql耗时:5000ms
func (s *sBase) InitMiddleMenu(rs *utils.RedisStore) (middleMenu []model.Menu, err error) {
	if hasMiddleMenu := rs.Get(rs.PreKey+"middleMenu", &middleMenu); hasMiddleMenu == true {
		return middleMenu, nil
	} else {
		err := g.DB.Select(&middleMenu,
			"select * from menu where status=1 and position=2 order by sort desc")
		if err != nil {
			return nil, err
		}
		g.Logger.Debugf("%v\n", time.Now())
		// 获取关联商品 TODO:获取时间长，待优化
		for i := 0; i < len(middleMenu); i++ {
			middleMenu[i].Relation = strings.ReplaceAll(middleMenu[i].Relation, ". ", ",")
			relation := strings.Split(middleMenu[i].Relation, ",")
			var product []model.Product
			query, args, err := dao.In("select * from product where id in (?) limit 6", relation)
			if err != nil {
				return nil, err
			}
			err = g.DB.Select(&product, query, args...)
			if err != nil {
				return nil, err
			}
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
