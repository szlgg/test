package api

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/app/model"
	"test/app/system/index/internal/define"
	"test/app/system/index/internal/service"
)

//var req define.UserRegisterReq

// @summary 展示注册页面
// @tags    前台-注册
// @produce html
// @router  /register [GET]
// @success 200 {string} html "页面HTML"
func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register/index.html", gin.H{"Title": "注册"})
}

// 提交注册
func Do(c *gin.Context) {
	var (
		req *define.UserRegisterReq
	)
	gob.Register(define.UserRegisterReq{})
	// 获取表单数据
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}
	// 判断验证码
	if !service.Captcha.VerifyAndClear(c, model.CaptchaDefaultName, req.Captcha) {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "验证码错误"})
		return
	}

	// 插入数据
	if err := service.User.Register(req.UserRegisterInput); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "注册成功"})
}