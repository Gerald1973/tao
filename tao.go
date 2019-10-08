package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"./monkey"
	"./properties"
	"./turtle"

	"github.com/hajimehoshi/ebiten"
)

var background *ebiten.Image
var backgroundDrawOptions *ebiten.DrawImageOptions
var maxSinkingTurtles int = 1
var turtleDrawImageOptions ebiten.DrawImageOptions
var monkeyDrawImageOptions ebiten.DrawImageOptions
var selectedTurtle int
var turtles [4]turtle.Turtle
var theMonkey monkey.Monkey

func selectTurtle() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(4)
}

func countSinkingTurtle(turtles [4]turtle.Turtle) int {
	c := 0
	for i := 0; i < len(turtles); i++ {
		if turtles[i].IsSinking {
			c++
		}
	}
	return c
}

func init() {

	theMonkey = *monkey.Init()

	for i := 0; i < len(turtles); i++ {
		turtles[i] = turtle.Init()
		turtles[i].X = i*(turtles[i].Width+properties.Gutter) + (properties.Borderwidth + properties.Gutter)
	}

	background, _ = ebiten.NewImage(properties.Screenwidth, properties.Screenheight, ebiten.FilterDefault)
	background.Fill(color.White)

	var leftBorderDrawOptions = new(ebiten.DrawImageOptions)
	leftBorderDrawOptions.GeoM.Translate(0, properties.Groundheight)
	leftBorder, _ := ebiten.NewImage(properties.Borderwidth, 100, ebiten.FilterDefault)
	leftBorder.Fill(color.RGBA{0, 255, 0, 255})
	background.DrawImage(leftBorder, leftBorderDrawOptions)

	rightBorderDrawOptions := new(ebiten.DrawImageOptions)
	rightBorderDrawOptions.GeoM.Translate(247, properties.Groundheight)
	rightBorder, _ := ebiten.NewImage(properties.Borderwidth, 100, ebiten.FilterDefault)
	rightBorder.Fill(color.RGBA{0, 255, 0, 255})
	background.DrawImage(rightBorder, rightBorderDrawOptions)

}

func update(screen *ebiten.Image) error {

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	//The background
	screen.DrawImage(background, backgroundDrawOptions)

	//The turtles
	if countSinkingTurtle(turtles) < maxSinkingTurtles {
		selectedTurtle = selectTurtle()
	} else {
		selectedTurtle = -1
	}
	for i := 0; i < len(turtles); i++ {
		turtleDrawImageOptions.GeoM.Reset()
		if turtles[i].IsSinking {
			turtles[i].Sink()
		}
		if i == selectedTurtle && selectedTurtle != -1 {
			if !turtles[i].IsSinking {
				turtles[i].Sink()
			}
		}
		turtleDrawImageOptions.GeoM.Translate(float64(turtles[i].X), float64(turtles[i].Y))
		screen.DrawImage(turtles[i].Image, &turtleDrawImageOptions)
	}
	//The monkey
	monkeyDrawImageOptions.GeoM.Reset()
	monkeyDrawImageOptions.GeoM.Translate(float64(theMonkey.X), float64(theMonkey.Y))
	screen.DrawImage(theMonkey.Image, &monkeyDrawImageOptions)
	if theMonkey.IsJumping {
		theMonkey.Jump()
	} else {
		if !isMonkeyOnGround() {
			//todo: monkey die
			//todo: one live less
			//monkey reinit
			theMonkey.Reset()
		}
	}

	//input poller
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if !theMonkey.IsJumping && isMonkeyOnGround() {
			theMonkey.InitJump(true)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if !theMonkey.IsJumping && isMonkeyOnGround() {
			theMonkey.InitJump(false)
		}
	}
	return nil
}

func isMonkeyOnGround() bool {
	result := false
	if theMonkey.Y == 57 {
		if theMonkey.X < 70 {
			result = true
		} else if theMonkey.X+theMonkey.Width > properties.Screenwidth-properties.Borderwidth {
			result = true
		} else {
			for i := 0; i < len(turtles); i++ {
				xOk := turtles[i].X <= theMonkey.X && turtles[i].X+turtles[i].Width >= theMonkey.X+theMonkey.Width
				yOk := turtles[i].Y < properties.Groundheight+turtles[i].Height/2
				if xOk && yOk {
					result = true
				}
			}
		}
	}
	return result
}

func main() {
	if err := ebiten.Run(update, properties.Screenwidth, 240, 2, "Tao"); err != nil {
		log.Fatal(err)
	}
}
