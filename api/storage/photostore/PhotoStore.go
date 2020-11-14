package photostore

import "github.com/ZacharyDuve/photoz/api/model"

/*
	PhotoStore is capable of storing Photo data (the meta data of a photo)
*/
type PhotoStore interface {
	AddPhoto(model.Photo) error
	GetPhoto(model.PhotoId) (model.Photo, error)
	RemovePhoto(model.PhotoId) error
}
