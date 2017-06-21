package awslocal

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/Zhanat87/go/helpers"
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"strconv"
	"encoding/xml"
	"github.com/Zhanat87/go/util"
)

// go test github.com/Zhanat87/go/awslocal
func TestAwsS3Local_MoveFile(t *testing.T) {
	helpers.LoadEnvFile()
	file := "files_for_tests/1.txt"
	awsS3Local := NewAwsS3Local()
	ok, err := awsS3Local.MoveFile(file)
	assert.True(t, ok)
	assert.Nil(t, err)

	checkFileContent(t, file, 1)
	awsS3Local.RemoveFile(file)
}

func checkFileContent(t *testing.T, file string, content int) {
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s", os.Getenv("AWS_S3_URL"), os.Getenv("AWS_BUCKET"), file))
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	res, err := strconv.Atoi(string(body))
	assert.Nil(t, err)
	assert.Equal(t, content, res)
}

func TestAwsS3Local_MoveDir(t *testing.T) {
	helpers.LoadEnvFile()
	path := "files_for_tests/files/"
	awsS3Local := NewAwsS3Local()
	ok, err := awsS3Local.MoveDir(path)
	assert.True(t, ok)
	assert.Nil(t, err)

	checkFileContent(t, path + "2.txt", 2)
	awsS3Local.RemoveFile(path + "2.txt")

	checkFileContent(t, path + "3.txt", 3)
	awsS3Local.RemoveFile(path + "3.txt")
}

type AwsS3Error struct {
	XMLName xml.Name  `xml:"Error"`
	Code    string `xml:"Code"`
	Message string `xml:"Message"`
}

func TestAwsS3Local_RemoveFile(t *testing.T) {
	helpers.LoadEnvFile()
	file := "files_for_tests/4.txt"
	awsS3Local := NewAwsS3Local()
	ok, err := awsS3Local.MoveFile(file)
	assert.True(t, ok)
	assert.Nil(t, err)

	checkFileContent(t, file, 4)

	ok, err = awsS3Local.RemoveFile(file)
	assert.True(t, ok)
	assert.Nil(t, err)

	var awsS3Error = &AwsS3Error{}
	xmlDoc := util.FetchURL(fmt.Sprintf("%s/%s/%s", os.Getenv("AWS_S3_URL"), os.Getenv("AWS_BUCKET"), file))
	util.ParseXML(xmlDoc, &awsS3Error)
	assert.Equal(t, "AccessDenied", awsS3Error.Code)
	assert.Equal(t, "Access Denied", awsS3Error.Message)
}