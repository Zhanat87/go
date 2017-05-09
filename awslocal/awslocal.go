package awslocal

import (
	"github.com/Zhanat87/go/helpers"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func GetS3() *s3.S3 {
	aws_access_key_id := os.Getenv("AWS_KEY")
	aws_secret_access_key := os.Getenv("AWS_SECRET")
	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	_, err := creds.Get()
	helpers.FailOnError(err, fmt.Sprintf("bad credentials: %s", err), true)
	cfg := aws.NewConfig().WithRegion(os.Getenv("AWS_REGION")).WithCredentials(creds)
	return s3.New(session.New(), cfg)
}

func RemoveFile(file string) bool {
	params := &s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(file),
	}
	resp, err := GetS3().DeleteObject(params)
	helpers.FailOnError(err, "failed remove file from aws s3", true)
	helpers.LogInfo(awsutil.StringValue(resp), true)
	return true
}

func ListObjects(prefix string) *s3.ListObjectsOutput {
	params := &s3.ListObjectsInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Prefix: aws.String(prefix),
	}
	resp, err := GetS3().ListObjects(params)
	helpers.FailOnError(err, "failed get list objects in aws s3", true)
	return resp
}
