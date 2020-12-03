package uploadedphoto

import (
	"errors"

	"github.com/ZacharyDuve/photoz/api/model/photo"
	"github.com/ZacharyDuve/photoz/api/model/photo/image"
)

//Status is the current processing status of the photo
type Status string

//UploadedPhoto is a photo that was uploaded but hasn't been completely processed yet* (*or could just being migrated to photo store)
type UploadedPhoto struct {
	id          photo.ID
	masterImage image.Image
	status      Status
}

const (
	missingIDErrorMsg     string = "Cannot create UploadedPhoto as id was missing"
	missingMasterErrorMsg string = "Cannot create UploadedPhoto as master image is missing"
	missingStatusErrorMsg string = "Cannot create UploadedPhoto as status is missing"
)

//NewUploadedPhoto creates a new UploadedPhoto from the photo.ID, image.Image, and Status that was passed in. All fields are required or error will return and UploadedPhoto will be nil
func NewUploadedPhoto(id photo.ID, master image.Image, status Status) (*UploadedPhoto, error) {
	var err error
	var uPht *UploadedPhoto
	if id == photo.ID("") {
		err = newMissingIDError()
	} else if master == nil {
		err = newMissingMasterError()
	} else if status == Status("") {
		err = newMissingStatusError()
	} else {
		uPht = &UploadedPhoto{}
		uPht.id = id
		uPht.masterImage = master
		uPht.status = status
	}

	return uPht, err
}

//ID returns the id of the UploadedPhoto
func (uPht *UploadedPhoto) ID() photo.ID {
	return uPht.id
}

//Master returns the master image that was uploaded
func (uPht *UploadedPhoto) Master() image.Image {
	return uPht.masterImage
}

//Status returns the current status of the Uploaded Photo
func (uPht *UploadedPhoto) Status() Status {
	return uPht.status
}

func newMissingIDError() error {
	return errors.New(missingIDErrorMsg)
}

//IsMissingIDError checks if the error passed in is a MissingIDError
func IsMissingIDError(err error) bool {
	return err.Error() == missingIDErrorMsg
}

func newMissingMasterError() error {
	return errors.New(missingMasterErrorMsg)
}

//IsMissingMasterError checks if the error passed in is a MissingMasterError
func IsMissingMasterError(err error) bool {
	return err.Error() == missingMasterErrorMsg
}

func newMissingStatusError() error {
	return errors.New(missingStatusErrorMsg)
}

//IsMissingStatusError checks if the passed in error is a MissingStatusError
func IsMissingStatusError(err error) bool {
	return err.Error() == missingStatusErrorMsg
}
