package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"./monkey"
	"./turtle"

	"github.com/hajimehoshi/ebiten"
)

var background *ebiten.Image
var turtleImage *ebiten.Image
var drawOptions *ebiten.DrawImageOptions
var maxSinkingTurtles int = 1
var turtleDrawImageOptions ebiten.DrawImageOptions
var monkeyDrawImageOptions ebiten.DrawImageOptions
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
		turtles[i] = turtle.BuildTurtle()
		turtles[i].X = i * 45
	}

	turtleImage, _ = ebiten.NewImage(turtles[0].Width, turtles[1].Height, ebiten.FilterDefault)
	turtleImage.Fill(color.RGBA{0, 200, 0, 255})

	background, _ = ebiten.NewImage(320, 200, ebiten.FilterDefault)
	background.Fill(color.White)

	var leftBorderDrawOptions = new(ebiten.DrawImageOptions)
	leftBorderDrawOptions.GeoM.Translate(0, 100)
	leftBorder, _ := ebiten.NewImage(70, 100, ebiten.FilterDefault)
	leftBorder.Fill(color.RGBA{0, 255, 0, 255})
	background.DrawImage(leftBorder, leftBorderDrawOptions)

	rightBorderDrawOptions := new(ebiten.DrawImageOptions)
	rightBorderDrawOptions.GeoM.Translate(250, 100)
	rightBorder, _ := ebiten.NewImage(70, 100, ebiten.FilterDefault)
	rightBorder.Fill(color.RGBA{0, 255, 0, 255})
	background.DrawImage(rightBorder, rightBorderDrawOptions)

}

var selectedTurtle int

func update(screen *ebiten.Image) error {

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	//The background
	screen.DrawImage(background, drawOptions)

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
		turtleDrawImageOptions.GeoM.Translate(float64(72+turtles[i].X), float64(turtles[i].Y))
		screen.DrawImage(turtleImage, &turtleDrawImageOptions)
	}
	//The monkey
	monkeyDrawImageOptions.GeoM.Reset()
	monkeyDrawImageOptions.GeoM.Translate(float64(theMonkey.X), float64(theMonkey.Y))
	screen.DrawImage(theMonkey.Image, &monkeyDrawImageOptions)
	if theMonkey.IsJumping {
		theMonkey.Jump()
	}

	//input poller

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if !theMonkey.IsJumping {
			theMonkey.InitJump(true)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if !theMonkey.IsJumping {
			theMonkey.InitJump(false)
		}
	}
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Tao"); err != nil {
		log.Fatal(err)
	}
}