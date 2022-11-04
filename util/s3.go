package util

import (
	"fmt"
	"github.com/arvians-id/go-microservice/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"mime/multipart"
	"strings"
)

type StorageS3 struct {
	AccessKeyID     string
	SecretAccessKey string
	MyRegion        string
	MyBucket        string
}

func NewStorageS3(cofiguration *config.Config) *StorageS3 {
	return &StorageS3{
		AccessKeyID:     cofiguration.AwsAccessKeyId,
		SecretAccessKey: cofiguration.AwsSecretKey,
		MyRegion:        cofiguration.AwsRegion,
		MyBucket:        cofiguration.AwsBucket,
	}
}

func (storageS3 *StorageS3) DefaultPath() string {
	filePath := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/assets/%s", storageS3.MyBucket, storageS3.MyRegion, "no-image.png")
	return filePath
}

func (storageS3 *StorageS3) GenerateNewFile(fileName string) (string, string) {
	headerFileName := strings.Split(fileName, ".")
	randomName := RandomString(10) + "." + headerFileName[len(headerFileName)-1]
	filePath := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/assets/%s", storageS3.MyBucket, storageS3.MyRegion, randomName)
	return filePath, randomName
}

func (storageS3 *StorageS3) ConnectToAWS() (*session.Session, error) {
	sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String(storageS3.MyRegion),
			Credentials: credentials.NewStaticCredentials(storageS3.AccessKeyID, storageS3.SecretAccessKey, ""),
		},
	)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func (storageS3 *StorageS3) UploadToAWS(file multipart.File, fileName string, contentType string) error {
	sess, err := storageS3.ConnectToAWS()
	if err != nil {
		log.Println("[AWS][UploadToAWS][ConnectToAWS] problem in connecting to aws, err: ", err.Error())
		return err
	}

	_, err = s3.New(sess).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(storageS3.MyBucket),
		ACL:                  aws.String("public-read"),
		Key:                  aws.String(fileName),
		Body:                 file,
		ContentType:          aws.String(contentType),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	if err != nil {
		log.Println("[AWS][UploadToAWS][PutObject] problem in upload to aws, err: ", err.Error())
		return err
	}

	return nil
}

// Deprecated on production
func (storageS3 *StorageS3) DeleteFromAWS(filePath string) error {
	headerFilePathName := strings.Split(filePath, "/")
	fileName := headerFilePathName[len(headerFilePathName)-1]
	if fileName == "no-image.png" {
		return nil
	}

	sess, err := storageS3.ConnectToAWS()
	if err != nil {
		log.Println("[AWS][DeleteFromAWS][ConnectToAWS] problem in connecting to aws, err: ", err.Error())
		return err
	}

	svc := s3.New(sess)
	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(storageS3.MyBucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		log.Println("[AWS][DeleteFromAWS][DeleteObject] problem in delete from aws, err: ", err.Error())
		return err
	}

	return nil
}
