package oss

import (
	"bitbucket.org/kawancicil/core-library/constant"
	"bitbucket.org/kawancicil/core-library/external/oss/request"
	"bytes"
	"fmt"
	oss_sdk "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"log"
	"os"
	"path/filepath"
)

type OssUtils struct {
	/*todo fill this block if need connection with another interface*/
}

func InitOssUtils() Oss {
	return &OssUtils{
	}
}



func (r *OssUtils)  UploadFile(bucketConfig *request.OssBucketConfig, fileServerPath string) (string, error) {
	// Creates an OSSClient instance.
	client, err := oss_sdk.New(bucketConfig.Endpoint, bucketConfig.AccessKeyID, bucketConfig.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	bucketName := bucketConfig.BucketName
	objectName := filepath.Base(fileServerPath)
	localFilename := fileServerPath
	// Obtains a bucket.
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	chunks, err := oss_sdk.SplitFileByPartNum(localFilename, 1)
	fd, err := os.Open(localFilename)
	defer fd.Close()
	defer os.Remove(localFilename)
	// Step 1: Initiates a multipart upload event.
	imur, err := bucket.InitiateMultipartUpload(objectName)
	// Step 2: Uploads parts.
	var parts []oss_sdk.UploadPart
	for _, chunk := range chunks {
		_, _ = fd.Seek(chunk.Offset, os.SEEK_SET)
		// Calls UploadPart to upload each part.
		part, err := bucket.UploadPart(imur, fd, chunk.Size, chunk.Number)
		if err != nil {
			fmt.Println("Error:", err)
			return "", err
		}
		parts = append(parts, part)
	}
	// Step 3: Completes the multipart upload task.
	cmur, err := bucket.CompleteMultipartUpload(imur, parts)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	return cmur.Location, nil
}

func (r *OssUtils)  DownloadFile(bucketConfig *request.OssBucketConfig, file string) (*bytes.Buffer, error) {
	// Create an OSSClient instance.
	client, err := oss_sdk.New(bucketConfig.Endpoint, bucketConfig.AccessKeyID, bucketConfig.AccessKeySecret)
	if err != nil {
		log.Println(constant.ErrorConnectAliCloud, err.Error())
		return &bytes.Buffer{}, err
	}
	bucketName := bucketConfig.BucketName
	// Obtain the bucket.
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		log.Println(constant.ErrorBucketNameNotFound, err.Error())
		return &bytes.Buffer{}, err
	}
	// Download the object to the local buffer.
	body, err := bucket.GetObject(file)
	if err != nil {
		log.Println(constant.ErrorDownloadFromAliCloud, err.Error())
		return &bytes.Buffer{}, err
	}
	defer body.Close()

	buf := new(bytes.Buffer)
	_, _ = io.Copy(buf, body)
	return buf, nil
}
