package uploadedphoto

import (
	"testing"

	"github.com/ZacharyDuve/photoz/api/model/photo"
	testingImage "github.com/ZacharyDuve/photoz/testing/api/model/photo"
)

func newUploadedPhotoWithBlankID() (*UploadedPhoto, error) {
	var blankid photo.ID
	return NewUploadedPhoto(blankid, nil, Status("SomeStatus"))
}

func newUploadedPhotoWithoutMaster() (*UploadedPhoto, error) {
	return NewUploadedPhoto(photo.ID("1"), nil, Status("SomeStatus"))
}

func newUploadedPhotoWithBlankStatus() (*UploadedPhoto, error) {
	var blankStatus Status
	return NewUploadedPhoto(photo.ID("1"), &testingImage.MockImage{}, blankStatus)
}

func newUploadedPhoto() (*UploadedPhoto, error) {
	return NewUploadedPhoto(photo.ID("1"), &testingImage.MockImage{}, Status("SomeStatus"))
}

func TestNewUploadedPhotoWithoutIDReturnsMissingIDError(t *testing.T) {
	_, err := newUploadedPhotoWithBlankID()
	if err == nil || !IsMissingIDError(err) {
		t.Fatal("Failed to return MissingIDError when creating UploadedPhoto without an ID")
	}
}

func TestNewUploadedPhotoWithoutIDReturnsNilUploadedPhoto(t *testing.T) {
	pht, _ := newUploadedPhotoWithBlankID()

	if pht != nil {
		t.Fatal("When creating a new UploadedPhoto with a blank id then we shouldn't have gotten an uploaded photo")
	}
}

func TestNewUploadedPhotoWithoutMasterReturnsMissingMasterError(t *testing.T) {
	_, err := newUploadedPhotoWithoutMaster()

	if err == nil || !IsMissingMasterError(err) {
		t.Fatal("Failed to return MissingMasterError when creating UploadedPhoto without a master image")
	}
}

func TestNewUploadedPhotoWithoutMasterReturnsNilUploadedPhoto(t *testing.T) {
	pht, _ := newUploadedPhotoWithoutMaster()

	if pht != nil {
		t.Fatal("NewUploadedPhoto without a master image should fail as it is a required field")
	}
}

func TestNewUploadedPhotoWithABlankStatusShouldReturnMissingStatusError(t *testing.T) {
	_, err := newUploadedPhotoWithBlankStatus()

	if err == nil || !IsMissingStatusError(err) {
		t.Fatal("NewUploadedPhoto wihtout a status should return a MissingStatusError as it is a required field")
	}
}

func TestNewUploadedPhotoWithValidInputReturnsNoError(t *testing.T) {
	_, err := newUploadedPhoto()

	if err != nil {
		t.Fatal("NewUploadedPhoto with valid parameters should return no error")
	}
}
