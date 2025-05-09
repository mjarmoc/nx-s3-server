package aws

import (
	"bytes"
	"io"

	"github.com/mjarmoc/x-s3-server/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)


type S3 struct{
	client *s3.S3
}

var service *S3

func GetService() (*S3) {
	return service;
}

func (s S3) List() (*s3.ListObjectsV2Output, error) {
	bucket:= config.GetConfig().GetString("aws.bucket")
	input := s3.ListObjectsV2Input{Bucket: aws.String(bucket)}
	file, err := s.client.ListObjectsV2(&input)
	if err !=nil {
		return nil, err
	}
	return file, nil
}

func (s S3) Get(hash string) ([]byte, error) {
	bucket:= config.GetConfig().GetString("aws.bucket")
	input := s3.GetObjectInput{Key: aws.String(hash), Bucket: aws.String(bucket)}
	output, err := s.client.GetObject(&input)
	if err !=nil {
		if aerr, ok := err.(awserr.Error); ok {
        	switch aerr.Code() {
        		case s3.ErrCodeNoSuchKey:
            		return nil, &CacheNotFoundError{Hash: hash}
        	}
		}
		return nil, err
	}
	file, err2 := io.ReadAll(output.Body)
	defer output.Body.Close()
	if err2 !=nil {
		return nil, err
	}
	return file, nil
}

func (s S3) Upload(hash string, file []byte) (*s3.PutObjectOutput, error) {
	bucket:= config.GetConfig().GetString("aws.bucket")
	input := s3.PutObjectInput{
		Key: aws.String(hash), 
		Bucket: aws.String(bucket),
		Body: bytes.NewReader(file),
	}
	output, err := s.client.PutObject(&input)
	if err !=nil {
		return nil, err
	}
	return output, nil
}

