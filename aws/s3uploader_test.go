package main

import (
	"log"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
)

type fileFetcher struct{}

func (f *fileFetcher) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	log.Println("Mock Uploaded to S3:", *input.Key)
	return &s3.PutObjectOutput{}, nil
}

func TestUploader(t *testing.T) {

	f := fileFetcher{}
	err := Uploader(&f)

	if err != nil {
		t.Errorf("TestUploader returned an error: %s", err)
	}
}
