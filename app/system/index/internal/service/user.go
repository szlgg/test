package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/o1egl/govatar"
	"test/app/model"
	"test/app/system/index/internal/define"
	//"test/app/system/index/internal/service"
	"test/database"
	"test/library/gerroraa"
	"test/library/utils"
	"xorm.io/builder"
)

//定义全局配置文件*.ini
//var cfg *ini.File

//var engine *xorm.Engine
var engine = database.InitMysql()

// 用户管理服务
var User = userService{
	//avatarUploadPath:      cfg.Section("upload").Key("path").String() + `/avatar`,
	avatarUploadPath:      `upload/avatar`,
	avatarUploadUrlPrefix: `/upload/avatar`,
}

type userService struct {
	avatarUploadPath      string // 头像上传路径
	avatarUploadUrlPrefix string // 头像上传对应的URL前缀
}

// 将密码按照内部算法进行加密
func (s *userService) EncryptPassword(password string) string {
	return utils.MustEncrypt(password)
}

// 根据账号和密码查询用户信息，一般用于账号密码登录。
// 注意password参数传入的是按照相同加密算法加密过后的密码字符串。
func (s *userService) GetUserByPassportAndPassword(passport, password string) (user *model.User, err error, result bool) {
	user = &model.User{Passport: passport, Password: password}
	result, err = engine.Where(builder.Eq{"passport": passport, "password": password}).Get(user)
	return
}

// 执行登录
func (s *userService) Login(c *gin.Context, input define.UserLoginInput) error {
	userEntity, err, result := s.GetUserByPassportAndPassword(
		input.Passport,
		s.EncryptPassword(input.Password),
	)
	if err != nil {
		return err
	}
	if result == false {
		return gerroraa.New(`账号或密码错误`)
	}
	if err := Session.SetUser(c, userEntity); err != nil {
		return err
	}
	//存储session
	//session := sessions.DefaultMany(c, "Login")
	//session.Set(service.SessionKeyUser, &model.ContextUser{
	//	Id:       userEntity.Id,
	//	Passport: userEntity.Passport,
	//	Nickname: userEntity.Nickname,
	//	Avatar:   userEntity.Avatar,
	//})
	//err = session.Save()

	return err
}

// 检测给定的账号是否唯一
func (s *userService) CheckPassportUnique(passport string) error {
	user := &model.User{Passport: passport}
	n, err := engine.Exist(user)
	if err != nil {
		return err
	}
	if n == true {
		return gerroraa.Newf(`账号"%s"已被占用`, passport)
	}
	return nil
}

// 检测给定的昵称是否唯一
func (s *userService) CheckNicknameUnique(nickname string) error {
	user := &model.User{Nickname: nickname}
	n, err := engine.Exist(user)
	if err != nil {
		return err
	}
	if n == true {
		return gerroraa.Newf(`昵称"%s"已被占用`, nickname)
	}
	return nil
}

// 用户注册。
func (s *userService) Register(input define.UserRegisterInput) error {
	user := model.User{Passport: input.Passport, Password: input.Password, Nickname: input.Nickname}
	// 检测给定的账号是否唯一
	if err := s.CheckPassportUnique(input.Passport); err != nil {
		return err
	}
	// 检测给定的昵称是否唯一
	if err := s.CheckNicknameUnique(user.Nickname); err != nil {
		return err
	}
	user.Password = s.EncryptPassword(user.Password)
	// 自动生成头像
	avatarFilePath := fmt.Sprintf(`%s/%s.jpg`, s.avatarUploadPath, user.Passport)
	if err := govatar.GenerateFileForUsername(govatar.MALE, user.Passport, avatarFilePath); err != nil {
		return gerroraa.Wrapf(err, `自动创建头像失败`)
	}
	user.Avatar = fmt.Sprintf(`%s/%s.jpg`, s.avatarUploadUrlPrefix, user.Passport)
    // 插入数据
	_, err := engine.Insert(&user)

	return err
}

