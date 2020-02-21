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

func (ctx *Ball) changeDirection() {
	ballX := ctx.x + int(math.Ceil(ctx.vector.x))
	ballY := ctx.y + int(math.Ceil(ctx.vector.y))

	// bounce from text box (where we draw FPS and score)
	// bounce from right border of the text box
	if ballX-BallSize <= TextRight && ballY < TextBottom {
		ctx.vector.x = -ctx.vector.x
	}
	// bounce from bottom of the text box
	if ballX <= TextRight && ballY-BallSize < TextBottom {
		ctx.vector.y = -ctx.vector.y
	}

	// right and left of the playground
	if ballX > ctx.windowWidth-BallSize {
		ctx.vector.x = -ctx.vector.x
	} else if ballX < BallSize {
		ctx.vector.x = -ctx.vector.x
	}

	// bottom and top of the playground
	if ballY > ctx.windowHeight-BallSize {
		ctx.vector.y = -ctx.vector.y
	} else if ballY < BallSize {
		ctx.vector.y = -ctx.vector.y
	}

	platform := ctx.platform
	points := [...]Point{
		// bounce from platform top
		{x: ballX, y: platform.y},
		// bounce from platform bottom
		{x: ballX, y: platform.y + platform.height},
		// bounce from platform left
		{x: platform.x, y: ballY},
		// bounce from platform right
		{x: platform.x + platform.width, y: ballY},

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
		if ctx.Contains(point) && ctx.platform.Contains(point) {
			ctx.BounceFromPoint(point)
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
