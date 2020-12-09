package imagestore

import (
	"fmt"
	"testing"

	"github.com/ZacharyDuve/photoz/api/model/photo"
	testingphoto "github.com/ZacharyDuve/photoz/testing/api/model/photo"
)

func TestGetPathForRootPathAndImageForTestImageGeneratesCorrectPath(t *testing.T) {
	rootPath := "/mnt/somwhere/photoz"
	testUUID := "123e4567-e89b-12d3-a456-426614174000"
	testImage := &testingphoto.MockImage{}
	testImage.PhotoIDFunc = func() photo.ID {
		return photo.ID(testUUID)
	}
	testImage.FormatFunc = func() photo.Format {
		return photo.Format("jpeg")
	}
	testImage.TypeFunc = func() photo.ImageType {
		return photo.Master
	}

	generatedPath := getPathForRootPathAndImage(rootPath, testImage)
	expectedPath := fmt.Sprintf("%s/%s/%s/%s/%s/%s.%s", rootPath, photo.Master, "12", "3e", "45", testUUID, "jpeg")

	if generatedPath != expectedPath {
		t.Fatalf("got %s but expected %s", generatedPath, expectedPath)
	}
}
