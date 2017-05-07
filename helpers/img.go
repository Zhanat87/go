package helpers

import (
	"encoding/base64"
	"errors"
	"strings"

	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"bytes"
)

var (
	//ErrBucket       = errors.New("Invalid bucket!")
	//ErrSize         = errors.New("Invalid size!")
	ErrInvalidImage = errors.New("Invalid image!")
)

func SaveImageToDisk(path, fileNameBase, data string) (string, error) {
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return "", ErrInvalidImage
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[idx+8:]))
	buff := bytes.Buffer{}
	_, err := buff.ReadFrom(reader)
	if err != nil {
		return "", err
	}
	//imgCfg, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	_, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	if err != nil {
		return "", err
	}

	//if imgCfg.Width != 750 || imgCfg.Height != 685 {
	//	return "", ErrSize
	//}

	fileName := fileNameBase + "." + fm
	ioutil.WriteFile(path + fileName, buff.Bytes(), 0644)

	return fileName, err
}
