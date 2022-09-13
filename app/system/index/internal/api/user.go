package api

import (
	"fmt"
	"net/http"
	"strconv"
	"test/app/system/index/internal/define"
	"test/app/system/index/internal/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var User = userApi{}

type userApi struct{}

// Logout 注销退出
func (a *userApi) Logout(c *gin.Context) {
	fmt.Println("开始退出")
	session := sessions.DefaultMany(c, "Login")
	session.Clear()
	err := session.Save()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}

	// 跳转到首页
	// c.Redirect(http.StatusMovedPermanently, "/")
}

// 用户主页
func (a *userApi) Index(c *gin.Context) {
	var (
		req *define.UserGetListReq
	)
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}
	// 获取内容用户id
	idStr := c.Param("id")
	uid, _ := strconv.Atoi(idStr)
	req.Uid = uint(uid)
	req.UserGetListInput.Uid = uint(uid)

	req.Type = c.DefaultQuery("type", "article")
	req.UserGetListInput.Type = c.DefaultQuery("type", "article")

	sort := GetSort(c.Query("sort"))
	cate := GetCate(c.Query("cate"))

	// 获取栏目
	categoryList, err := service.Category.GetList(req.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": err.Error()})
	}

	// 获取回复
	getReplyList := make([]define.ReplyGetListOutput, 0)
	if req.Type == "message" {
		getReplyList, err = service.User.GetMessageList(req.UserGetListInput)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": err.Error()})
		}
	}

	// 获取文章
	if output, getListRes, err := service.User.GetList(req.UserGetListInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": err.Error()})
	} else {
		user := service.Session.GetUser(c)

		c.HTML(http.StatusOK, "user/index.html", gin.H{
			"Title":        "个人中心",
			"Type":         req.Type,
			"categoryList": categoryList,
			"User":         user,
			"Cate":         cate,
			"Sort":         sort,
			"Data":         getListRes,
			"Output":       output,
			"ReplyList":    getReplyList,
		})
	}
}

// Profile 基本资料
func (a *userApi) Profile(c *gin.Context) {
	c.HTML(http.StatusOK, "user/user_info.html", gin.H{
		"Title": "基本资料",
		"Type":  "profile",
		"User":  service.Session.GetUser(c),
	})
}

// UpdateProfile 编辑基本资料
func (a *userApi) UpdateProfile(c *gin.Context) {
	var (
		req *define.UserUpdateProfileInput
	)
	// 获取表单数据
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}
	if err := service.User.UpdateProfile(c, *req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "修改成功"})
	}

}

// Avatar 修改头像
func (a *userApi) Avatar(c *gin.Context) {
	c.HTML(http.StatusOK, "user/user_info.html", gin.H{
		"Title": "修改头像",
		"Type":  "avatar",
		"User":  service.Session.GetUser(c),
	})
}

// 上传图片
func (a *userApi) UpdateAvatar(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	if file == nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "请选择上传的文件"})
		return
	}
	// 上传图片
	uploadResult, err := service.File.UploadAvatar(c, file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}
	// 更新头像
	if err := service.User.UploadAvatar(c, uploadResult); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "上传成功"})
	}
}

// Password 修改密码
func (a *userApi) Password(c *gin.Context) {
	c.HTML(http.StatusOK, "user/user_info.html", gin.H{
		"Title": "修改密码",
		"Type":  "password",
		"User":  service.Session.GetUser(c),
	})
}

// UpdatePassword 更新密码
func (a *userApi) UpdatePassword(c *gin.Context) {
	fmt.Println("0000")
	var (
		req *define.UserPasswordInput
	)
	// 获取表单数据
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
		return
	}
	fmt.Println("req:", req)
	if err := service.User.UpdatePassword(c, *req); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "修改成功"})
	}
}
