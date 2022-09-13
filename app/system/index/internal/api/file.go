package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/app/system/index/internal/service"
)

// 文件管理
var File = fileApi{}

type fileApi struct{}

// Upload 文件上传
func (a *fileApi) Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	if form == nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "请选择上传的文件"})
		return
	}
	files := form.File["file[]"]

	uploadResult, err := service.File.Upload(c, files)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "上传成功",
		"data": uploadResult,
	})
}