package api

import (
	"net/http"
	"strconv"
	"test/app/model"
	"test/app/system/index/internal/define"
	"test/app/system/index/internal/service"

	"github.com/gin-gonic/gin"
)

var Article = articleApi{}

type articleApi struct{}

func (a *articleApi) Index(c *gin.Context) {
	var (
		req *define.ContentGetListReq
	)
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"code": 1, "message": err.Error()})
		return
	}
	req.Type = model.ContentTypeArticle

	sort := GetSort(c.Query("sort"))
	cate := GetCate(c.Query("cate"))

	// 获取栏目
	categoryList, err := service.Category.GetList(req.Type)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"code": 1, "message": err.Error()})
	}

	// 获取文章
	if PageData, getListRes, err := service.Content.GetList(req.ContentGetListInput); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"code": 1, "message": err.Error()})
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

		c.HTML(http.StatusOK, "article/index.html", gin.H{
			"Title":        "问答",
			"categoryList": categoryList,
			"ContentType":  req.Type,
			"Cate":         cate,
			"Sort":         sort,
			"Data":         getListRes,
			"Page":         PageData,
			"User":         userInfo,
		})
	}
}

// Detail 文章详情
func (a *articleApi) Detail(c *gin.Context) {
	// 获取内容id
	idStr := c.Param("id")
	// 转换成int类型
	id, _ := strconv.Atoi(idStr)

	// 获取内容信息与用户信息
	if getDetailRes, err := service.Content.GetDetail(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取内容信息与用户信息失败", "err": err.Error()})
	} else {
		if getDetailRes == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "数据不存在"})
		}
		// 新增浏览量
		if getDetailRes.Content.ViewCount, err = service.Content.AddViewCount(getDetailRes.Content.Id, getDetailRes.Content.ViewCount); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "新增浏览量失败", "err": err.Error()})
		}
		// 获取面包屑信息
		BreadCrumb := service.View.GetBreadCrumb(getDetailRes.Content.Type, getDetailRes.Content.CategoryId)
		// 获取回复
		ReplyList, err := service.Reply.GetList(getDetailRes.Content.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "获取回复失败", "err": err.Error()})
		}
		// 获取登录用户信息
		user := service.Session.GetUser(c)
		userInfo := model.User{}
		if user != nil {
			userInfo.Id = user.Id
		}

		c.HTML(http.StatusOK, "article/detail.html", gin.H{
			"Title":      "详情",
			"Data":       getDetailRes,
			"BreadCrumb": BreadCrumb,
			"User":       userInfo,
			"ReplyList":  ReplyList,
		})
	}
}