// 内容列表
func (s *userService) GetList(input define.UserGetListInput) (output *define.UserGetListOutput, list []define.ContentGetListOutputItem, err error) {
	output = &define.UserGetListOutput{}
	list = make([]define.ContentGetListOutputItem, 0)

	if input.Type != "message" {
		var where map[string]interface{}
		var order string
		if input.CategoryId == 0 {
			// 默认查询全部
			where = builder.Eq{"content.type": input.Type}
		} else {
			// 查询指定栏目
			where = builder.Eq{"content.type": input.Type, "content.category_id": input.CategoryId}
		}
		// 排序方式
		switch input.Sort {
		case model.ContentSortNewly:
			// 按照创建时间
			order = "id"
		case model.ContentSortActive:
			// 按照更新时间
			order = "updated_time"
		case model.ContentSortHot:
			// 按照浏览量
			order = "view_count"
		default:
			// 默认，安装创建时间查询
			order = "content.id"
		}
		// 查询内容列表
		if err := engine.Table("content").
			Where(where).
			Where(builder.Eq{"content.uid": input.Uid}).
			Desc(order).
			Join("INNER", "user", "user.id = content.uid").
			Join("INNER", "category", "category.id = content.category_id").
			Find(&list); err != nil {
			return output, list, err
		}
	}

	// 查询用户信息
	_, err = engine.ID(input.Uid).Get(&output.User)
	if err != nil {
		return output, list, err
	}
	// 查询统计信息
	output.Stats, err = s.GetUserStats(input.Uid)

	return
}

// 获取文章数量
func (s *userService) GetUserStats(userId uint) (map[string]int64, error) {
	content := new(model.Content)
	statsMap := make(map[string]int64)
	statsMap["topic"], _ = engine.Where(builder.Eq{"uid": userId, "type": "topic"}).Count(content)
	statsMap["ask"], _ = engine.Where(builder.Eq{"uid": userId, "type": "ask"}).Count(content)
	statsMap["article"], _ = engine.Where(builder.Eq{"uid": userId, "type": "article"}).Count(content)
	statsMap["message"], _ = engine.Where(builder.Eq{"uid": userId}).Count(new(model.Reply))

	return statsMap, nil
}

// 回复列表
func (s userService) GetMessageList(input define.UserGetListInput) (list []define.ReplyGetListOutput, err error) {
	list = make([]define.ReplyGetListOutput, 0)
	err = engine.Table("reply").
		Where(builder.Eq{"reply.uid": input.Uid}).
		Desc("id").
		Join("INNER", "user", "user.id = reply.uid").
		Find(&list)

	return
}

// 修改基本资料
func (s userService) UpdateProfile(c *gin.Context, input define.UserUpdateProfileInput)  error {
	user := new(model.User)
	user.Nickname = input.Nickname
	user.Gender = input.Gender
	_, err := engine.ID(input.Id).Cols("gender,nickname").Update(user)
	if err != nil {
		return err
	}
	// 更新缓存
	if err := Session.UpUser(c, input.Id); err != nil {
		return err
	}
	return nil
}

// 上传头像
func (s userService) UploadAvatar(c *gin.Context, input string)  error {
	uid := Session.GetUid(c)
	user := new(model.User)
	user.Avatar = input
	_, err := engine.ID(uid).Cols("avatar").Update(user)
	if err != nil {
		return err
	}
	// 更新缓存
	if err := Session.UpUser(c, uid); err != nil {
		return err
	}

	return nil
}

// UpdatePassword 修改密码
func (s userService) UpdatePassword(c *gin.Context, input define.UserPasswordInput)  error {
	// 查询密码是否正确
	input.OldPassword = s.EncryptPassword(input.OldPassword)
	has, err := engine.Table("user").Where(builder.Eq{"id": input.Id, "password": input.OldPassword}).Exist()
	fmt.Println("input:", input)
	if err != nil {
		return err
	}
	if has == false {
		return gerroraa.New("原始密码错误")
	}
	// 修改密码
	user := new(model.User)
	user.Password = s.EncryptPassword(input.NewPassword)
	_, err = engine.ID(input.Id).Cols("password").Update(user)
	if err != nil {
		return err
	}

	return nil
}