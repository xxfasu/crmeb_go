package oss

import (
	"crmeb_go/consts"
	"crmeb_go/internal/redis"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	redis.InitRedis()
	InitOSS()
	GetOSSClient()
	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)
}

func TestUploadFile(t *testing.T) {
	err := ossClient.UploadFile("test/123456.docx", "D:\\goproject\\cloud-backend\\test\\123456.docx")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	t.Logf("成功")
}

func TestDownload(t *testing.T) {
	err := syncOSSClient.Download("test/test.pdf", "D:\\goproject\\cloud-backend\\test\\testaabb.pdf")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

}

func TestGetOnly(t *testing.T) {

	file, err := ossClient.GetOnly("test/1")
	if err != nil {
		t.Error(err)
	}
	t.Log(file)
}

func TestSignURL(t *testing.T) {
	path := "test/"
	fileName := "ttaa.pdf"
	signedURL, err := ossClient.SignURL(path, fileName, consts.UrlExpireTime, consts.PreviewFile)
	if err != nil {
		t.Errorf("Error generating signed URL: %v", err)
	}
	// 输出签名URL
	t.Logf("Signed URL: %s", signedURL)
}

func TestRename(t *testing.T) {
	err := ossClient.Rename("test/test.pdf", "test/test1.pdf")
	if err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	err := ossClient.Delete("test/test.pdf")
	if err != nil {
		t.Error(err)
	}
}

func TestCopy(t *testing.T) {
	err := ossClient.Copy("test/test.pdf", "test/copy.pdf")
	if err != nil {
		t.Error(err)
	}
}
