package turtle

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

var image *ebiten.Image

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

func (s *Turtle) Sink() {
	if s.Y == 100 {
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

func Init() Turtle {
	var tmp Turtle
	tmp.Width = 41
	tmp.Height = 10
	tmp.Y = 100
	tmp.IncY = 1
	tmp.IsSinking = false
	if image == nil {
		image, _ = ebiten.NewImage(tmp.Width, tmp.Height, ebiten.FilterDefault)
		image.Fill(color.RGBA{0, 200, 0, 255})
	}
	tmp.Image = image
	return tmp
}
