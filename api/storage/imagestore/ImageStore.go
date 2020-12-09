package imagestore

import (
	"github.com/ZacharyDuve/photoz/api/model/photo"
)

//ImageStore stores the image object and not the specific data for the picture, see ImageDataStore for that
type ImageStore interface {
	Add(photo.Image) error
	Update(photo.Image) error
	GetByPhotoIDAndImageType(photo.ID, photo.ImageType) (photo.Image, error)
	Delete(photo.Image)
}
