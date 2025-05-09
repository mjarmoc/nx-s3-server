package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/s3"
)


func Init() {
	Session, err := createSession()
	if err != nil {
		panic(err)
	}
	session = Session
	fmt.Println("Session established")
	service = &S3{client: s3.New(session)}
	fmt.Println("S3 Service created")
}