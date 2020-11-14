package model

import (
	"time"
)

type Photo interface {
	GetId() PhotoId
	GetTimeTaken() time.Time
	SetTimeTaken(time.Time)
	GetDescription() string
	SetDescription(string)
}

type photo struct {
	id          PhotoId
	timeTaken   time.Time
	description string

	// GetOriginalImage() Image
	// GetMasterImage() Image
	// GetThumbnailImage() Image
	// GetImages() []Image
	// //Could error if we are adding invalid image
	// AddImage(Image) error
}

//New Photo that contains the information provided
func NewPhoto(id PhotoId, timeTaken time.Time, desc string) Photo {
	return &photo{id: id, timeTaken: timeTaken, description: desc}
}

//Get the Id for this photo. Thinking that this will be the hash of the image that was uploaded
func (this *photo) GetId() PhotoId {
	return this.id
}

func (this *photo) GetTimeTaken() time.Time {
	return this.timeTaken
}

func (this *photo) SetTimeTaken(tT time.Time) {
	this.timeTaken = tT
}

func (this *photo) GetDescription() string {
	return this.description
}

func (this *photo) SetDescription(desc string) {
	this.description = desc
}
