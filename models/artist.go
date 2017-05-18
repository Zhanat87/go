package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"database/sql"
	"os"
	"github.com/satori/go.uuid"
	"github.com/Zhanat87/go/helpers"
	"sync"
	"image/jpeg"
	"github.com/nfnt/resize"
	"strconv"
)

const imagePath = "static/artists/images/"

// Artist represents an artist record.
type Artist struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	ImageBase64 JsonNullString `json:"image_base_64" db:"-"`
	Image       JsonNullString `json:"image,omitempty" db:"image"`
}

// Validate validates the Artist fields.
func (m Artist) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(4, 120)),
	)
}

func (m *Artist) BeforeInsert() {
	m.SaveImage()
}

func (m *Artist) BeforeUpdate() {
	m.SaveImage()
}

func (m *Artist) BeforeDelete() {
	m.RemoveImage()
}

func (m *Artist) SaveImage() {
	imageNotEmpty := m.Image.Valid
	imageString := m.Image.String
	if m.ImageBase64.Valid {
		if m.ImageBase64.String == "remove" {
			err := os.Remove(imagePath + imageString)
			if err != nil {
				panic(err)
			}
			m.Image = JsonNullString{sql.NullString{String:"", Valid:false}}
		} else {
			imageUUID := uuid.NewV4()
			img, err := helpers.SaveImageToDisk(imagePath, imageUUID.String(), m.ImageBase64.String)
			if err != nil {
				panic(err)
			}
			makeThumbnails(img, imagePath)
			if imageNotEmpty {
				err := os.Remove(imagePath + imageString)
				if err != nil {
					panic(err)
				}
			}
			m.Image = JsonNullString{sql.NullString{String:img, Valid:true}}
		}
	} else {
		m.Image = JsonNullString{sql.NullString{String:"", Valid:false}}
	}
}

func (m *Artist) RemoveImage() {
	imageNotEmpty := m.Image.Valid
	imageString := m.Image.String
	if imageNotEmpty {
		err := os.Remove(imagePath + imageString)
		if err != nil {
			panic(err)
		}
	}
}

func makeThumbnails(basename, path string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines
	widths := [3]uint{100, 300, 500}
	for _, width := range widths {
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

func saveThumbnail(width uint, basename, path string) string {
	file, err := os.Open(path + basename)
	if err != nil {
		panic(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}
	file.Close()

	m := resize.Resize(width, 0, img, resize.Lanczos3)

	fullName := path + strconv.Itoa(int(width)) + "_" + basename
	out, err := os.Create(fullName)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// write new image to file
	err = jpeg.Encode(out, m, nil)
	if err != nil {
		panic(err)
	}
	return fullName
}
