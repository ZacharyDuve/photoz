package photostore

import (
	"github.com/ZacharyDuve/photoz/api/model/photo"
)

//PhotoStore stores data for photo.Photo
type PhotoStore interface {
	AddPhoto(photo.Photo) error
	GetPhotoByID(photo.ID) (photo.Photo, error)
	RemovePhoto(photo.ID) error
}
