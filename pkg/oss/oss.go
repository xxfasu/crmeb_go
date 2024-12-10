package oss

import (
	"context"
	"crmeb_go/consts"
	"crmeb_go/internal/conf"
	redis2 "crmeb_go/internal/redis"
	"crmeb_go/pkg/logs"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"sync"
	"time"
)

type Client interface {
	UploadFile(objectName string, filePath string) error
	Download(objectKey, downloadPath string) error
	Delete(srcPath string) error
	BatchDelete(srcPath []string) error
	Rename(srcPath, destPath string) error
	Copy(srcPath, destPath string) error
	GetOnly(path string) (string, error)
	SignURL(ossFilePath, fileName string, expiresInSec int64, category int) (string, error)
	IsExist(filePath string) (bool, error)
	GetFileSize(filePath string) (int64, error)
}

var ossClient Client
var syncOSSClient Client

func InitOSS() error {
	result, err := redis2.Client.Get(context.Background(), consts.OSSSTSTokenKey).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			logs.Log.Error("redis err:", zap.Error(err))
			return err
		}
	}
	if len(result) == 0 {
		jsonData, err := GenerateSTSToken()
		if err != nil {
			logs.Log.Error("GenerateSTSToken err: ", zap.Error(err))
			return err
		}
		redis2.Client.Set(context.Background(), consts.OSSSTSTokenKey, string(jsonData), time.Hour)
	}
	registerFactory(consts.AliyunOSS, &AliyunOSSFactory{})
	return nil
}

var once sync.Once

func GetOSSClient() (Client, error) {
	var flag error
	once.Do(func() {
		if len(conf.Config.OSS.Provider) == 0 {
			conf.Config.OSS.Provider = consts.AliyunOSS
		}
		factory, ok := getFactory(conf.Config.OSS.Provider)
		if !ok {
			logs.Log.Error("oss provider not found")
			flag = errors.New("oss provider not found")
			return
		}
		client, err := factory.Create()
		if err != nil {
			panic(err)
		}
		ossClient = client
	})
	if flag != nil {
		return nil, flag
	}
	if ossClient == nil {
		return nil, errors.New("oss client is nil")
	}
	return ossClient, nil
}

func GenerateSTSToken() ([]byte, error) {
	client, err := sts.NewClientWithAccessKey(conf.Config.AliyunOSS.Region, conf.Config.AliyunOSS.KeyID, conf.Config.AliyunOSS.KeySecret)
	if err != nil {
		logs.Log.Error("Error: ", zap.Error(err))
	}
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"
	request.RoleArn = conf.Config.AliyunOSS.RoleArn
	request.RoleSessionName = uuid.NewString()
	request.DurationSeconds = consts.STSTokenDurationSeconds

	// 4. 获取临时凭证
	response, err := client.AssumeRole(request)
	if err != nil {
		fmt.Println("Error getting STS token:", err)
		return nil, err
	}

	var ossInfo = struct {
		OSSAccessKeyId     string `json:"access_key_id"`
		OSSAccessKeySecret string `json:"access_key_secret"`
		OSSBucket          string `json:"bucket"`
		OSSEndpoint        string `json:"endpoint"`
		OSSRegion          string `json:"region"`
		OSSSecurityToken   string `json:"security_token"`
	}{
		OSSAccessKeyId:     response.Credentials.AccessKeyId,
		OSSAccessKeySecret: response.Credentials.AccessKeySecret,
		OSSSecurityToken:   response.Credentials.SecurityToken,
		OSSBucket:          conf.Config.AliyunOSS.Bucket,
		OSSEndpoint:        conf.Config.AliyunOSS.Endpoint,
		OSSRegion:          conf.Config.AliyunOSS.Region,
	}
	logs.Log.Info("GenerateSTSToken OSSConfig:", zap.Any("ossInfo", ossInfo))
	// 将结构体转换为JSON格式的字节切片
	jsonData, err := json.Marshal(ossInfo)
	if err != nil {
		logs.Log.Error("json Marshal err: ", zap.Error(err))
		return nil, err
	}
	return jsonData, nil
}
