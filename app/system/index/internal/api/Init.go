package api

import (
	"net/http"
	"test/app/system/index/internal/define"
	"test/app/system/index/internal/service"

	"github.com/gin-gonic/gin"
)

var IsLogin = false

//
// Init
//  @Description:
//  @param c
//
func Init(c *gin.Context) {
	var userInfo define.UserGetProfileOutput

	user := service.Session.GetUser(c)
	if user != nil {
		userInfo = define.UserGetProfileOutput{
			Id:       user.Id,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
			Gender:   user.Gender,
		}
		IsLogin = true
	} else {
		IsLogin = false
	}

	c.JSON(http.StatusOK, gin.H{
		"isLogin": IsLogin,
		"User":    userInfo,
	})
}
