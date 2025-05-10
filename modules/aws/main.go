package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/mjarmoc/x-s3-server/config"
)


func Init() {	
	ctx := context.Background()
	sdkConfig := createSession(ctx)
	service = &S3{client: s3.NewFromConfig(sdkConfig, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(config.GetConfig().GetString("aws.endpoint"))
		o.UsePathStyle = true
	})}
	fmt.Println("S3 Service created")
}