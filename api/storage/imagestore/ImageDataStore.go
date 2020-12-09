package imagestore

import (
	"io"

	"github.com/ZacharyDuve/photoz/api/model/photo"
)

//ImageDataStore stores the image data tied to an image
type ImageDataStore interface {
	//Save takes the data on the image and saves it
	Save(photo.Image, io.ReadCloser) error
	//Update updates the data that is store with what is currently on the image.Image
	Update(photo.Image, io.ReadCloser) error
	//Get returns the data for this image, doesn't modify the image
	Get(photo.Image) (io.ReadCloser, error)
	//Delete removes the image data from the datastore
	Delete(photo.Image) error
}
