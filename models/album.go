package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/satori/go.uuid"
	"os"
	"github.com/Zhanat87/go/helpers"
	"github.com/Zhanat87/go/awslocal"
	"database/sql"
	"strings"
	"strconv"
)

const ALBUMS_BASE_IMAGE_PATH = "static/albums/images/"

/*
https://github.com/go-ozzo/ozzo-dbx/issues/23
dbx is not orm and can do relational queries, etc with relations
need select by separate queries,
find parent model and then related models by foreign id(s)
 */
type Album struct {
	Id          int    	   `json:"id" db:"id"`
	Title       string 	   `json:"title" db:"title"`
	ArtistId    uint           `json:"artistId,string" db:"artist_id"`
	ImageBase64 JsonNullString `json:"image_base_64" db:"-"`
	Image       JsonNullString `json:"image,omitempty" db:"image"`
}

type AlbumForClient struct {
	Album
	ArtistName string `json:"artistName" db:"artist_name"`
}

// Validate validates the Album fields.
func (m Album) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required, validation.Length(4, 120)),
		validation.Field(&m.ArtistId, validation.Required),
	)
}

func (m *Album) BeforeInsert() {
	m.SaveImage()
}

func (m *Album) BeforeUpdate() {
	m.SaveImage()
}

func (m *Album) BeforeDelete() {
	if m.Image.Valid {
		m.RemoveImage(m.Image.String)
	}
}

func (m *Album) SaveImage() {
	imageNotEmpty := m.Image.Valid
	imageString := m.Image.String
	if m.ImageBase64.Valid {
		if m.ImageBase64.String == "remove" {
			m.RemoveImage(imageString)
			m.Image = JsonNullString{sql.NullString{String:"", Valid:false}}
		} else {
			imageUUID := uuid.NewV4()
			path := ALBUMS_BASE_IMAGE_PATH + imageUUID.String() + "/"
			os.Mkdir(path, 0755)
			img, err := helpers.SaveImageToDisk(path, imageUUID.String(), m.ImageBase64.String)
			if err != nil {
				panic(err)
			}
			helpers.MakeThumbnails(img, path)
			ok, err := awslocal.NewAwsS3Local().MoveDir(path)
			if err != nil {
				panic(err)
			}
			if ok {
				go helpers.RemoveImageDirWithLatency(path, 5)
			}
			if imageNotEmpty {
				m.RemoveImage(imageString)
			}
			m.Image = JsonNullString{sql.NullString{String:img, Valid:true}}
		}
	}
}

func (m *Album) RemoveImage(image string) {
	basePath := ALBUMS_BASE_IMAGE_PATH + image[:strings.Index(image, ".")]
	awsS3Local := awslocal.NewAwsS3Local()
	awsS3Local.RemoveFile(basePath + "/" + image)
	for _, width := range helpers.ThumbnailsSizes {
		awsS3Local.RemoveFile(basePath + "/" + strconv.Itoa(int(width)) + "_" + image)
	}
}


func (m *Album) AfterSave() {
	helpers.LogInfoF("after", "save")
}
