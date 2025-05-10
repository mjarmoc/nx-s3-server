package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mjarmoc/x-s3-server/config"
	"github.com/mjarmoc/x-s3-server/modules/aws"
	"github.com/mjarmoc/x-s3-server/server"
)

func main() {
	// ToDo Move to DotEnv

	environment := flag.String("e", "development", "")
	if *environment == "development" {
		os.Setenv("NX_S3_TOKEN", "abcdef")
	}
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	aws.Init()
	server.Init() 
}
