package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	bucketOpt   = flag.String("bucket", "", "bucket name")
	fileNameOpt = flag.String("file", "", "file name")
	expireOpt   = flag.Duration("exp", 12*time.Hour, "e.g 600(5 minutes). default 12 hours")
)

func main() {
	flag.Parse()

	if *bucketOpt == "" || *fileNameOpt == "" {
		panic("--bucket and --file is required")
	}

	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := s3.New(sess, &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})

	poi := &s3.PutObjectInput{
		Bucket: bucketOpt,
		Key:    fileNameOpt,
	}
	req, _ := svc.PutObjectRequest(poi)

	preSignURL, err := req.Presign(*expireOpt)
	if err != nil {
		panic(err)
	}

	curl := "curl -X PUT --upload-file " + *fileNameOpt + "" + preSignURL
	fmt.Printf("【URL】\n%s\n【curl】\n%s\n【expire time】\n%s\n", preSignURL, curl, *expireOpt)
}
