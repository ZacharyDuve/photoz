package upload

import (
	"github.com/ZacharyDuve/photoz/api/storage/imagestore"
	"github.com/ZacharyDuve/photoz/api/storage/photostore"
)

//AsyncPhotoUploadHandler is a PhotoUploadHandler that handle the processing of the photo asynchronously. This is to stage photos to be processed at an unknown later time.
type AsyncPhotoUploadHandler struct {
	photoStore    photostore.PhotoStore
	masterImStore imagestore.ImageStore
}
