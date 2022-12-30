package upload

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/geekr-dev/go-blog-app/global"
	"github.com/geekr-dev/go-blog-app/pkg/util"
)

type FileType int

const TypeImage FileType = iota + 1

// 获取文件名称（基于md5生成新的文件名称）
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

// 获取文件后缀扩展
func GetFileExt(name string) string {
	return path.Ext(name)
}

// 获取文件保存路径
func GetSavePath() string {
	return global.AppConfig.UploadSavePath
}

// 检查文件路径是否存在
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

// 检查文件是否是允许上传的后缀类型
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppConfig.UploadImageAllowExts {
			if strings.EqualFold(allowExt, ext) {
				return true
			}
		}
	}
	return false
}

// 检查上传文件尺寸
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppConfig.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

// 检查文件操作权限
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

// 创建保存文件路径
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

// 保存上传文件到指定路径
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
