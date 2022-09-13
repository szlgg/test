package model

//type User struct {
//	Id          int    `xorm:"not null pk autoincr integer"`
//	Passport    string `xorm:"not null default '' TEXT"`
//	Password    string `xorm:"not null default '' TEXT"`
//	Nickname    string `xorm:"not null default '' TEXT"`
//	Avatar      string `xorm:"not null default '' TEXT"`
//	Gender      int    `xorm:"not null default 0 integer"`
//	CreatedTime int    `xorm:"default NULL integer"`
//	UpdatedTime int    `xorm:"default NULL integer"`
//}
//
//type Content struct {
//	Id             int    `xorm:"not null pk autoincr integer"`
//	Type           string `xorm:"not null default '' text"`
//	Uid            int    `xorm:"not null default '' integer"`
//	AdoptedReplyId int    `xorm:"not null default '' INTEGER"`
//	Title          string `xorm:"not null default '' TEXT"`
//	Content        string `xorm:"not null default '' TEXT"`
//	Sort           int    `xorm:"not null default '' integer"`
//	Brief          string `xorm:"not null default '' TEXT"`
//	Thumb          string `xorm:"not null default '' TEXT"`
//	Tags           string `xorm:"not null default '' TEXT"`
//	Referer        string `xorm:"not null default '' TEXT"`
//	Status         int    `xorm:"not null default 0 integer"`
//	ReplyCount     int    `xorm:"not null default 0 integer"`
//	ViewCount      int    `xorm:"not null default 0 integer"`
//	ZanCount       int    `xorm:"not null default 0 integer"`
//	CaiCount       int    `xorm:"not null default 0 integer"`
//	CreatedTime    int    `xorm:"default NULL integer"`
//	UpdatedTime    int    `xorm:"default NULL integer"`
//}
//
//type Reply struct {
//	Id          int    `xorm:"not null pk autoincr integer"`
//	ParentId    int    `xorm:"not null default 0 integer"`
//	Title       string `xorm:"not null default '' TEXT"`
//	Content     string `xorm:"not null default '' TEXT"`
//	TargetType  string `xorm:"not null default '' TEXT"`
//	TargetId    int    `xorm:"not null default '' integer"`
//	Uid         int    `xorm:"not null default '' integer"`
//	ZanCount    int    `xorm:"not null default 0 integer"`
//	CaiCount    int    `xorm:"not null default 0 integer"`
//	CreatedTime int    `xorm:"default NULL integer"`
//	UpdatedTime int    `xorm:"default NULL integer"`
//}
//
//type _interactOld20220803 struct {
//	Id int `xorm:"not null pk autoincr INTEGER"`
//}
//
//type Interact struct {
//	Id          int    `xorm:"not null pk autoincr INTEGER"`
//	Type        int    `xorm:"not null default '' integer"`
//	Uid         int    `xorm:"not null default '' integer"`
//	TargetType  string `xorm:"not null default '' TEXT"`
//	TargetId    int    `xorm:"not null default '' integer"`
//	Count       int    `xorm:"not null default '' INTEGER"`
//	CreatedTime int    `xorm:"default NULL integer"`
//	UpdatedTime int    `xorm:"default NULL integer"`
//}
//
//type Category struct {
//	Id          int    `xorm:"not null pk autoincr INTEGER"`
//	ContentType string `xorm:"not null default '' TEXT"`
//	Name        string `xorm:"not null default '' TEXT"`
//	Sort        int    `xorm:"not null default 10 integer"`
//	CreatedTime int    `xorm:"default NULL integer"`
//	UpdatedTime int    `xorm:"default NULL integer"`
//}
