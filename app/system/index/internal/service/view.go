package service

import (
	"fmt"
	"test/app/model"
)

// 视图管理服务
var View = viewService{}

type viewService struct{}

// GetBreadCrumb 获取内容面包屑
func (s *viewService) GetBreadCrumb(Type string, CategoryId int) []model.ViewBreadCrumb {
	breadcrumb := []model.ViewBreadCrumb{
		{Name: "首页", Url: "/"},
	}

	var name string
	var url string

	switch Type {
	case "topic":
		name = "主题"
		url = "/topic"
	case "ask":
		name ="问答"
		url = "/ask"
	case "article":
		name ="文章"
		url = "/article"
	}
	breadcrumb = append(breadcrumb, model.ViewBreadCrumb{
		Name: name,
		Url:  url,
	})

	category := &model.Category{Id: CategoryId}
	_, err := engine.Cols("id", "name").Get(category)
	if err != nil {
		panic(err)
	}
	cid := fmt.Sprintf("%d", category.Id)
	breadcrumb = append(breadcrumb, model.ViewBreadCrumb{
		Name: category.Name,
		Url:  url + "&cate=" + cid,
	})

	return breadcrumb
}