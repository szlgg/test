package api

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/app/model"
	"test/app/system/index/internal/define"
	"test/app/system/index/internal/service"
)

// 登录展示页面
func Login(c *gin.Context) {

	c.HTML(http.StatusOK, "login/index.html", gin.H{"Title": "登录"})
}

// 提交登录
func LoginPost(c *gin.Context)  {
	var (
		req *define.UserLoginReq
	)
	gob.Register(define.UserLoginReq{})
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
	// 验证数据
	if err := service.User.Login(c, req.UserLoginInput); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}

	session := sessions.DefaultMany(c, "Login")

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录成功", "url": session.Get("url")})
}
