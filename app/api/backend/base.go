package backend

import (
	"errors"
	"github.com/gin-gonic/gin"
	g "main/app/global"
	"main/app/internal/model"
	"main/utils"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

type BaseApi struct{}

var insBase = BaseApi{}

func (a *BaseApi) Success(c gin.Context, message string, redirect string) {
	if strings.Contains(redirect, "http") {
		c.HTML(http.StatusOK, "backend/public/success.tmpl", gin.H{
			"Message":  message,
			"Redirect": redirect,
		})
	} else {
		c.HTML(http.StatusOK, "backend/public/success.tmpl", gin.H{
			"Message":  message,
			"Redirect": "/backend" + redirect,
		})
	}
}

func (a *BaseApi) Error(c gin.Context, message string, redirect string) {
	if strings.Contains(redirect, "http") {
		c.HTML(http.StatusOK, "backend/public/error.tmpl", gin.H{
			"Message":  message,
			"Redirect": redirect,
		})
	} else {
		c.HTML(http.StatusOK, "backend/public/error.tmpl", gin.H{
			"Message":  message,
			"Redirect": "/backend" + redirect,
		})
	}
}

func (a *BaseApi) Goto(c gin.Context, redirect string) {
	c.Redirect(302, "/backend"+redirect)
}

func (a *BaseApi) UploadImg(c gin.Context, picName string) (string, error) {
	return a.LocalUploadImg(c, picName)
}

func (a *BaseApi) LocalUploadImg(c gin.Context, picName string) (string, error) {
	h, err := c.FormFile(picName) // TODO:
	if err != nil {
		return "", err
	}
	//2、关闭文件流
	//3、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(h.Filename)

	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("图片后缀名不合法")
	}
	//4、创建图片保存目录  static/upload/20200623
	day := utils.FormatDay()
	dir := "static/upload/" + day

	if err := os.MkdirAll(dir, 0666); err != nil {
		return "", err
	}
	//5、生成文件名称   144325235235.png
	fileUnixName := strconv.FormatInt(utils.GetUnixNano(), 10)
	//static/upload/20200623/144325235235.png
	saveDir := path.Join(dir, fileUnixName+extName)
	//6、保存图片
	c.FileAttachment(saveDir, picName) // TODO:
	return saveDir, nil
}

func (a *BaseApi) GetSetting() model.Setting {
	setting := model.Setting{}
	setting.Id = 1
	g.DB.Where("id=?", setting.Id)
	return setting
}
