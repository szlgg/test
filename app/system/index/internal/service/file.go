package service

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"test/app/model"
	"test/app/system/index/internal/define"
	"test/config"
	"test/library/gerroraa"
	"test/library/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// 文件管理服务
var File = fileService{}

type fileService struct{}

// 统一上传文件
func (s *fileService) Upload(c *gin.Context, input []*multipart.FileHeader) ([]*define.FileUploadRes, error) {
	files := make([]*model.File, 0)
	uploadResult := make([]*define.FileUploadRes, 0)

	uploadPath := config.Conf.Upload.Path
	if uploadPath == "" {
		return nil, gerroraa.New("上传文件路径配置不存在")
	}

	// 返回当前的时间戳
	now := time.Now()
	// 文件路径
	fileDir := fmt.Sprintf("%s/img/%d%d%d", uploadPath, now.Year(), now.Month(), now.Day())
	// ModePerm是0777，这样拥有该文件夹路径的执行权限
	err := os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	for _, fileObj := range input {
		fileObj.Filename = utils.RandomString() + filepath.Ext(fileObj.Filename)

		// 文件路径
		filePathStr := filepath.Join(fileDir, fileObj.Filename)
		// 将浏览器客户端上传的文件拷贝到本地路径的文件里面
		err := c.SaveUploadedFile(fileObj, filePathStr)
		if err != nil {
			return nil, err
		}

		file := &model.File{
			Name: fileObj.Filename,
			Src:  filePathStr,
			Url:  filePathStr,
		}
		files = append(files, file)
		uploadResult = append(uploadResult, &define.FileUploadRes{
			Name: fileObj.Filename,
			Url:  filePathStr,
		})
	}

	_, err = engine.Insert(files)
	if err != nil {
		return nil, err
	}

	return uploadResult, nil
}

// 上传头像
func (s *fileService) UploadAvatar(c *gin.Context, input *multipart.FileHeader) (string, error) {
	uploadPath := config.Conf.Upload.Path
	if uploadPath == "" {
		return "", gerroraa.New("上传文件路径配置不存在")
	}

	// 文件路径
	fileDir := fmt.Sprintf("%s/avatar", uploadPath)
	// ModePerm是0777，这样拥有该文件夹路径的执行权限
	err := os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		return "", err
	}
	// 文件路径
	filePathStr := filepath.Join(fileDir, input.Filename)
	// 上传文件到指定的路径
	err = c.SaveUploadedFile(input, filePathStr)
	if err != nil {
		return "", err
	}
	filePathStr = "/" + filePathStr

	return filePathStr, err
}
