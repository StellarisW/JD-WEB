package router

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"main/utils"
	"main/utils/captcha"
)

type TplFuncMap = template.FuncMap

var tplFunMap = TplFuncMap{}

func SetFuncMap(r *gin.Engine) {
	cpt := captcha.NewCaptcha()

	tplFunMap["formatAttribute"] = utils.FormatAttribute
	tplFunMap["str2html"] = utils.Str2html
	tplFunMap["mul"] = utils.Mul
	tplFunMap["eq"] = utils.Eq

	r.SetFuncMap(tplFunMap)

	r.Any(cpt.UrlPrefix, cpt.Handler)
}
