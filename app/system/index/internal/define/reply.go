package define

import "test/app/model"

type ReplyDoCreateReq struct {
	ReplyCreateInput
	ParentId   uint   `form:"parent_id"`   // 回复对应的上一级回复ID(没有的话默认为0)
	TargetType string `form:"target_type"` // 评论类型: topic, ask, article, reply
	ContentId  uint   `form:"content_id"`  // 对应内容ID
	Content    string `form:"content"`     // 回复内容
}

// 执行删除内容
type ReplyDoDeleteReq struct {
	Id uint `form:"id"` // 对应内容ID
}

// 创建内容
type ReplyCreateInput struct {
	Title      string
	ParentId   int    `form:"parent_id"`   // 回复对应的上一级回复ID(没有的话默认为0)
	TargetType string `form:"target_type"` // 评论类型: topic, ask, article, reply
	ContentId  int    `form:"content_id"`  // 对应内容ID
	Content    string `form:"content"`     // 回复内容
	Uid        int
}

// 查询回复
type ReplyGetListOutput struct {
	Reply   *model.Reply   `json:"reply"   xorm:"extends"`
	User    *model.User    `json:"user"    xorm:"extends"`
}