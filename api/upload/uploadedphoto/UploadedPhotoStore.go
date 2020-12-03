package uploadedphoto

import "github.com/ZacharyDuve/photoz/api/model/photo"

//Store is a store that is to save uploadedphotos while they are being processed
type Store interface {
	AddPhoto(*UploadedPhoto) error
	UpdatePhoto(*UploadedPhoto) error
	GetPhotoByID(photo.ID) (error, *UploadedPhoto)
	RemovePhotoByID(photo.ID) error
}
