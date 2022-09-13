package model

// 视图面包屑结构
type ViewBreadCrumb struct {
	Name string // 显示名称
	Url  string // 链接地址，当为空时表示被选中
}
