package api

import (
	"net/http"
	"test/app/system/index/internal/define"
	"test/app/system/index/internal/service"

	"github.com/gin-gonic/gin"
)

// 内容管理
var Content = contentApi{}

type contentApi struct{}

// Create 展示创建内容页面
func (a *contentApi) Create(c *gin.Context) {
	var (
		req *define.ContentCreateReq
	)
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}
	// 获取栏目
	categoryList, err := service.Category.GetList(req.Type)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
	}
	c.HTML(http.StatusOK, "create.html", gin.H{"Title": "发布内容", "categoryList": categoryList, "ContentType": req.Type, "UID": c.Param("UID"), "URL": c.Param("URL")})
}

// DoCreate 创建内容
func (a *contentApi) DoCreate(c *gin.Context) {
	var (
		req *define.ContentDoCreateReq
	)
	// 获取表单数据
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}

	if res, err := service.Content.Create(c, req.ContentCreateInput); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "发布成功", "data": res})
	}
}

// Update 展示更新内容页面
func (a *contentApi) Update(c *gin.Context) {
	var (
		req *define.ContentUpdateReq
	)
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}
	if getDetailRes, err := service.Content.GetDetail(int(req.Id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": err.Error()})
	} else {
		// 获取栏目
		categoryList, err := service.Category.GetList(getDetailRes.Content.Type)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		}

		c.HTML(http.StatusOK, "update.html", gin.H{
			"Title":        "修改内容",
			"Data":         getDetailRes,
			"CategoryList": categoryList,
			"User":         service.Session.GetUser(c),
		})
	}

}

// DoUpdate 更新内容
func (a *contentApi) DoUpdate(c *gin.Context) {
	var (
		req *define.ContentDoUpdateReq
	)
	// 获取表单数据
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}

	if err := service.Content.Update(req.ContentUpdateInput); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "保存成功"})
	}
}

// DoDelete 删除内容
func (a *contentApi) DoDelete(c *gin.Context) {
	var (
		req *define.ContentDoDeleteReq
	)
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
	}
	if err := service.Content.Delete(req.Id); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除成功"})
	}
}
