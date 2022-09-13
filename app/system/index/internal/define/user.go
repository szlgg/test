package define

import "test/app/model"

// 查询用户列表请求
type UserGetListReq struct {
	UserGetListInput
	Id uint `form:"id,default=0"`
}

// 查询用户详情结果
type UserGetListOutput struct {
	Content *ContentGetListOutput `json:"content"`
	User    model.User            `json:"user"`
	Stats   map[string]int64      // 发布内容数
}

// 用户信息
type UserGetProfileOutput struct {
	Id       uint           // 用户ID
	Nickname string         // 昵称
	Avatar   string         // 头像地址
	Gender   int            // 性别 0: 未设置 1: 男 2: 女
}

// 查询用户列表输入
type UserGetListInput struct {
	ContentGetListInput
	Id uint `form:"id"`
}

// 用户注册
type UserRegisterReq struct {
	UserRegisterInput
	Passport string `form:"passport" binding:"required"` // 账号
	Password string `form:"password" binding:"required"` // 密码
	Nickname string `form:"nickname" binding:"required"` // 昵称
	Captcha  string `form:"captcha"  binding:"required"`  // 验证码
}

// 创建用户
type UserRegisterInput struct {
	Passport string `form:"passport"` // 账号
	Password string `form:"password"` // 密码(明文)
	Nickname string `form:"nickname"` // 昵称
}

// Api用户登录
type UserLoginReq struct {
	UserLoginInput
	Passport string `form:"passport"` // 账号
	Password string `form:"password"` // 密码
	Captcha  string `form:"captcha"`  // 验证码
}

// 用户登录
type UserLoginInput struct {
	Passport string `form:"passport"` // 账号
	Password string `form:"password"` // 密码(明文)
}

// 修改用户信息
type UserUpdateProfileInput struct {
	Id       uint   `form:"id"`       // 用户ID
	Nickname string `form:"nickname"` // 昵称
	Avatar   string `form:"avatar"`   // 头像地址
	Gender   int    `form:"gender"`   // 性别 0: 未设置 1: 男 2: 女
}

// 修改用户密码
type UserPasswordInput struct {
	Id          uint   `form:"id"`          // 用户ID
	OldPassword string `form:"oldPassword"` // 原密码
	NewPassword string `form:"newPassword"` // 新密码
}
