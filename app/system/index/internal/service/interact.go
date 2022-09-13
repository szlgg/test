package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"test/app/model"
	"test/library/gerroraa"
	"xorm.io/builder"
)

// 交互管理服务
var Interact = interactService{}

type interactService struct{}

const (
	contextMapKeyForMyInteractList = "ContextMapKeyForMyInteractList"
)

// DidIZan 我是否有对指定内容赞
func (s *interactService) DidIZan(targetType string, targetId int, uid int) bool {
	has, _ := engine.Where(builder.Eq{"uid": uid, "target_type": targetType, "target_id": targetId, "type": model.InteractTypeZan}).Exist(new(model.Interact))
	return has
}

// DidICai 我是否有对指定内容踩
func (s *interactService) DidICai(targetType string, targetId int, uid int) bool {
	has, _ := engine.Where(builder.Eq{"uid": uid, "target_type": targetType, "target_id": targetId, "type": model.InteractTypeCai}).Exist(new(model.Interact))
	return has
}

// Zan 赞
func (s *interactService) Zan(c *gin.Context, targetType string, targetId int) error {
	// 获取登录用户信息
	user := Session.GetUser(c)

	interact := new(model.Interact)
	interact.TargetType = targetType
	interact.TargetId = targetId
	interact.Uid = int(user.Id)
	interact.Type = model.InteractTypeZan

	where := builder.Eq{
		"uid": interact.Uid,
		"target_type": interact.TargetType,
		"target_id": interact.TargetId,
		"type": interact.Type,
	}
	if has, _ := engine.Where(where).Exist(interact); has {
		return gerroraa.New("你已经赞过啦")
	}
	_, err := engine.Insert(interact)
	if err != nil {
		return err
	}

	return s.updateCount(model.InteractTypeZan, targetType, targetId)
}

// CancelZan 取消赞
func (s *interactService) CancelZan(c *gin.Context, targetType string, targetId int) error {
	// 获取登录用户信息
	user := Session.GetUser(c)

	interact := new(model.Interact)
	interact.TargetType = targetType
	interact.TargetId = targetId
	interact.Uid = int(user.Id)
	interact.Type = model.InteractTypeZan

	where := builder.Eq{
		"uid": interact.Uid,
		"target_type": interact.TargetType,
		"target_id": interact.TargetId,
		"type": interact.Type,
	}
	_, err := engine.Where(where).Delete(interact)
	if err != nil {
		return err
	}

	return s.updateCount(model.InteractTypeZan, targetType, targetId)
}

// Cai 踩
func (s *interactService) Cai(c *gin.Context, targetType string, targetId int) error {
	// 获取登录用户信息
	user := Session.GetUser(c)

	interact := new(model.Interact)
	interact.TargetType = targetType
	interact.TargetId = targetId
	interact.Uid = int(user.Id)
	interact.Type = model.InteractTypeCai

	where := builder.Eq{
		"uid": interact.Uid,
		"target_type": interact.TargetType,
		"target_id": interact.TargetId,
		"type": interact.Type,
	}
	if has, _ := engine.Where(where).Exist(interact); has {
		return gerroraa.New("你已经踩过啦")
	}
	_, err := engine.Insert(interact)
	if err != nil {
		return err
	}

	return s.updateCount(model.InteractTypeCai, targetType, targetId)
}

// CancelCai 取消赞
func (s *interactService) CancelCai(c *gin.Context, targetType string, targetId int) error {
	// 获取登录用户信息
	user := Session.GetUser(c)

	interact := new(model.Interact)
	interact.TargetType = targetType
	interact.TargetId = targetId
	interact.Uid = int(user.Id)
	interact.Type = model.InteractTypeCai

	where := builder.Eq{
		"uid": interact.Uid,
		"target_type": interact.TargetType,
		"target_id": interact.TargetId,
		"type": interact.Type,
	}
	_, err := engine.Where(where).Delete(interact)
	if err != nil {
		return err
	}

	return s.updateCount(model.InteractTypeCai, targetType, targetId)
}

// 根据业务类型更新指定模块的赞/踩数量
func (s *interactService) updateCount(interactType int, targetType string, targetId int) error {
	interact := new(model.Interact)

	var err error
	switch targetType {
	// 内容赞踩
	case model.InteractTargetTypeContent:
		content := new(model.Content)
		switch interactType {
		case model.InteractTypeZan:
			count, _ := engine.Where(builder.Eq{"target_type": targetType, "target_id": targetId, "type": interactType}).Count(interact)
			content.ZanCount = int(count)
			_, err = engine.ID(targetId).Cols("zan_count").Update(content)
			if err != nil {
				return err
			}
		case model.InteractTypeCai:
			count, _ := engine.Where(builder.Eq{"target_type": targetType, "target_id": targetId, "type": interactType}).Count(interact)
			content.CaiCount = int(count)
			_, err = engine.ID(targetId).Cols("cai_count").Update(content)
			fmt.Println("content:", content.CaiCount)
			fmt.Println("err:", err)
			if err != nil {
				return err
			}
		}
	// 评论赞踩
	case model.InteractTargetTypeReply:
		reply := new(model.Reply)
		switch interactType {
		case model.InteractTypeZan:
			count, _ := engine.Where(builder.Eq{"target_type": targetType, "target_id": targetId, "type": interactType}).Count(interact)
			reply.ZanCount = int(count)
			_, err = engine.ID(targetId).Cols("zan_count").Update(reply)
			fmt.Println("targetId：", targetId)
			fmt.Println("reply：", reply)
			fmt.Println("err：", err)
			if err != nil {
				return err
			}
		case model.InteractTypeCai:
			count, _ := engine.Where(builder.Eq{"target_type": targetType, "target_id": targetId, "type": interactType}).Count(interact)
			reply.CaiCount = int(count)
			_, err = engine.ID(targetId).Cols("cai_count").Update(reply)
			if err != nil {
				return err
			}
		}

	}
	return nil
}
