package monkey

import (
	"../properties"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	_ "image/png"
	"log"
)

var image *ebiten.Image

//Monkey The Monkey is the lmain character from the game.
type Monkey struct {
	//Position of the upper left corner of the monkey
	X int
	//Position of the left side of the monkey
	Y int
	//Width of the monkey
	Width int
	//Height of the monkey
	Height int
	//incrementation on the y axis
	IncJump int
	//if true the monkey is jumping
	IsJumping bool
	//direction true = right false=left
	Direction bool
	//X position for the Jumping start
	JumpingStart int
	//Image tro represent the monkey
	Image *ebiten.Image
}

//Jump contains the precalculate Y of the quadratic function
var Jump [50]int

//Init monkey initialisation and building
func Init() *Monkey {
	length := len(Jump)
	for i := 0; i < length; i++ {
		x := (float64(i+1) / float64(len(Jump))) * float64(43)
		Jump[i] = int(-0.1082*x*x + 4.654*x)
	}
	monkey := new(Monkey)
	monkey.Reset()
	image, _, err := ebitenutil.NewImageFromFile("assets/monkey.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	monkey.Image = image
	return monkey
}

//Reset reset the parameter to their initial state
func (m *Monkey) Reset() {
	m.X = properties.Borderwidth - 32
	m.Y = 57
	m.Width = 32
	m.Height = 43
	m.IncJump = -1
	m.IsJumping = false
	m.Direction = false
}

//InitJump initiate the monkey jumping
func (m *Monkey) InitJump(direction bool) {
	m.IsJumping = true
	m.JumpingStart = m.X
	m.Direction = direction
}

//Jump the monkey jumps
func (m *Monkey) Jump() {
	if m.IncJump == len(Jump)-1 {
		if m.Direction {
			m.X = m.JumpingStart + 43
		} else {
			m.X = m.JumpingStart - 43
		}
		m.IsJumping = false
		m.IncJump = -1
	} else {
		m.IncJump = m.IncJump + 1
		m.Y = properties.Groundheight - m.Height - Jump[m.IncJump]
		floatX := float64(m.IncJump) / float64(len(Jump)) * 43
		if m.Direction {
			m.X = int(floatX) + m.JumpingStart
		} else {
			m.X = m.JumpingStart - int(floatX)
		}
	}
}
