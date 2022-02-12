package main

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Snake struct {
	body         []PointInGameArea
	head         PointInGameArea
	vector2      pixel.Vec
	direction    rune
	auxDirection rune
}

func snakeInit(window *pixelgl.Window) Snake { // inicjalizacja snake'a
	gridArray[gameAreaX/2][gameAreaY/2] = 1
	point := PointInGameArea{x: gameAreaX / 2, y: gameAreaY / 2}
	snake := Snake{head: PointInGameArea{x: 1, y: 1}, direction: 'd', auxDirection: 'a'}
	snake.body = append(snake.body, point)
	point.x++
	snake.body = append(snake.body, point)
	point.x++
	snake.body = append(snake.body, point)
	point.x++
	snake.body = append(snake.body, point)

	return snake
}

func drawSnake(snake *Snake, imd *imdraw.IMDraw, window *pixelgl.Window) { // rysowanie snake'a
	imd.Clear()
	imd.Color = colornames.Darkgreen
	for x := 0; x < gameAreaX; x++ {
		for y := 0; y < gameAreaY; y++ {
			if gridArray[x][y] == 1 {
				body := pixel.V(float64(x*grid+gameBorder+(grid/2)), float64(y*grid+gameBorder+(grid/2)))
				snake.vector2 = body
				imd.Push(body)
				imd.Circle(grid, defaultThickness)
				imd.Draw(window)
			}
		}
	}
}

func snakeMovement(snake *Snake, food *Food, window *pixelgl.Window, gameOverBool *bool, score *uint) { // poruszanie snake'a
	if snake.direction == 'd' {
		for index := 0; index < len(snake.body)-1; index++ {
			snake.body[index] = snake.body[index+1]
		}
		snake.body[len(snake.body)-1].x++
		snake.head = snake.body[len(snake.body)-1]
	} else if snake.direction == 'a' {
		for index := 0; index < len(snake.body)-1; index++ {
			snake.body[index] = snake.body[index+1]
		}
		snake.body[len(snake.body)-1].x--
		snake.head = snake.body[len(snake.body)-1]
	} else if snake.direction == 'w' {
		for index := 0; index < len(snake.body)-1; index++ {
			snake.body[index] = snake.body[index+1]
		}
		snake.body[len(snake.body)-1].y++
		snake.head = snake.body[len(snake.body)-1]
	} else if snake.direction == 's' {
		for index := 0; index < len(snake.body)-1; index++ {
			snake.body[index] = snake.body[index+1]
		}
		snake.body[len(snake.body)-1].y--
		snake.head = snake.body[len(snake.body)-1]
	}
	snakeCollision(snake, gameOverBool, window)
	snakeUpdate(snake, food, score)
}

func snakeUpdate(snake *Snake, food *Food, score *uint) { // aktualizowanie snake'a
	snakeHead_X_float := float64(snake.head.x)
	snakeHead_Y_float := float64(snake.head.y)
	foodPosition_X_float := float64(food.position.x)
	foodPosition_Y_float := float64(food.position.y)
	distanceSnakeFood := math.Sqrt(math.Pow(snakeHead_X_float-foodPosition_X_float, 2) + math.Pow(snakeHead_Y_float-foodPosition_Y_float, 2))
	for x := 0; x < gameAreaX; x++ {
		for y := 0; y < gameAreaY; y++ {
			if distanceSnakeFood < 2 {
				if snake.direction == 'd' && food.count == 1 {
					point := PointInGameArea{x: snake.head.x + 1, y: snake.head.y}
					snake.body = append(snake.body, point)
					*score += 1
					food.count = 0
					break
				} else if snake.direction == 'a' && food.count == 1 {
					point := PointInGameArea{x: snake.head.x - 1, y: snake.head.y}
					snake.body = append(snake.body, point)
					*score += 1
					food.count = 0
					break
				} else if snake.direction == 'w' && food.count == 1 {
					point := PointInGameArea{x: snake.head.x, y: snake.head.y + 1}
					snake.body = append(snake.body, point)
					*score += 1
					food.count = 0
					break
				} else if snake.direction == 's' && food.count == 1 {
					point := PointInGameArea{x: snake.head.x, y: snake.head.y - 1}
					snake.body = append(snake.body, point)
					*score += 1
					food.count = 0
					break
				}
			}
		}
	}
}

func snakeCollision(snake *Snake, gameOverBool *bool, window *pixelgl.Window) { // sprawdzanie kolizji snake'a

	lastIndex := len(snake.body) - 1
	if snake.head == snake.body[lastIndex] {
		for index := lastIndex - 1; index >= 0; index-- {
			if snake.head.x == snake.body[index].x && snake.head.y == snake.body[index].y {
				*gameOverBool = true
			}
		}
	}
	if snake.head.x == 0 || snake.head.x == gameAreaX-1 || snake.head.y == gameAreaY-1 || snake.head.y == 0 {
		*gameOverBool = true
	}
	if snake.vector2.Y < 85 || snake.vector2.Y > 715 || snake.vector2.X >= 1125 || snake.vector2.X < 85 {
		*gameOverBool = true
	}
}
