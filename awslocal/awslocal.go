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
	"bytes"
	"net/http"
	"io/ioutil"
	"sync"
)

type AwsS3Local struct {
	AwsS3 *s3.S3
}

func NewAwsS3Local() *AwsS3Local {
	aws_access_key_id := os.Getenv("AWS_KEY")
	aws_secret_access_key := os.Getenv("AWS_SECRET")
	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	_, err := creds.Get()
	if err != nil {
		panic(err)
	}
	helpers.FailOnError(err, fmt.Sprintf("bad credentials: %s", err), true)
	cfg := aws.NewConfig().WithRegion(os.Getenv("AWS_REGION")).WithCredentials(creds)
	return &AwsS3Local{AwsS3: s3.New(session.New(), cfg)}
}

func (awsS3Local *AwsS3Local) RemoveFile(file string) bool {
	params := &s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(file),
	}
	resp, err := awsS3Local.AwsS3.DeleteObject(params)
	helpers.FailOnError(err, "failed remove file from aws s3", true)
	helpers.LogInfo(awsutil.StringValue(resp), true)
	return true
}

// https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGET.html
func (awsS3Local *AwsS3Local) ListObjects(prefix string) *s3.ListObjectsOutput {
	params := &s3.ListObjectsInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Prefix: aws.String(prefix),
	}
	resp, err := awsS3Local.AwsS3.ListObjects(params)
	helpers.FailOnError(err, "failed get list objects in aws s3", true)
	return resp
}

func (awsS3Local *AwsS3Local) MoveFile(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
		return false, err
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size) // read file content to buffer

	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	params := &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key: aws.String(file.Name()),
		Body: fileBytes,
		ContentLength: aws.Int64(size),
		ContentType: aws.String(fileType),
	}
	_, err = awsS3Local.AwsS3.PutObject(params)
	if err != nil {
		panic(err)
		return false, err
	}

	return true, nil
}

func (awsS3Local *AwsS3Local) MoveDir(dirPath string) (bool, error) {
	helpers.LogInfoF("MoveDir", "start")
	defer helpers.LogInfoF("MoveDir", "end")
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
		return false, err
	}

	filesCountChannel := make(chan bool)
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		pathToFile := dirPath + file.Name()
		go func(pathToFile string) {
			defer wg.Done()
			awsS3Local.MoveFile(pathToFile)
			filesCountChannel <- true
		}(pathToFile)
	}
	go func() {
		wg.Wait()
		close(filesCountChannel)
	}()

	go helpers.RemoveImageDirWithLatency(dirPath)

	return true, nil
}
