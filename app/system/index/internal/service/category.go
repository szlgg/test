package service

import (
	"test/app/model"
	"time"
	"xorm.io/builder"
)

// 栏目管理服务
var (
	Category = categoryService{}
)

type categoryService struct{}

const (
	mapCacheKey       = "category_map_cache"
	mapCacheDuration  = time.Hour
	treeCacheKey      = "category_tree_cache"
	treeCacheDuration = time.Hour
)

// GetList 获取栏目列表
func (s *categoryService) GetList(contentType string) (map[int64]model.Category,error) {
	//var category []model.Category
	category := make(map[int64]model.Category)
	err := engine.Where(builder.Eq{"content_type":contentType}).Cols("id", "name").Find(&category)
	if err != nil {
		return category, err
	}
	return category, nil
}

func (s *categoryService) GetCategoryName(id int) string {
	content := new(model.Content)
	_, err := engine.ID(id).Get(content)
	if err != nil {
		return ""
	}
	category := new(model.Category)
	_, err = engine.ID(content.CategoryId).Get(category)
	if err != nil {
		return ""
	}

	return category.Name
}