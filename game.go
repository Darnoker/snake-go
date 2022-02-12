package main

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

var gridArray [gameAreaX][gameAreaY]int

type PointInGameArea struct { //punkty znajdujace sie w strefie gry
	x int
	y int
}

func drawBorder(window *pixelgl.Window) { // rysowanie granic
	leftBorder := imdraw.New(nil)
	bottomBorder := imdraw.New(nil)
	rightBorder := imdraw.New(nil)
	upBorder := imdraw.New(nil)

	leftBorder.Color = colornames.Burlywood
	leftBorder.Push(pixel.V(0, windowHeight))
	leftBorder.Push(pixel.V((gameBorder + 25), 0))
	leftBorder.Rectangle(0)

	bottomBorder.Color = colornames.Burlywood
	bottomBorder.Push(pixel.V(0, (gameBorder + 25)))
	bottomBorder.Push(pixel.V(windowWidth, 0))
	bottomBorder.Rectangle(0)

	rightBorder.Color = colornames.Burlywood
	rightBorder.Push(pixel.V(windowWidth, 0))
	rightBorder.Push(pixel.V(windowWidth-(gameBorder+25), windowHeight))
	rightBorder.Rectangle(0)

	upBorder.Color = colornames.Burlywood
	upBorder.Push(pixel.V(0, windowHeight))
	upBorder.Push(pixel.V(windowWidth, windowHeight-(gameBorder+25)))
	upBorder.Rectangle(0)

	leftBorder.Draw(window)
	bottomBorder.Draw(window)
	rightBorder.Draw(window)
	upBorder.Draw(window)
}

func checkKey(snake *Snake, window *pixelgl.Window) { //sprawdzanie jaki klawisz zostal nacisniety
	if window.JustReleased(pixelgl.KeyW) && snake.auxDirection != 'w' {
		snake.direction = 'w'
		snake.auxDirection = 's'
	} else if window.JustReleased(pixelgl.KeyS) && snake.auxDirection != 's' {
		snake.direction = 's'
		snake.auxDirection = 'w'
	} else if window.JustReleased(pixelgl.KeyA) && snake.auxDirection != 'a' {
		snake.direction = 'a'
		snake.auxDirection = 'd'
	} else if window.JustReleased(pixelgl.KeyD) && snake.auxDirection != 'd' {
		snake.direction = 'd'
		snake.auxDirection = 'a'
	} else if window.JustReleased(pixelgl.KeyW) && window.JustReleased(pixelgl.KeyA) {
		snake.direction = 'w'
		snake.auxDirection = 's'
	} else if window.JustReleased(pixelgl.KeyW) && window.JustReleased(pixelgl.KeyD) {
		snake.direction = 'w'
		snake.auxDirection = 's'
	} else if window.JustReleased(pixelgl.KeyS) && window.JustReleased(pixelgl.KeyA) {
		snake.direction = 's'
		snake.auxDirection = 'w'
	} else if window.JustReleased(pixelgl.KeyS) && window.JustReleased(pixelgl.KeyD) {
		snake.direction = 's'
		snake.auxDirection = 'w'
	}
}

func updateGrid(snake *Snake, food *Food) { // update'owanie siatki sluzacej do rysowania snake'a i jedzenia
	for x := 0; x < gameAreaX; x++ {
		for y := 0; y < gameAreaY; y++ {
			gridArray[x][y] = 0
			if x == food.position.x && y == food.position.y {
				gridArray[x][y] = 2
			}
		}
	}
	for index := range snake.body {
		gridArray[snake.body[index].x][snake.body[index].y] = 1
	}
}

func gameOver(gameOn *bool, window *pixelgl.Window, score *uint) { //funkcja wywolujaca koniec gry
	gameOverScreen(window, score)
	*gameOn = false
}

func gameOverScreen(window *pixelgl.Window, score *uint) { // funkcja wywolujaca ekran konca gry
	time.Sleep(1 * time.Second)
	window.Clear(colornames.Black)
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicText := text.New(pixel.V(400, 400), basicAtlas)
	fmt.Fprintln(basicText, "GAME OVER")
	fmt.Fprintln(basicText, "Score: ", *score)
	basicText.Draw(window, pixel.IM.Scaled(basicText.Orig, 6))
	window.Update()
	time.Sleep(2 * time.Second)
}

func drawScore(window *pixelgl.Window, score *uint) { // funkcja rysujaca licznik punktow
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicText := text.New(pixel.V(600, 775), basicAtlas)
	basicText.Color = colornames.Black
	if *score > 20 {
		fmt.Fprintln(basicText, "Score: ", *score)
		fmt.Fprintln(basicText, "HARD")
	} else {
		fmt.Fprintln(basicText, "Score: ", *score)
	}
	basicText.Draw(window, pixel.IM.Scaled(basicText.Orig, 2))
}

func updateSpeed(speed *int, score *uint) { // funkcja aktualizujaca predkosc
	if *score > 5 {
		*speed = 10
	}
	if *score > 10 {
		*speed = 5
	}
	if *score > 20 {
		*speed = 0
	}
}
