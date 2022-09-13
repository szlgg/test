package model

const (
	ContentListDefaultSize = 10
	ContentListMaxSize     = 50
	ContentSortDefault     = 0 // 排序：默认，按照创建时间
	ContentSortNewly       = 1 // 排序：按照创建时间
	ContentSortActive      = 2 // 排序：按照更新时间
	ContentSortHot         = 3 // 排序：按照浏览量
	ContentSortScore       = 3 // 排序：按照搜索结果关联性
	ContentTypeArticle     = "article"
	ContentTypeAsk         = "ask"
	ContentTypeTopic       = "topic"
)

// ==========================================================================================
// 通用数据结构
// ==========================================================================================
// 主要用于列表展示
type ContentListItem struct {
	Id          uint   `json:"id"`           // 自增ID
	CategoryId  uint   `json:"category_id"`  // 栏目ID
	Uid         uint   `json:"uid"`          // 用户ID
	Title       string `json:"title"`        // 标题
	Sort        uint   `json:"sort"`         // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	Brief       string `json:"brief"`        // 摘要
	Thumb       string `json:"thumb"`        // 缩略图
	Tags        string `json:"tags"`         // 标签名称列表，以JSON存储
	Referer     string `json:"referer"`      // 内容来源，例如github/gitee
	Status      uint   `json:"status"`       // 状态 0: 正常, 1: 禁用
	ViewCount   uint   `json:"view_count"`   // 浏览数量
	ReplyCount  uint   `json:"reply_count"`  // 回复数量
	ZanCount    uint   `json:"zan_count"`    // 赞
	CaiCount    uint   `json:"cai_count"`    // 踩
	CreatedTime int64  `json:"created_time"` // 创建时间
	UpdatedTime int64  `json:"updated_time"` // 修改时间
	// CreatedAt  *gtime.Time `json:"created_at"`  // 创建时间
	// UpdatedAt  *gtime.Time `json:"updated_at"`  // 修改时间
}

// 绑定到Content列表中的栏目信息
type ContentListCategoryItem struct {
	Id          uint   `json:"id"`           // 分类ID，自增主键
	Name        string `json:"name"`         // 分类名称
	Thumb       string `json:"thumb"`        // 封面图
	ContentType string `json:"content_type"` // 内容类型：content, ask, article, reply

}

// 绑定到Content列表中的用户信息
type ContentListUserItem struct {
	Id       uint   `json:"id"`       // UID
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像地址
}

// // 绑定到Content列表中的回复信息
// type ContentListReplyItem struct {
// 	Id       uint   `json:"id"`
// 	Content  string `json:"content"`
// 	ParentId string `json:"parent_id"`
// }

func (ContentListItem) TableName() string {
	return "user"
}

// // Content详情
// type ContentDetail struct {
//	Content Content
//	User    User
// }
