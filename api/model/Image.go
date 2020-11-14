package model

import "io"

type ImageType string

const (
	Original  ImageType = "ORIGINAL"
	FullRes   ImageType = "FULL_RESOLUTION"
	ThumbNail ImageType = "THUMB_NAIL"
)

type Image interface {
	io.Reader
	GetSize() uint64
	GetDimensions() (height, width uint32)
	GetType() ImageType
}

type basicImage struct {
	data          io.Reader
	size          uint64
	height, width uint32
	imageType     ImageType
}

func NewImage(data io.Reader, size uint64, height, width uint32, iType ImageType) Image {
	return &basicImage{data: data, size: size, height: height, width: width, imageType: iType}
}

func (this *basicImage) Read(p []byte) (int, error) {
	return this.data.Read(p)
}

func (this *basicImage) GetSize() uint64 {
	return this.size
}

func (this *basicImage) GetDimensions() (uint32, uint32) {
	return this.height, this.width
}

func (this *basicImage) GetType() ImageType {
	return this.imageType
}
