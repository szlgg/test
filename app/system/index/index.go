package index

import (
	"encoding/gob"
	"html/template"
	"test/app/model"
	"test/app/system/index/internal/api"
	"test/app/system/index/internal/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"FormatTime":   service.FormatTime,
		"HTML":         service.HTML,
		"DidIZan":      service.DidIZan,
		"DidICai":      service.DidICai,
		"GenderFont":   service.GenderFont,
		"Gender":       service.Gender,
		"ContentTitle": service.ContentTitle,
		"CategoryName": service.CategoryName,
	})

	// 注册结构体
	gob.Register(model.User{})
	// 设置密钥
	store := cookie.NewStore([]byte("secret"))
	sessionNames := []string{"Captcha", "Login"}
	router.Use(sessions.SessionsMany(sessionNames, store))

	// 设置模板目录
	router.LoadHTMLGlob("./template/index/**/*")

	// 路由组中添加中间件
	authorized := router.Group("/")
	authorized.Use(service.Middleware.Ctx)
	{
		// 初始化
		authorized.POST("/init", api.Init)
		// 登录
		authorized.GET("/login", api.Login)
		authorized.POST("/login", api.LoginPost)
		// 注册
		authorized.GET("/register", api.Register)
		authorized.POST("/register", api.Do)
		// 验证码
		authorized.GET("/captcha", api.Captcha)
		// 搜索
		authorized.GET("/search", api.Search.Index)
		// 首页
		authorized.GET("/", api.Index)
		// 主题
		authorized.GET("/topic", api.Topic.Index)
		authorized.GET("/topic/:id", api.Topic.Detail)
		// 问答
		authorized.GET("/ask", api.Ask.Index)
		authorized.GET("/ask/:id", api.Ask.Detail)
		// 文章
		authorized.GET("/article", api.Article.Index)
		authorized.GET("/article/:id", api.Article.Detail)
		// 回复
		authorized.POST("/reply", api.Reply.DoCreate)
		authorized.POST("/reply/do-delete", api.Reply.DoDelete)
	}

	// 路由组中添加中间件
	authority := router.Group("/")
	authority.Use(service.Middleware.Auth)
	{
		// 文件上传
		authority.POST("/file/upload", api.File.Upload)

		// 用户
		user := authority.Group("/user")
		{
			// 用户主页
			user.GET("/:id", api.User.Index)
			// 基本资料
			user.GET("/profile", api.User.Profile)
			user.POST("/update-profile", api.User.UpdateProfile)
			// 修改头像
			user.GET("/avatar", api.User.Avatar)
			user.POST("/update-avatar", api.User.UpdateAvatar)
			// 修改密码
			user.GET("/password", api.User.Password)
			user.POST("/update-password", api.User.UpdatePassword)
			// 注销
			user.GET("/logout", api.User.Logout)
		}
		// 交互
		interact := authority.Group("/interact")
		{
			// 赞
			interact.POST("/zan", api.Interact.Zan)
			// 取消赞
			interact.POST("/cancel-zan", api.Interact.CancelZan)
			// 踩
			interact.POST("/cai", api.Interact.Cai)
			// 取消踩
			interact.POST("/cancel-cai", api.Interact.CancelCai)
		}

		// 发布
		C1 := authority.Group("/content")
		{
			// 显示发布页面
			C1.GET("/create", api.Content.Create)
			// 创建发布
			C1.POST("/docreate", api.Content.DoCreate)
			// 显示更新页面
			C1.GET("/update", api.Content.Update)
			// 更新内容
			C1.POST("/do-update", api.Content.DoUpdate)
			// 删除内容
			C1.POST("/do-delete", api.Content.DoDelete)
		}
	}

	return router
}
