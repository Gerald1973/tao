package monkey

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

var image *ebiten.Image

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

var Jump [100]int

func Init() *Monkey {
	length := len(Jump)
	for i := 0; i < length; i++ {
		x := (float64(i+1) / float64(100)) * float64(47)
		Jump[i] = int(-0.09054*x*x + 4.255*x)
	}
	monkey := new(Monkey)
	monkey.Reset()
	if image == nil {
		image, _ = ebiten.NewImage(monkey.Width, monkey.Height, ebiten.FilterDefault)
		image.Fill(color.Black)
	}
	monkey.Image = image
	return monkey
}

func (m *Monkey) Reset() {
	m.X = 36
	m.Y = 50
	m.Width = 20
	m.Height = 50
	m.IncJump = -1
	m.IsJumping = false
	m.Direction = false
}

func (m *Monkey) InitJump(direction bool) {
	m.IsJumping = true
	m.JumpingStart = m.X
	m.Direction = direction
}

func (m *Monkey) Jump() {
	if m.IncJump == 99 {
		if m.Direction {
			m.X = m.JumpingStart + 47
		} else {
			m.X = m.JumpingStart - 47
		}
		m.IsJumping = false
		m.IncJump = -1
	} else {
		m.IncJump = m.IncJump + 1
		m.Y = 100 - m.Height - Jump[m.IncJump]
		floatX := float64(m.IncJump) / float64(100) * 47
		if m.Direction {
			m.X = int(floatX) + m.JumpingStart
		} else {
			m.X = m.JumpingStart - int(floatX)
		}
	}
}
