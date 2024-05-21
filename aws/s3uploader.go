package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	AWS_ACCESS_KEY    = "<AWS_ACCESS_KEY>"
	AWS_SECRET_ACCESS = "<AWS_SECRET_ACCESS>"
	BUCKET_NAME       = "<BUCKET_NAME>"
)

type ReportS3 interface {
	PutObject(*s3.PutObjectInput) (*s3.PutObjectOutput, error)
}

func main() {

	// Create AWS Session
	s3Config := &aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY, AWS_SECRET_ACCESS, ""),
	}
	s3Session, _ := session.NewSession(s3Config)
	svc := s3.New(s3Session)

	if Uploader(svc) == nil {
		fmt.Println("File uploaded")
	}
}

// Uploader uploads a file to S3 - This is the function to be tested!
func Uploader(s3Svc ReportS3) error {
	f, err := os.Open("dog.jpeg")
	if err != nil {
		return err
	}
	return uploadToS3(s3Svc, f)
}

// UploadToS3 will upload a single file to S3, it will require a pre-built aws s3 service.
func uploadToS3(s3Svc ReportS3, file *os.File) error {

	// Get file size and read the file content into a buffer
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// S3 Name
	s3Key := path.Base(file.Name())
	_, err = s3Svc.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(BUCKET_NAME),
		Key:           aws.String(s3Key),
		ACL:           aws.String("private"),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(http.DetectContentType(buffer)),
	})
	return err
}
