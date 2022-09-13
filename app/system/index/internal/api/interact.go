package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/app/system/index/internal/define"
	"test/app/system/index/internal/service"
)

// 交互管理
var Interact = interactApi{}

type interactApi struct{}

// Zan 赞
func (a *interactApi) Zan(c *gin.Context) {
	var (
		req *define.InteractZanReq
	)
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}
	if err := service.Interact.Zan(c, req.Type, int(req.Id)); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": ""})
}

// CancelZan 取消赞
func (a *interactApi) CancelZan(c *gin.Context) {
	var (
		req *define.InteractCancelZanReq
	)
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}
	if err := service.Interact.CancelZan(c, req.Type, int(req.Id)); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": ""})
}

// Cai 踩
func (a *interactApi) Cai(c *gin.Context) {
	var (
		req *define.InteractCaiReq
	)
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}
	if err := service.Interact.Cai(c, req.Type, int(req.Id)); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": ""})
}

// CancelCai 取消踩
func (a *interactApi) CancelCai(c *gin.Context) {
	var (
		req *define.InteractCancelCaiReq
	)
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}
	if err := service.Interact.CancelCai(c, req.Type, int(req.Id)); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": ""})
}

