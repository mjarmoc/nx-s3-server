package main

import (
	"flag"
	"fmt"
	"os"

	"nx-s3-server/config"
	"nx-s3-server/server"
)

func main() {
	// ToDo Move to DotEnv
	os.Setenv("NX_S3_TOKEN", "abcdef")
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	server.Init() 
}
