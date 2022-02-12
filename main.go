package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func run() {
	windowCFG := pixelgl.WindowConfig{
		Title:  "Snake Game",
		Bounds: pixel.R(0, 0, 1200, 800),
		VSync:  true,
	}
	gameOn := true
	gameOverBool := false
	speedGame := 15
	var score uint = 0
	window, err := pixelgl.NewWindow(windowCFG)
	snake := snakeInit(window)
	food := Food{count: 0}
	imdSnake := imdraw.New(nil)
	imdFood := imdraw.New(nil)
	if err != nil {
		panic(err)
	}

	for !window.Closed() && gameOn {
		if !gameOverBool {
			window.Update()
			window.Clear(colornames.Antiquewhite)
			drawBorder(window)
			drawSnake(&snake, imdSnake, window)

			if food.count == 0 {
				foodGenerate(&food)
			}
			if food.count == 1 {
				drawFood(&food, imdFood, window)
			}
			checkKey(&snake, window)
			updateGrid(&snake, &food)
			snakeMovement(&snake, &food, window, &gameOverBool, &score)
			updateSpeed(&speedGame, &score)
			drawScore(window, &score)
			time.Sleep(time.Duration(speedGame) * time.Millisecond)
		} else if gameOverBool {
			gameOver(&gameOn, window, &score)
		}
	}
}

func main() {
	pixelgl.Run(run)
}
