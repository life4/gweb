package main

import (
	"math"

	"github.com/life4/gweb/canvas"
)

type Ball struct {
	Point
	vector Vector

	windowWidth  int
	windowHeight int

	context  canvas.Context2D
	platform *Platform
}

func (ball Ball) Contains(point Point) bool {
	hypotenuse := math.Pow(float64(BallSize), 2)
	cathetus1 := math.Pow(float64(point.x-ball.x), 2)
	cathetus2 := math.Pow(float64(point.y-ball.y), 2)
	return cathetus1+cathetus2 < hypotenuse
}

func (ball *Ball) BounceFromPoint(point Point) {
	normal := Vector{
		x: float64(point.x - ball.x),
		y: float64(point.y - ball.y),
	}
	normal = normal.Normalized()
	dot := ball.vector.Dot(normal)
	ball.vector = ball.vector.Sub(normal.Mul(2 * dot))
}

func (ball *Ball) changeDirection() {
	// bounce from text box (where we draw FPS and score)
	// bounce from right border of the text box
	if ball.x-BallSize <= TextRight && ball.y < TextBottom {
		ball.vector.x = -ball.vector.x
	}
	// bounce from bottom of the text box
	if ball.x <= TextRight && ball.y-BallSize < TextBottom {
		ball.vector.y = -ball.vector.y
	}

	// right and left of the playground
	if ball.x > ball.windowWidth-BallSize {
		ball.vector.x = -ball.vector.x
	} else if ball.x < BallSize {
		ball.vector.x = -ball.vector.x
	}

	// bottom and top of the playground
	if ball.y > ball.windowHeight-BallSize {
		ball.vector.y = -ball.vector.y
	} else if ball.y < BallSize {
		ball.vector.y = -ball.vector.y
	}

	platform := ball.platform
	points := [...]Point{
		// bounce from platform top
		{x: ball.x, y: platform.y},
		// bounce from platform bottom
		{x: ball.x, y: platform.y + platform.height},
		// bounce from platform left
		{x: platform.x, y: ball.y},
		// bounce from platform right
		{x: platform.x + platform.width, y: ball.y},

		// left-top corner of the platform
		{x: platform.x, y: platform.y},
		// right-top corner of the platform
		{x: platform.x + platform.width, y: platform.y},
		// left-bottom corner of the platform
		{x: platform.x, y: platform.y + platform.height},
		// right-bottom corner of the platform
		{x: platform.x + platform.width, y: platform.y + platform.height},
	}

	for _, point := range points {
		if ball.Contains(point) && ball.platform.Contains(point) {
			ball.BounceFromPoint(point)
			return
		}
	}
}

func (ball *Ball) handle() {
	// clear out previous render
	ball.context.SetFillStyle(BGColor)
	ball.context.BeginPath()
	ball.context.Arc(ball.x, ball.y, BallSize+1, 0, math.Pi*2)
	ball.context.Fill()
	ball.context.ClosePath()

	ball.changeDirection()

	// move the ball
	ball.x += int(math.Round(ball.vector.x))
	ball.y += int(math.Round(ball.vector.y))

	// draw the ball
	ball.context.SetFillStyle(BallColor)
	ball.context.BeginPath()
	ball.context.Arc(ball.x, ball.y, BallSize, 0, math.Pi*2)
	ball.context.Fill()
	ball.context.ClosePath()
}
