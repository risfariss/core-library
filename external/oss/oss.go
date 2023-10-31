package oss

import (
	"bitbucket.org/kawancicil/core-library/external/oss/request"
	"bytes"
)

type Oss interface {
	UploadFile(bucketConfig *request.OssBucketConfig, fileServerPath string) (string, error)
	DownloadFile(bucketConfig *request.OssBucketConfig, file string) (*bytes.Buffer, error)
}
