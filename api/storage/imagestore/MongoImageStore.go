package imagestore

import (
	"errors"

	"github.com/ZacharyDuve/photoz/api/model/photo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//MongoImageStore implements ImageStore using a mongo db
type mongoImageStore struct {
	iDS             ImageDataStore
	imageCollection *mongo.Collection
}

//NewMongoImageStore creates a new ImageStore that uses a mongo db that the client for has been passed in.
func NewMongoImageStore(iDS ImageDataStore, db *mongo.Database) ImageStore {
	mIS := &mongoImageStore{}
	mIS.imageCollection = db.Collection("image", nil)
	mIS.iDS = iDS
	return mIS
}

func (this *mongoImageStore) Add(img photo.Image) error {
	this.imageCollection.InsertOne()
	return errors.New("Unimplemented")
}

func (this *mongoImageStore) Update(photo.Image) error {
	return errors.New("Unimplemented")
}

func (this *mongoImageStore) GetByPhotoIDAndImageType(photo.ID, photo.ImageType) (photo.Image, error) {
	return nil, errors.New("Unimplemented")
}

func (this *mongoImageStore) Delete(photo.Image) error {
	return errors.New("Unimplemented")
}

func imageToBSONImage(img photo.Image) bson.D
