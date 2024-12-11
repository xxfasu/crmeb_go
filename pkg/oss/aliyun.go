package oss

import (
	"crmeb_go/internal/conf"
	"crmeb_go/pkg/logs"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
	"net/url"
	"strconv"
)

type AliyunOSSClient struct {
	bucket *oss.Bucket
}

type AliyunOSSFactory struct{}

func (f *AliyunOSSFactory) Create() (Client, error) {
	client, err := oss.New(conf.Config.AliyunOSS.Endpoint, conf.Config.AliyunOSS.KeyID, conf.Config.AliyunOSS.KeySecret)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(conf.Config.AliyunOSS.Bucket)
	if err != nil {
		return nil, err
	}
	return &AliyunOSSClient{bucket: bucket}, nil
}

func (c *AliyunOSSClient) UploadFile(objectName string, filePath string) error {
	logs.Log.Info("UploadFile", zap.String("objectName", objectName), zap.String("filePath", filePath))
	err := c.bucket.PutObjectFromFile(objectName, filePath)
	if err != nil {
		return err
	}
	fmt.Printf("File uploaded to Aliyun OSS: %s\n", objectName)
	return nil
}

func (c *AliyunOSSClient) Download(objectKey, downloadPath string) error {
	logs.Log.Info("Download", zap.String("objectKey", objectKey), zap.String("downloadPath", downloadPath))
	if len(objectKey) == 0 || len(downloadPath) == 0 {
		return errors.New("objectKey downloadPath is empty")
	}
	return c.bucket.GetObjectToFile(objectKey, downloadPath)
}

func (c *AliyunOSSClient) Delete(srcPath string) error {
	logs.Log.Info("Delete", zap.String("srcPath", srcPath))
	err := c.bucket.DeleteObject(srcPath)
	if err != nil {
		return err
	}
	return nil
}

func (c *AliyunOSSClient) BatchDelete(srcPath []string) error {
	logs.Log.Info("BatchDelete", zap.Strings("srcPath", srcPath))
	_, err := c.bucket.DeleteObjects(srcPath, oss.DeleteObjectsQuiet(false))
	if err != nil {
		return err
	}
	return nil
}

func (c *AliyunOSSClient) Rename(srcPath, destPath string) error {
	logs.Log.Info("Rename", zap.String("srcPath", srcPath), zap.String("srcPath", destPath))
	_, err := c.bucket.CopyObject(srcPath, destPath)
	if err != nil {
		return err
	}
	err = c.bucket.DeleteObject(srcPath)
	if err != nil {
		return err
	}
	return nil
}

func (c *AliyunOSSClient) Copy(srcPath, destPath string) error {
	logs.Log.Info("Copy ", zap.String("srcPath", srcPath), zap.String("srcPath", destPath))
	_, err := c.bucket.CopyObject(srcPath, destPath)
	return err
}

func (c *AliyunOSSClient) GetOnly(path string) (string, error) {
	logs.Log.Info("GetOnly ", zap.String("path", path))
	marker := oss.Marker("")
	lsRes, err := c.bucket.ListObjects(
		oss.Prefix(path),
		oss.MaxKeys(1),
		marker,
	)
	if err != nil {
		return "", err
	}
	if len(lsRes.Objects) == 0 {
		return "", errors.New("file does not exist")
	}

	return lsRes.Objects[0].Key, nil
}

func (c *AliyunOSSClient) SignURL(ossFilePath, fileName string, expiresInSec int64, category int) (string, error) {
	logs.Log.Info("SignURL", zap.String("ossFilePath", ossFilePath), zap.String("fileName", fileName), zap.Int64("expiresInSec", expiresInSec), zap.Int("category", category))

	options := make([]oss.Option, 0)
	if category == enums.DownloadFile {
		options = append(options, oss.ResponseContentDisposition(fmt.Sprintf("attachment; filename*=utf-8''%s", url.PathEscape(fileName))))
	}
	if category == enums.UploadFile {
		objectURL, err := c.bucket.SignURL(ossFilePath, oss.HTTPPut, expiresInSec)
		return objectURL, err
	}
	objectURL, err := c.bucket.SignURL(ossFilePath, oss.HTTPGet, expiresInSec, options...)
	return objectURL, err
}

func (c *AliyunOSSClient) IsExist(filePath string) (bool, error) {
	logs.Log.Info("IsExist", zap.String("filePath", filePath))
	found, err := c.bucket.IsObjectExist(filePath)
	return found, err
}

func (c *AliyunOSSClient) GetFileSize(filePath string) (int64, error) {
	logs.Log.Info("GetFileSize", zap.String("filePath", filePath))
	// 获取对象元数据，其中包含了文件大小（Content-Length）
	props, err := c.bucket.GetObjectDetailedMeta(filePath)
	if err != nil {
		return 0, err
	}
	// 提取文件大小（单位：字节）
	fileSize := props.Get("Content-Length")
	if len(fileSize) == 0 {
		return 0, errors.New("failed to retrieve file size from metadata")
	}
	result, err := strconv.ParseInt(fileSize, 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
