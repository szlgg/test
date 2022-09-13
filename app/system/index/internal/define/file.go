package define

import (
	"mime/multipart"
)

// 上传文件返回结果
type FileUploadRes struct {
	Name string // 文件名称
	Url  string // 访问URL，可能只是URI
}

// 上传文件输入参数
type FileUploadInput struct {
	File       []*multipart.FileHeader // 上传文件对象
	Name       string            				  // 自定义文件名称
	RandomName bool              				  // 是否随机命名文件
}

// 上传文件返回参数
type FileUploadOutput struct {
	Id   uint   // 数据表ID
	Name string // 文件名称
	Path string // 本地路径
	Url  string // 访问URL，可能只是URI
}