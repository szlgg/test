package service

import (
	"fmt"
	"html/template"
	"test/app/model"
	"time"
)

// ContentTitle 根据内容ID获取内容标题
func ContentTitle(id int) string {
	result := Content.GetTitle(id)
	return result
}

// CategoryName 根据内容ID获取栏目名称
func CategoryName(id int) string {
	result := Category.GetCategoryName(id)
	return result
}

// GenderFont 根据性别字段内容返回性别的font。
func GenderFont(gender int) string {
	switch gender {
	case model.UserGenderMale:
		return "icon-male"
		// return "&#xe651;"
	case model.UserGenderFemale:
		return "icon-female"
		// return "&#xe636;"
	default:
		return "icon-genderless"
		// return "&#xead2;"
	}
}

// Gender 根据性别字段内容返回性别。
func Gender(gender int) string {
	switch gender {
	case model.UserGenderMale:
		return "男"
	case model.UserGenderFemale:
		return "女"
	default:
		return "未知"
	}
}

// DidIZan 我是否赞了这个内容
func DidIZan(targetType string, targetId int, uid uint) bool {
	result := Interact.DidIZan(targetType, targetId, int(uid))
	return result
}

// DidICai 我是否踩了这个内容
func DidICai(targetType string, targetId int, uid uint) bool {
	result := Interact.DidICai(targetType, targetId, int(uid))
	return result
}

// HTML 字符串转html
func HTML(h string) interface{} {
	return template.HTML(h)
}

// FormatTime 格式化时间
func FormatTime(t int64) string {
	if t == 0 {
		return ""
	}
	n := time.Now().Unix()

	var ys int64 = 31536000
	var ds int64 = 86400
	var hs int64 = 3600
	var ms int64 = 60
	var ss int64 = 1

	var rs string

	d := n - t
	switch {
	case d > ys:
		rs = fmt.Sprintf("%d年前", int(d/ys))
	case d > ds:
		rs = fmt.Sprintf("%d天前", int(d/ds))
	case d > hs:
		rs = fmt.Sprintf("%d小时前", int(d/hs))
	case d > ms:
		rs = fmt.Sprintf("%d分钟前", int(d/ms))
	case d > ss:
		rs = fmt.Sprintf("%d秒前", int(d/ss))
	default:
		rs = "刚刚"
	}

	return rs
}
