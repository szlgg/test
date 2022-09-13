package service

import (
	"github.com/gin-gonic/gin"
	"test/app/model"
	"test/app/system/index/internal/define"
	"test/library/gerroraa"
	"xorm.io/builder"
)

// 评论/回复管理服务
var Reply = replyService{}

type replyService struct{}

// GetList 获取回复列表
func (s *replyService) GetList(id int) (output []define.ReplyGetListOutput, err error) {
	output = make([]define.ReplyGetListOutput, 0)
	err = engine.Table("reply").
		Where(builder.Eq{"reply.content_id": id}).
		Desc("id").
		Join("INNER", "user", "user.id = reply.uid").
		Find(&output)

	return
}

// Create 创建回复
func (s *replyService) Create(c *gin.Context, input define.ReplyCreateInput) error {
	user := Session.GetUser(c)
	if user == nil {
		return gerroraa.Newf("未登录或会话已过期，请您登录后再继续")
	}

	reply := new(model.Reply)
	reply = &model.Reply{
		Uid:        int(user.Id),
		ParentId:   input.ParentId,
		TargetType: input.TargetType,
		ContentId:  input.ContentId,
		Content:    input.Content,
	}

	_, err := engine.Insert(reply)
	if err != nil {
		return err
	}

	return nil
}

// Delete 删除回复(硬删除)
func (s *replyService) Delete(id uint) error {
	reply := new(model.Reply)
	_, err := engine.ID(id).Delete(reply)
	if err != nil {
		return err
	}

	return nil
}
