package pkg

import (
	"bytes"
	"context"
	"fiber-template/config"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"mime/multipart"
	"strings"
)

var awsConfig = config.Config.Aws

func AwsUpload(file multipart.File, fileName string) (string, error) {
	// 获取文件的 Content-Type
	contentType := ""
	// 获取输出文件的扩展名
	ext := strings.ToLower(GetFileExtension(fileName))

	// 根据文件扩展名选择适当的编码器
	switch ext {
	case ".png":
		contentType = "image/png"
	case ".jpeg", ".jpg":
		contentType = "image/jpeg"
	case ".bmp":
		contentType = "image/bmp"
	case ".webp":
		contentType = "image/webp"
	}

	fmt.Println("Detected Content-Type:", contentType)
	region := awsConfig.Region

	//client := s3.New(s3.Options{
	//	Region:      region,
	//	Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(awsConfig.AccessKeyID, awsConfig.SecretAccessKey, "")),
	//})
	client := GetAwsClient()

	bucket := awsConfig.Bucket
	_, err := client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(fileName),
		Body:        file, // 这里也可以使用其他 io.Reader 实例实现对数据流的上传
		ACL:         types.ObjectCannedACLPublicRead,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, fileName)
	return url, nil
}

func GetAwsClient() *s3.Client {
	return s3.New(s3.Options{
		Region:      awsConfig.Region,
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(awsConfig.AccessKeyID, awsConfig.SecretAccessKey, "")),
	})
}

func AwsUploadBody(fileName string, fileBody []byte) (string, error) {
	// 获取文件的 Content-Type
	contentType := ""
	// 获取输出文件的扩展名
	ext := strings.ToLower(GetFileExtension(fileName))

	// 根据文件扩展名选择适当的编码器
	switch ext {
	case ".png":
		contentType = "image/png"
	case ".jpeg", ".jpg":
		contentType = "image/jpeg"
	case ".bmp":
		contentType = "image/bmp"
	case ".webp":
		contentType = "image/webp"
	}
	bucket := awsConfig.Bucket
	region := awsConfig.Region
	client := GetAwsClient()
	_, err := client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(fileBody), // 这里也可以使用其他 io.Reader 实例实现对数据流的上传
		ACL:         types.ObjectCannedACLPublicRead,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, fileName)
	return url, nil
}

func AwsUploadAudio(file multipart.File, fileName string) (string, error) {

	// 获取文件的 Content-Type
	contentType := ""
	// 获取输出文件的扩展名
	ext := strings.ToLower(GetFileExtension(fileName))

	// 根据文件扩展名选择适当的编码器
	switch ext {
	case ".mp3", ".mpga", ".mp2": // 这三种格式共享同一 MIME 类型
		contentType = "audio/mpeg"
	case ".wav":
		contentType = "audio/wav"
	case ".flac":
		contentType = "audio/flac"
	case ".aac":
		contentType = "audio/aac"
	case ".ogg":
		contentType = "audio/ogg"
	case ".m4a":
		contentType = "audio/mp4"
	case ".wma":
		contentType = "audio/x-ms-wma"
	case ".opus":
		contentType = "audio/opus"
	case ".amr":
		contentType = "audio/amr"
	case ".webm":
		contentType = "audio/webm" // 如果是音频部分
		// 或者 contentType = "video/webm" 如果是视频部分
	}

	//fmt.Println("Detected Content-Type:", contentType)

	//client := s3.New(s3.Options{
	//	Region:      region,
	//	Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(config.Config("Aws_AccessKeyID"), config.Config("Aws_SecretAccessKey"), "")),
	//})

	client := GetAwsClient()
	region := awsConfig.Region
	bucket := awsConfig.Bucket
	_, err := client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(fileName),
		Body:        file, // 这里也可以使用其他 io.Reader 实例实现对数据流的上传
		ACL:         types.ObjectCannedACLPublicRead,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, fileName)
	return url, nil
}

// GetFileExtension 获取文件扩展名
func GetFileExtension(filePath string) string {
	ext := strings.ToLower(filePath[strings.LastIndex(filePath, "."):])
	return ext
}
