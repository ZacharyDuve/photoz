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

type image struct {
	data          io.Reader
	size          uint64
	height, width uint32
	imageType     ImageType
}

func NewImage(data io.Reader, size uint64, height, width uint32, iType ImageType) Image {
	return &image{data: data, size: size, height: height, width: width, imageType: iType}
}

func (this *image) Read(p []byte) (int, error) {
	return this.data.Read(p)
}

func (this *image) GetSize() uint64 {
	return this.size
}

func (this *image) GetDimensions() (uint32, uint32) {
	return this.height, this.width
}

func (this *image) GetType() ImageType {
	return this.imageType
}
