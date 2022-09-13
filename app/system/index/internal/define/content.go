package define

import "test/app/model"

type SearchGetListReq struct {
	SearchInput
	Key  string `form:"key"`
	Type string `form:"type"`
	Sort int    `form:"sort"` // 排序类型(0:最新, 默认。1:最新, 2:活跃, 3:热度)

}

// 搜索列表
type SearchInput struct {
	Key  string `form:"key"`
	Type string `form:"type"`
	Sort int    `form:"sort"` // 排序类型(0:最新, 默认。1:最新, 2:活跃, 3:热度)
	Page int    `form:"page"` // 分页号码
	Size int    `form:"size"` // 分页数量，最大50
}

type ContentGetSearchOutput struct {
	PageData ContentGetListOutput
	// SearchGetListOutputItem
	Stats map[string]int64
}

type SearchGetListOutputItem struct {
	Content  *model.ContentListItem         `json:"content"  xorm:"extends"`
	Category *model.ContentListCategoryItem `json:"category" xorm:"extends"`
	User     *model.ContentListUserItem     `json:"user"     xorm:"extends"`
}

// 展示创建内容页面
type ContentCreateReq struct {
	Type string `form:"type"`
}

// 展示修改内容页面
type ContentUpdateReq struct {
	Id uint `form:"id"`
}

// 获取内容列表
type ContentGetListReq struct {
	ContentGetListInput
	CategoryId uint `form:"cate"`            // 栏目ID
	Page       int  `form:"page,default=1"`  // 分页号码
	Size       int  `form:"size,default=10"` // 分页数量，最大50
}

// 查询列表结果
type ContentGetListOutput struct {
	List  []*ContentGetListOutputItem `json:"list"`  // 列表
	Page  int                         `json:"page"`  // 分页码
	Size  int                         `json:"size"`  // 分页数量
	Total int                         `json:"total"` // 数据总数
}

type ContentGetListOutputItem struct {
	Content  *model.ContentListItem         `json:"content"  xorm:"extends"`
	Category *model.ContentListCategoryItem `json:"category" xorm:"extends"`
	User     *model.ContentListUserItem     `json:"user"     xorm:"extends"`
}

// 查询详情结果
type ContentGetDetailOutput struct {
	Content model.Content `json:"content" xorm:"extends"`
	User    model.User    `json:"user"    xorm:"extends"`
}

// 获取内容列表
type ContentGetListInput struct {
	Type       string `form:"type"` // 内容模型
	CategoryId uint   `form:"cate"` // 栏目ID
	Page       int    `form:"page"` // 分页号码
	Size       int    `form:"size"` // 分页数量，最大50
	Sort       int    `form:"sort"` // 排序类型(0:最新, 默认。1:最新, 2:活跃, 3:热度)
	Uid        uint   `form:"uid"`  // 要查询的用户ID
}

// 执行创建内容
type ContentDoCreateReq struct {
	ContentCreateInput
}

// 创建内容
type ContentCreateInput struct {
	ContentCreateUpdateBase
	Uid uint
}

// 执行修改内容
type ContentDoUpdateReq struct {
	ContentUpdateInput
	Id uint `form:"id"`
}

// 修改内容
type ContentUpdateInput struct {
	ContentCreateUpdateBase
	Id uint `form:"id"`
}

// 执行删除内容
type ContentDoDeleteReq struct {
	Id uint `form:"id"` // 对应内容ID
}

// 创建内容返回结果
type ContentCreateOutput struct {
	ContentId uint `json:"content_id"`
}

// 创建/修改内容基类
type ContentCreateUpdateBase struct {
	Type       string   `form:"type"`       // 内容模型
	CategoryId uint     `form:"categoryId"` // 栏目ID
	Title      string   `form:"title"`      // 标题
	Content    string   `form:"content"`    // 内容
	Brief      string   // 摘要
	Thumb      string   // 缩略图
	Tags       []string // 标签名称列表，以JSON存储
	Referer    string   // 内容来源，例如github/gitee
}
