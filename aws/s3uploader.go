package main

import (
	"bytes"
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	AWS_ACCESS_KEY    = "<AWS_ACCESS_KEY>"
	AWS_SECRET_ACCESS = "<AWS_SECRET_ACCESS>"
	BUCKET_NAME       = "<BUCKET_NAME>"
)

var uploader *s3manager.Uploader

func main() {
	uploader = NewUploader()

	upload()
}

func NewUploader() *s3manager.Uploader {
	s3Config := &aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY, AWS_SECRET_ACCESS, ""),
	}

	s3Session := session.New(s3Config)

	return s3manager.NewUploader(s3Session)
}

func upload() {
	log.Println("uploading")

	file, err := os.ReadFile("aws/dog.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	upInput := &s3manager.UploadInput{
		Bucket:      aws.String(BUCKET_NAME), // bucket's name
		Key:         aws.String("dog.jpeg"),  // files destination location
		Body:        bytes.NewReader(file),   // content of the file
		ContentType: aws.String("image/jpg"), // content type
	}
	res, err := uploader.UploadWithContext(context.Background(), upInput)
	log.Printf("res %+v\n", res)
	log.Printf("err %+v\n", err)
}
