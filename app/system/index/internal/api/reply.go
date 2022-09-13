package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/app/system/index/internal/define"
	"test/app/system/index/internal/service"
)

// 回复控制器
var Reply = replyApi{}

type replyApi struct{}

// DoCreate 新增回复
func (a *replyApi) DoCreate(c *gin.Context) {
	var (
		req *define.ReplyDoCreateReq
	)

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}

	if err := service.Reply.Create(c, req.ReplyCreateInput); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "评论成功"})
	}
}

// DoDelete 删除回复
func (a *replyApi) DoDelete(c *gin.Context) {
	var (
		req *define.ReplyDoDeleteReq
	)
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
	}
	if err := service.Reply.Delete(req.Id); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除成功"})
	}
}