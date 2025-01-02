package oss

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
