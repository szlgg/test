package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var exceptionExit interface{} = "exit"

// 数据返回通用JSON数据结构
type JsonRes struct {
	Code     int         `json:"code"`     // 错误码((0:成功, 1:失败, >1:错误码))
	Message  string      `json:"message"`  // 提示信息
	Data     interface{} `json:"data"`     // 返回数据(业务接口定义具体数据结构)
	Redirect string      `json:"redirect"` // 引导客户端跳转到指定路由
}

// 返回标准JSON数据。
func Json(c *gin.Context, code int, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = make(map[string]interface{})
	}
	c.JSON(http.StatusBadRequest, JsonRes{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

// 返回标准JSON数据并退出当前HTTP执行函数。
func JsonExit(c *gin.Context, code int, message string, data ...interface{}) {
	Json(c, code, message, data...)
	//panic(exceptionExit)
	return
}

