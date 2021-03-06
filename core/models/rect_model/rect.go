package rect_model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//Rectangle
//model that will store inside database
type Rectangle struct {
	Id        string
	X         int64
	Y         int64
	Width     int64
	Height    int64
	CreatedAt time.Time
}

//CreateRectangle
//utility for creation of Rectangle
func CreateRectangle(x int64, y int64, w int64, h int64) *Rectangle {

	return &Rectangle{
		Id:     uuid.NewV4().String(),
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
	}
}
