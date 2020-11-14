package imagestore

import "github.com/ZacharyDuve/photoz/api/model"

type ImageStore interface {
	AddImage(model.PhotoId, model.Image) error
	GetImage(model.PhotoId, model.ImageType) (model.Image, error)
	DeleteImage(model.PhotoId, model.ImageType) error
}
