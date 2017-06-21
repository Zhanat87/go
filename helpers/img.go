package helpers

import (
	"encoding/base64"
	"errors"
	"strings"
	"io/ioutil"
	"bytes"
	"sync"
	"os"
	"strconv"
	"path"

	"image"
	"image/png"
	"image/jpeg"
	"image/gif"

	"github.com/nfnt/resize"
	"time"
)

var (
	//ErrBucket       = errors.New("Invalid bucket!")
	//ErrSize         = errors.New("Invalid size!")
	ErrInvalidImage = errors.New("Invalid image!")
	ThumbnailsSizes = [3]uint{100, 300, 500}
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

func MakeThumbnails(basename, path string) int64 {
	//LogInfoF("MakeThumbnails", "start")
	//defer LogInfoF("MakeThumbnails", "end")
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines
	for _, width := range ThumbnailsSizes {
		wg.Add(1)
		// worker
		go func(width uint, basename, path string) {
			defer wg.Done()
			thumb := saveThumbnail(width, basename, path)
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(width, basename, path)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}

func saveThumbnail(width uint, basename, imagePath string) string {
	file, err := os.Open(imagePath + basename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	ext := path.Ext(imagePath + basename)

	var img image.Image
	switch ext {
	case ".jpeg":
		// decode jpeg into image.Image
		img, err = jpeg.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	case ".gif":
		img, err = gif.Decode(file)

	}
	if err != nil {
		panic(err)
	}

	m := resize.Resize(width, 0, img, resize.Lanczos3)

	fullName := imagePath + strconv.Itoa(int(width)) + "_" + basename
	out, err := os.Create(fullName)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// write new image to file
	switch ext {
	case ".jpeg":
		err = jpeg.Encode(out, m, nil)
	case ".png":
		err = png.Encode(out, m)
	case ".gif":
		err = gif.Encode(out, m, nil)
	}
	if err != nil {
		panic(err)
	}

	return fullName
}

func RemoveImageWithThumbnails(basePath, imageName string) {
	err := os.Remove(basePath + imageName)
	if err != nil {
		panic(err)
	}

	for _, width := range ThumbnailsSizes {
		err := os.Remove(basePath + strconv.Itoa(int(width)) + "_" + imageName)
		if err != nil {
			panic(err)
		}
	}
}

func RemoveImageDirWithLatency(basePath string, latency int) {
	time.Sleep(time.Duration(latency) * time.Second)
	err := os.RemoveAll(basePath)
	if err != nil {
		LogInfo("RemoveImageDirWithLatency error: " + err.Error(), false)
	}
}
