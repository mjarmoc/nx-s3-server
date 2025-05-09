package aws

import (
	"os"

	"github.com/mjarmoc/x-s3-server/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	AWSSession "github.com/aws/aws-sdk-go/aws/session"
)

var session *AWSSession.Session

func createSession() (*AWSSession.Session, error) {
	
	config := config.GetConfig()
	region := config.GetString("aws.region")
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID") 
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	endpoint := config.GetString("aws.endpoint")

 return AWSSession.NewSession(&aws.Config{
	Endpoint: aws.String(endpoint),
  	Region: aws.String(region),
	S3ForcePathStyle: aws.Bool(true),
  	Credentials: credentials.NewStaticCredentials(
   		accessKey,
   		secretKey,
   		"",
  	),
 })
}