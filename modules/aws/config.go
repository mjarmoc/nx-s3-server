package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
)

func createSdkConfig(ctx context.Context) (aws.Config) {
	sdkConfig, err := awsConfig.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println(err)
		panic("Couldn't load default configuration. Have you set up your AWS account?")
	}	
	return sdkConfig
} 