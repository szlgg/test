package api

import (
	"net/http"
	"strconv"
	"test/app/model"
	"test/app/system/index/internal/define"
	"test/app/system/index/internal/service"

	"github.com/gin-gonic/gin"
)

var int1 int

// Index
// @summary 展示站点首页
// @tags    前台-首页
// @produce html
// @router  / [GET]
// @success 200 {string} html "页面HTML"
func Index(c *gin.Context) {
	var (
		req *define.ContentGetListReq
	)

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": err.Error()})
		return
	}

	req.Type = model.ContentTypeTopic

	sort := GetSort(c.Query("sort"))
	cate := GetCate(c.Query("cate"))

	// 获取栏目
	categoryList, err := service.Category.GetList(req.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": err.Error()})
	}
	// 获取文章
	if PageData, getListRes, err := service.Content.GetList(req.ContentGetListInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": err.Error()})
	} else {
		// 获取登录用户信息
		var userInfo define.UserGetProfileOutput
		user := service.Session.GetUser(c)
		if user != nil {
			userInfo = define.UserGetProfileOutput{
				Id:       user.Id,
				Nickname: user.Nickname,
				Avatar:   user.Avatar,
				Gender:   user.Gender,
			}
		}

		c.HTML(http.StatusOK, "home/index.html", gin.H{
			"Title":        "首页",
			"categoryList": categoryList,
			"ContentType":  "",
			"Cate":         cate,
			"Sort":         sort,
			"Data":         getListRes,
			"Page":         PageData,
			"User":         userInfo,
		})
	}
}

// GetCate 获取栏目ID转换为int
func GetCate(cate string) int {
	int1, _ := strconv.Atoi(cate)
	return int1
}

// GetSort 获取排序方式转换为int
func GetSort(sort string) int {
	switch {
	case sort == "1":
		int1 = 1
	case sort == "2":
		int1 = 2
	case sort == "3":
		int1 = 3
	default:
		int1 = 0
	}
	return int1
}
