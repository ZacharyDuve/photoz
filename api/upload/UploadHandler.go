package upload

import (
	"io"

	"github.com/ZacharyDuve/photoz/api/model/photo"
)

//PhotoUploadHandler is capable of having photos uploaded via a byte stream (io.Reader)
type PhotoUploadHandler interface {
	Upload(io.Reader) (err error, id photo.ID)
}
