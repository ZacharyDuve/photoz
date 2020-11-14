package model

import (
	"crypto/md5"
	"fmt"
	"image"
)

//ID of the photo
type PhotoId string

func GeneratePhotoIdFromImage(img image.Image) PhotoId {
	//Grab a slice that can fit all of the bytes that we need. Which is width * height * 4 (number of bytes to pixel)
	bts := make([]byte, img.Bounds().Dx()*img.Bounds().Dy()*4)
	curBTSIndex := 0

	for j := img.Bounds().Min.X; j < img.Bounds().Max.X; j++ {

		for k := img.Bounds().Min.Y; k < img.Bounds().Max.Y; k++ {

			curColor := img.At(j, k)
			bts[curBTSIndex], bts[curcurBTSIndex+1], bts[curcurBTSIndex+2], bts[curcurBTSIndex+3] = curColor.RGBA()
		}
	}

	return PhotoId(fmt.Sprintf("%x", md5.Sum(data)))
}
