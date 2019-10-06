package turtle

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

func BuildTurtle() Turtle {
	var tmp Turtle
	tmp.Width = 41
	tmp.Height = 10
	tmp.Y = 100
	tmp.IncY = 1
	tmp.IsSinking = false
	return tmp
}
