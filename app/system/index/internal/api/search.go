package api

import (
	"net/http"
	"test/app/system/index/internal/define"
	"test/app/system/index/internal/service"

	"github.com/gin-gonic/gin"
)

var Search = SearchApi{}

type SearchApi struct{}

// 搜索
func (a *SearchApi) Index(c *gin.Context) {
	var (
		req *define.SearchGetListReq
	)
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": err.Error()})
		return
	}

	// 获取文章
	if output, getListRes, err := service.Content.Search(req.SearchInput); err != nil {
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

		c.HTML(http.StatusOK, "search/index.html", gin.H{
			"Title": "搜索",
			"Stats": output.Stats,
			"Type":  req.Type,
			"Sort":  req.Sort,
			"Data":  getListRes,
			"Page":  output.PageData,
			"User":  userInfo,
			"Key":   req.Key,
		})
	}
}
