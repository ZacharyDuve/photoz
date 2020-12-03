package image

import (
	"io"

	"github.com/ZacharyDuve/photoz/api/model/photo/image"
)

//MockImage is an implementation of the Image interface to allow for easy testing
type MockImage struct {
	// Size() uint64
	// Width() uint64
	// Height() uint64
	// Hash() Hash
	// Format() Format
	// Data() io.ReadCloser
	SizeFunc   func() uint64
	WidthFunc  func() uint64
	HeightFunc func() uint64
	HashFunc   func() image.Hash
	FormatFunc func() image.Format
	DataFunc   func() io.ReadCloser
}

func (mI *MockImage) Size() uint64 {
	return mI.SizeFunc()
}

func (mI *MockImage) Width() uint64 {
	return mI.WidthFunc()
}

func (mI *MockImage) Height() uint64 {
	return mI.HeightFunc()
}

func (mI *MockImage) Hash() image.Hash {
	return mI.HashFunc()
}

func (mI *MockImage) Format() image.Format {
	return mI.FormatFunc()
}

func (mI *MockImage) Data() io.ReadCloser {
	return mI.DataFunc()
}
