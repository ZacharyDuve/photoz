package imagestore

//MongoImageStore implements ImageStore using a mongo db
type MongoImageStore struct {
	iDS ImageDataStore
}
