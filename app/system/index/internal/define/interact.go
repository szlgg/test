package define

// 赞
type InteractZanReq struct {
	Id   uint   `form:"id"`
	Type string `form:"type"`
}

// 取消赞
type InteractCancelZanReq struct {
	Id   uint   `form:"id"`
	Type string `form:"type"`
}

// 踩
type InteractCaiReq struct {
	Id   uint   `form:"id"`
	Type string `form:"type"`
}

// 取消踩
type InteractCancelCaiReq struct {
	Id   uint   `form:"id"`
	Type string `form:"type"`
}
