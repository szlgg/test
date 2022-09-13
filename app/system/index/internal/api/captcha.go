package api

import (
	"test/app/model"
	"test/app/system/index/internal/service"

	"github.com/gin-gonic/gin"
)

// @summary 获取默认的验证码
// @description 注意直接返回的是图片二进制内容。
// @tags    前台-验证码
// @produce png
// @router  /captcha [GET]
// @success 200 {file} body "验证码二进制内容"
func Captcha(c *gin.Context) {
	err := service.Captcha.NewAndStore(c, model.CaptchaDefaultName)
	if err != nil {
		return
	}
}
