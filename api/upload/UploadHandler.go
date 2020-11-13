package upload

import (
	"io"

	"github.com/ZacharyDuve/photoz/api/model"
)

type UploadHandler interface {
	Upload(io.Reader) (err error, id model.PhotoId)
}
