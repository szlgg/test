package service

import (
	"fmt"
	"strings"
	"test/app/model"
	"test/app/system/index/internal/define"
	"test/config"
	"test/library/ghtml"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"xorm.io/builder"
)

// ContentService 内容管理服务
var Content = contentService{}

type contentService struct{}

// GetTitle 获取标题
func (s *contentService) GetTitle(id int) string {
	content := new(model.Content)
	_, err := engine.ID(id).Get(content)
	if err != nil {
		return ""
	}
	return content.Title
}

// 增加浏览量
func (s *contentService) AddViewCount(id int, count int) (view_count int, err error) {
	content := model.Content{ViewCount: count + 1}

	_, err = engine.ID(id).Cols("view_count, updated_time").Update(&content)
	if err != nil {
		return 0, err
	}
	return content.ViewCount, nil
}

// Create 新增内容
func (s *contentService) Create(c *gin.Context, input define.ContentCreateInput) (output define.ContentCreateOutput, err error) {
	// 获取用户ID
	if input.Uid == 0 {
		session := sessions.DefaultMany(c, "Login")
		input.Uid = session.Get(SessionKeyUser).(model.User).Id
	}
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(input); err != nil {
		return output, err
	}
	// 插入数据
	content := model.Content{
		Type:       input.Type,
		Uid:        int(input.Uid),
		CategoryId: int(input.CategoryId),
		Title:      input.Title,
		Content:    input.Content,
	}

	_, err = engine.Insert(&content)
	if err != nil {
		return output, err
	}
	return define.ContentCreateOutput{ContentId: uint(content.Id)}, err
}

// GetList 获取列表
func (s *contentService) GetList(input define.ContentGetListInput) (output *define.ContentGetListOutput, list []define.ContentGetListOutputItem, err error) {
	output = &define.ContentGetListOutput{
		Page: input.Page,
		// Size: input.Size,
	}

	input.Page--

	// 默认查询topic
	if input.Type == "" {
		input.Type = model.ContentTypeTopic
	}

	// 执行查询
	list = make([]define.ContentGetListOutputItem, 0)
	content := new(model.Content)
	var where map[string]interface{}
	var order string
	if input.CategoryId == 0 {
		// 默认查询全部
		where = builder.Eq{"content.type": input.Type}
	} else {
		// 查询指定栏目
		where = builder.Eq{"content.type": input.Type, "content.category_id": input.CategoryId}
	}

	// 排序方式
	switch input.Sort {
	case model.ContentSortNewly:
		// 按照创建时间
		order = "id"
	case model.ContentSortActive:
		// 按照更新时间
		order = "updated_time"
	case model.ContentSortHot:
		// 按照浏览量
		order = "view_count"
	default:
		// 默认，安装创建时间查询
		order = "content.id"
	}
	// 查询数据
	if err := engine.Table("content").
		Where(where).
		Desc(order).
		Limit(config.Conf.Index.Num, input.Page*config.Conf.Index.Num).
		Join("INNER", "user", "user.id = content.uid").
		Join("INNER", "category", "category.id = content.category_id").
		Find(&list); err != nil {
		return output, list, err
	}
	// 获取数据总数
	num, _ := engine.Where(where).Count(content)
	output.Total = int(num)

	// 没有数据
	if len(list) == 0 {
		return output, list, nil
	}

	return
}

