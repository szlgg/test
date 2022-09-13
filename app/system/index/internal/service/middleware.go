package service

import (
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// // 登录状态
// var IsLogin int

// 中间件管理服务
var Middleware = middlewareService{
	loginUrl: "/login",
}

type middlewareService struct {
	loginUrl string // 登录路由地址
}

// 获取用户登录URL
func (s *middlewareService) GetLoginUrl() string {
	return s.loginUrl
}

// 自定义上下文对象
func (s *middlewareService) Ctx(c *gin.Context) {
	session := sessions.DefaultMany(c, "Login")
	session.Set("url", c.Request.RequestURI)
	session.Save()
}

// 前台系统权限控制，用户必须登录才能访问
func (s *middlewareService) Auth(c *gin.Context) {
	session := sessions.DefaultMany(c, "Login")
	session.Set("url", c.Request.RequestURI)
	session.Save()
	user := Session.GetUser(c)
	if user == nil {
		// c.JSON(http.StatusOK, gin.H{"code": 1, "message": "未登录或会话已过期，请您登录后再继续"})
		c.AddParam("UID", "")
		// return
	} else {
		c.AddParam("UID", strconv.Itoa(int(user.Id)))
	}
}
