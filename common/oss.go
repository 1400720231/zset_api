package common

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"os"
)

func UploadFileToOss(fd io.Reader, name string) string {
	// https://1400720231.oss-cn-shenzhen.aliyuncs.com/1.jpg
	scheme := "https://"
	endpoint := "oss-cn-shenzhen.aliyuncs.com"
	accessKeyID := "LTAI5tArgc71UVg8JRTkUar7"
	accessKeySecret := "ZEZHyWv8hh06qq6uVQbky1ejXR0iBA"
	bucketName := "1400720231"
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 上传文件流。
	err = bucket.PutObject(name, fd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	link := scheme + bucketName + "." + endpoint + "/" + name
	return link
}
