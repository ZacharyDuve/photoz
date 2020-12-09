package photo

//ImageType is well the type of image, see the const types for better explination
type ImageType string

const (
	//Master is the image that was sent to us, byte for byte. Shouldn't have been modified in anyway. This is for archiving
	Master ImageType = "master"
	//FullRes is the image that is the same exact resolution as original but has been modified (say rotated)
	FullRes ImageType = "fullres"
	//Thumbnail is the master that has been resized to fit in a 300x300 pixel square. - Note the size I might change, pulled that out of thin air
	Thumbnail ImageType = "thumbnail"
)

//Hash is the md5 has of the file storing the image
type Hash string

//Format is the format that the image is encoded in when stored to file, what formats that are avialble are dependant on what processors are available.
type Format string

//Image represents the data about a picture that is stored
type Image interface {
	PhotoID() ID
	Size() uint64
	Width() uint64
	Height() uint64
	Type() ImageType
	Hash() Hash
	Format() Format
}
