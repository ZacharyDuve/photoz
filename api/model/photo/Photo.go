package photo

import (
	"time"
)

//ID is the UUID of the photo
type ID string

//Photo represents human interested information about the original picture. Descibes the image.
type Photo interface {
	ID() ID
	TimeTaken() time.Time
	SetTimeTaken(time.Time)
	Description() string
	SetDescription(string)
	Original() Image
	Master() Image
	Thumbnail() Image
}
