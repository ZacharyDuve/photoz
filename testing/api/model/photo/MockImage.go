package photo

import (
	"github.com/ZacharyDuve/photoz/api/model/photo"
)

//MockImage is an implementation of the Image interface to allow for easy testing
type MockImage struct {
	// Size() uint64
	// Width() uint64
	// Height() uint64
	// Hash() Hash
	// Format() Format
	SizeFunc    func() uint64
	WidthFunc   func() uint64
	HeightFunc  func() uint64
	HashFunc    func() photo.Hash
	FormatFunc  func() photo.Format
	TypeFunc    func() photo.ImageType
	PhotoIDFunc func() photo.ID
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

func (mI *MockImage) Hash() photo.Hash {
	return mI.HashFunc()
}

func (mI *MockImage) Format() photo.Format {
	return mI.FormatFunc()
}

func (mI *MockImage) Type() photo.ImageType {
	return mI.TypeFunc()
}

func (mI *MockImage) PhotoID() photo.ID {
	return mI.PhotoIDFunc()
}
