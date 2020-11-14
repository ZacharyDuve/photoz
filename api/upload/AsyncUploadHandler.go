package upload

import (
	"github.com/ZacharyDuve/photoz/api/storage/imagestore"
	"github.com/ZacharyDuve/photoz/api/storage/photostore"
)

type AsyncUploadHandler struct {
	photoStore    photostore.PhotoStore
	masterImStore imagestore.ImageStore
}
