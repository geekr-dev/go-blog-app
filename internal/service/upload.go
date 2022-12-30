package service

import (
	"errors"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/geekr-dev/go-blog-app/global"
	"github.com/geekr-dev/go-blog-app/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, header *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(header.Filename)
	savePath := upload.GetSavePath()
	dst := filepath.Join(savePath, fileName)
	// 上传文件不在允许的类型范围内
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	// 保存路径不存在则主动创建
	if upload.CheckSavePath(savePath) {
		if err := upload.CreateSavePath(savePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	// 上传文件尺寸超出最大限制
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}
	// 文件上传保存后权限检查
	if upload.CheckPermission(savePath) {
		return nil, errors.New("insufficient file permission")
	}
	// 保存文件
	if err := upload.SaveFile(header, dst); err != nil {
		return nil, err
	}
	accessUrl := global.AppConfig.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
