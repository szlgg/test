package model

// 用户表
type User struct {
	Id          uint   `xorm:"not null pk autoincr integer"  json:"id"`           // 自增ID
	Passport    string `xorm:"not null default '' TEXT"      json:"passport"`     // 账号
	Password    string `xorm:"not null default '' TEXT"      json:"password"`     // MD5密码
	Nickname    string `xorm:"not null default '' TEXT"      json:"nickname"`     // 昵称
	Avatar      string `xorm:"not null default '' TEXT"      json:"avatar"`       // 头像地址
	Gender      int    `xorm:"not null default 0 integer"    json:"gender"`       // 性别 0:未设置 1:男 2:女
	CreatedTime int64  `xorm:"created "                      json:"created_time"` // 注册时间
	UpdatedTime int64  `xorm:"updated"                       json:"updated_time"` // 更新时间
}

// 内容列表
type Content struct {
	Id             int    `xorm:"not null pk autoincr integer" json:"id"`               // 自增ID
	Type           string `xorm:"not null default '' text"     json:"type"`             // 内容模型: topic, ask, article等，具体由程序定义
	Uid            int    `xorm:"not null default '' integer"  json:"uid"`              // 用户ID
	AdoptedReplyId int    `xorm:"not null default '' INTEGER"  json:"adopted_reply_id"` // 采纳的回复ID，问答模块有效
	CategoryId     int    `xorm:"not null default '' INTEGER"  json:"category_id"`      // 栏目ID
	Title          string `xorm:"not null default '' TEXT"     json:"title"`            // 标题
	Content        string `xorm:"not null default '' TEXT"     json:"content"`          // 内容
	Sort           int    `xorm:"not null default '' integer"  json:"sort"`             // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	Brief          string `xorm:"not null default '' TEXT"     json:"brief"`            // 摘要
	Thumb          string `xorm:"not null default '' TEXT"     json:"thumb"`            // 缩略图
	Tags           string `xorm:"not null default '' TEXT"     json:"tags"`             // 标签名称列表，以JSON存储
	Referer        string `xorm:"not null default '' TEXT"     json:"referer"`          // 内容来源，例如github/gitee
	Status         int    `xorm:"not null default 0 integer"   json:"status"`           // 状态 0: 正常, 1: 禁用
	ReplyCount     int    `xorm:"not null default 0 integer"   json:"reply_count"`      // 回复数量
	ViewCount      int    `xorm:"not null default 0 integer"   json:"view_count"`       // 浏览数量
	ZanCount       int    `xorm:"not null default 0 integer"   json:"zan_count"`        // 赞
	CaiCount       int    `xorm:"not null default 0 integer"   json:"cai_count"`        // 踩
	CreatedTime    int64  `xorm:"created"                      json:"created_time"`     // 注册时间
	UpdatedTime    int64  `xorm:"updated"                      json:"updated_time"`     // 更新时间
}

// 回复列表
type Reply struct {
	Id          int    `xorm:"not null pk autoincr integer" json:"id"`           // 回复ID
	ParentId    int    `xorm:"not null default 0 integer"   json:"parent_id"`    // 回复对应的上一级回复ID(没有的话默认为0)
	Title       string `xorm:"not null default '' TEXT"     json:"title"`        // 回复标题
	Content     string `xorm:"not null default '' TEXT"     json:"content"`      // 回复内容
	TargetType  string `xorm:"not null default '' TEXT"     json:"target_type"`  // 内容类型
	ContentId   int    `xorm:"not null default '' integer"  json:"content_id"`   // 对应内容ID
	Uid         int    `xorm:"not null default '' integer"  json:"uid"`          // 用户ID
	ZanCount    int    `xorm:"not null default 0 integer"   json:"zan_count"`    // 赞
	CaiCount    int    `xorm:"not null default 0 integer"   json:"cai_count"`    // 踩
	CreatedTime int64  `xorm:"created"                      json:"created_time"` // 注册时间
	UpdatedTime int64  `xorm:"updated"                      json:"updated_time"` // 更新时间
}

// 交互管理表
type Interact struct {
	Id          int    `xorm:"not null pk autoincr INTEGER" json:"id"`           // 自增ID
	Type        int    `xorm:"not null default '' integer"  json:"type"`         // 操作类型。0:赞，1:踩。
	Uid         int    `xorm:"not null default '' integer"  json:"uid"`          // 操作用户
	TargetType  string `xorm:"not null default '' TEXT"     json:"target_type"`  // 内容模型: content, reply, 具体由程序定义
	TargetId    int    `xorm:"not null default '' integer"  json:"target_id"`    // 对应内容ID，该内容可能是content, reply
	Count       int    `xorm:"not null default '' INTEGER"  json:"count"`        // 操作数据值
	CreatedTime int    `xorm:"created"                      json:"created_time"` // 注册时间
	UpdatedTime int    `xorm:"updated"                      json:"updated_time"` // 更新时间
}

// 栏目表
type Category struct {
	Id          int    `xorm:"not null pk autoincr INTEGER" json:"id"`           // 自增ID
	ContentType string `xorm:"not null default '' TEXT"     json:"content_type"` // 内容类型：topic, ask, article, reply
	Name        string `xorm:"not null default '' TEXT"     json:"name"`         // 分类名称
	Sort        int    `xorm:"not null default 10 integer"  json:"sort"`         // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	CreatedTime int    `xorm:"default NULL integer"         json:"created_time"` // 创建时间
	UpdatedTime int    `xorm:"default NULL integer"         json:"updated_time"` // 修改时间
}

// 文件表
type File struct {
	Id          int    `xorm:"not null pk autoincr INTEGER" json:"id"`           // 自增ID
	Uid         int    `xorm:"not null default '' integer"  json:"uid"`          // 操作用户
	Name        string `xorm:"not null default '' TEXT"     json:"name"`         // 文件名称
	Src         string `xorm:"not null default '' TEXT"     json:"src"`          // 本地存储路径
	Url         string `xorm:"not null default '' TEXT"     json:"url"`          // url地址
	CreatedTime int    `xorm:"default NULL integer"         json:"created_time"` // 创建时间
}
