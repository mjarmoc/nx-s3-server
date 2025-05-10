package aws

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/mjarmoc/x-s3-server/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
)


type S3 struct{
	client *s3.Client
}

var service *S3

func GetService() (*S3) {
	return service;
}

func (s S3) List(ctx context.Context) (*s3.ListObjectsV2Output, error) {
	bucket:= config.GetConfig().GetString("aws.bucket")
	input := s3.ListObjectsV2Input{Bucket: aws.String(bucket)}
	file, err := s.client.ListObjectsV2(ctx, &input)
	if err !=nil {
		return nil, err
	}
	return file, nil
}

func (s S3) Get(ctx context.Context, hash string) ([]byte, error) {
	bucket:= config.GetConfig().GetString("aws.bucket")
	input := s3.GetObjectInput{Key: aws.String(hash), Bucket: aws.String(bucket)}
	output, err := s.client.GetObject(ctx, &input)
	if err !=nil {
		var ae smithy.APIError
		if errors.As(err, &ae) {
        	switch ae.ErrorCode() {
        		case "NoSuchKey":
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

func (s S3) Upload(ctx context.Context, hash string, file []byte) (*s3.PutObjectOutput, error) {
	bucket:= config.GetConfig().GetString("aws.bucket")
	input := s3.PutObjectInput{
		Key: aws.String(hash), 
		Bucket: aws.String(bucket),
		Body: bytes.NewReader(file),
	}
	output, err := s.client.PutObject(ctx, &input)
	if err !=nil {
		return nil, err
	}
	return output, nil
}