// GetDetail 查询详情
func (s *contentService) GetDetail(id int) (output *define.ContentGetDetailOutput, err error) {
	output = &define.ContentGetDetailOutput{}
	// 查询内容
	result, err := engine.ID(id).Get(&output.Content)
	if err != nil {
		return output, err
	}
	if !result {
		return nil, nil
	}

	content := new(model.Content)
	// 查询回复数量
	reply := new(model.Reply)
	total, err := engine.Where(builder.Eq{"content_id": id}).Count(reply)
	// 查询点赞数量
	interact := new(model.Interact)
	zanTotal, err := engine.
		Where(builder.Eq{
			"type":        0,
			"target_type": "content",
			"target_id":   id,
		}).
		Count(interact)
	// 更新回复数量与点赞数量
	if output.Content.ReplyCount != int(total) || output.Content.ZanCount != int(zanTotal) {
		output.Content.ZanCount = int(zanTotal)
		output.Content.ReplyCount = int(total)
		content.ReplyCount = int(total)
		content.ZanCount = int(zanTotal)
		_, err = engine.ID(id).Update(content)
		if err != nil {
			return output, err
		}
	}
	// 查询用户信息
	result, err = engine.ID(output.Content.Uid).Cols("id", "nickname", "avatar").Get(&output.User)
	if err != nil {
		return output, err
	}
	if !result {
		return nil, nil
	}

	return output, nil
}

// Update 更新内容
func (s *contentService) Update(input define.ContentUpdateInput) error {
	// 不允许HTML代码
	if err := ghtml.SpecialCharsMapOrStruct(input); err != nil {
		return err
	}
	// 更新数据
	content := new(model.Content)
	content = &model.Content{
		Type:       input.Type,
		Title:      input.Title,
		CategoryId: int(input.CategoryId),
		Content:    input.Content,
	}
	_, err := engine.ID(input.Id).Update(content)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除内容
func (s *contentService) Delete(id uint) error {
	content := new(model.Content)
	_, err := engine.ID(id).Delete(content)
	if err != nil {
		return err
	}

	return nil
}

// 搜索内容列表
func (s *contentService) Search(input define.SearchInput) (output *define.ContentGetSearchOutput, list []define.SearchGetListOutputItem, err error) {
	output = &define.ContentGetSearchOutput{}
	output.PageData.Page = input.Page

	input.Page--

	list = make([]define.SearchGetListOutputItem, 0)
	content := new(model.Content)
	var where, order, likePattern string
	var whereType map[string]interface{}
	input.Key = strings.ReplaceAll(input.Key, "'", "''")
	likePattern = "%" + input.Key + "%"

	where = fmt.Sprintf(" content.title like '%s' or content.content like '%s' ", likePattern, likePattern)

	// 栏目
	if input.Type != "" {
		whereType = builder.Eq{"content.type": input.Type}
		fmt.Println("whereType:", whereType)
	}

	// 排序方式
	switch input.Sort {
	case model.ContentSortNewly:
		// 按照创建时间
		order = "content.id"
	case model.ContentSortActive:
		// 按照更新时间
		order = "updated_time"
	case model.ContentSortHot:
		// 按照浏览量
		order = "view_count"
	default:
		// 默认，安装创建时间查询
		order = "content.id"
	}
	// 查询数据
	if err := engine.Table("content").
		Where(where).
		And(whereType).
		Desc(order).
		Limit(config.Conf.Index.Num, input.Page*config.Conf.Index.Num).
		Join("INNER", "user", "user.id = content.uid").
		Join("INNER", "category", "category.id = content.category_id").
		Find(&list); err != nil {
		return output, list, err
	}

	// 获取数据总数
	num, _ := engine.Where(fmt.Sprintf(" content like '%s' or title like '%s' ", likePattern, likePattern)).And(whereType).Count(content)
	output.PageData.Total = int(num)

	// 没有数据
	if len(list) == 0 {
		return output, list, nil
	}

	// 查询统计信息
	output.Stats, err = s.GetSearchStats(where)
	if err != nil {
		return output, list, err
	}

	return
}

// 统计搜索数据
func (s contentService) GetSearchStats(where string) (map[string]int64, error) {
	content := new(model.Content)
	statsMap := make(map[string]int64)
	statsMap["all"], _ = engine.Where(where).Count(content)
	statsMap["topic"], _ = engine.Where(where).And(builder.Eq{"type": "topic"}).Count(content)
	statsMap["ask"], _ = engine.Where(where).And(builder.Eq{"type": "ask"}).Count(content)
	statsMap["article"], _ = engine.Where(where).And(builder.Eq{"type": "article"}).Count(content)
	return statsMap, nil
}
