package turtle

import (
	"../properties"
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

var image *ebiten.Image

//Turtle one turtle
type Turtle struct {
	//Position of the upper left corner of the turtle
	X int
	//Position of the left side of the turtle
	Y int
	//Width of the turtle
	Width int
	//Height of the turtle
	Height int
	//incrementation on the y axis
	IncY int
	//if true the turtle is sinking
	IsSinking bool
	//Image
	Image *ebiten.Image
}

//Sink the turtle sinks
func (s *Turtle) Sink() {
	if s.Y == properties.Groundheight {
		if s.IncY == -1 {
			s.IsSinking = false
			s.IncY = 0
		} else {
			s.IncY = 1
			s.IsSinking = true
		}
	}
	if s.Y == 200-s.Height {
		s.IncY = -1
	}
	s.Y = s.Y + s.IncY
}

//Init intialize and builds a turtle
func Init() Turtle {
	var tmp Turtle
	tmp.Width = 41
	tmp.Height = 10
	tmp.Y = properties.Groundheight
	tmp.IncY = 1
	tmp.IsSinking = false
	if image == nil {
		image, _ = ebiten.NewImage(tmp.Width, tmp.Height, ebiten.FilterDefault)
		image.Fill(color.RGBA{0, 200, 0, 255})
	}
	tmp.Image = image
	return tmp
}
