package imagestore

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/ZacharyDuve/photoz/api/model/photo"
)

//FSImageDataStore implements ImageDataStore with using a FileSystem (FS) to handle storing that data
type FSImageDataStore struct {
	rootPath string
}

//NewFSImageDataStore returns a new FSImageDataStore with the rootPath that was provided
func NewFSImageDataStore(rootPath string) ImageDataStore {
	fsIDS := &FSImageDataStore{rootPath: rootPath}

	fsIDS.rootPath = path.Clean(rootPath)

	if fsIDS.rootPath == "/" {
		fsIDS.rootPath = ""
	}

	return fsIDS
}

//Save stores the photo's data to disk
func (fsIDS *FSImageDataStore) Save(img photo.Image, r io.ReadCloser) error {
	var err error
	path := getPathForRootPathAndImage(fsIDS.rootPath, img)
	defer r.Close()

	err = writeData(path, r, os.Stat, os.IsNotExist, os.Create)

	return err
}

//Update replaces the photo's data on disk
func (fsIDS *FSImageDataStore) Update(img photo.Image, r io.ReadCloser) error {
	var err error
	path := getPathForRootPathAndImage(fsIDS.rootPath, img)
	defer r.Close()

	err = writeData(path, r, os.Stat, os.IsExist, os.Open)

	return err
}

//Trying something to combine update and create into a single core logic
func writeData(path string, r io.ReadCloser, statFunc func(string) (os.FileInfo, error), existsFunc func(error) bool, getFileFunc func(string) (*os.File, error)) error {
	var err error
	if _, statErr := statFunc(path); existsFunc(statErr) {
		var f *os.File
		f, err = getFileFunc(path)
		defer f.Close()

		if err == nil {
			_, err = io.Copy(f, r)
		}
	} else {
		err = statErr
	}

	return err
}

func getFile(path string, statFunc func(string) (os.FileInfo, error), existsFunc func(error) bool, getFileFunc func(string) (*os.File, error)) (f *os.File, err error) {
	if _, statErr := statFunc(path); existsFunc(statErr) {
		f, err = getFileFunc(path)
	} else {
		err = statErr
	}

	return f, err
}

//Get looks up the photo data that is stored
func (fsIDS *FSImageDataStore) Get(img photo.Image) (r io.ReadCloser, err error) {
	path := getPathForRootPathAndImage(fsIDS.rootPath, img)

	r, err = getFile(path, os.Stat, os.IsExist, os.Open)

	return r, err
}

//Delete removes the photo's data file from disk
func (fsIDS *FSImageDataStore) Delete(img photo.Image) error {
	path := getPathForRootPathAndImage(fsIDS.rootPath, img)
	err := os.Remove(path)
	//TODO: Probably want to cleanup any empty dirs that are left via deleting a last file.
	return err
}

//Assumption that the Image given was null checked before
func getPathForRootPathAndImage(rPath string, img photo.Image) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s.%s", rPath, img.Type(), string(img.PhotoID())[0:2], string(img.PhotoID())[2:4], string(img.PhotoID())[4:6], img.PhotoID(), img.Format())
}
