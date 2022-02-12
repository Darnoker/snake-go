package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Food struct {
	position PointInGameArea
	count    uint
}

func foodGenerate(food *Food) { // generowanie jedzenia
	food.count = 1
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	food.position.x = int(random.Int31())%(gameAreaX-10) + 5
	food.position.y = int(random.Int31())%(gameAreaY-10) + 5
	gridArray[food.position.x][food.position.y] = 2
}

func drawFood(food *Food, imd *imdraw.IMDraw, window *pixelgl.Window) { //rysowanie jedzenia
	imd.Clear()
	imd.Color = colornames.Darkred
	for x := 0; x < gameAreaX; x++ {
		for y := 0; y < gameAreaY; y++ {
			if gridArray[x][y] == 2 {
				foodPosition := pixel.V(float64(x*grid+gameBorder+(grid/2)), float64(y*grid+gameBorder+(grid/2)))
				imd.Push(foodPosition)
				imd.Circle(grid, defaultThickness)
				imd.Draw(window)
			}
		}
	}
}
