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
	katX := float64(point.x - ball.x)
	katY := float64(point.y - ball.y)
	hypo := math.Sqrt(math.Pow(katX, 2) + math.Pow(katY, 2))
	sin := katX / hypo
	cos := katY / hypo
	ball.vector.x = -ball.vector.x
	newX := ball.vector.x*cos - ball.vector.y*sin
	newY := ball.vector.x*sin + ball.vector.y*cos
	ball.vector.x = newX
	ball.vector.y = newY
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

	// bounce from platform top
	if ctx.vector.y > 0 && ctx.platform.Contains(ballX, ballY+BallSize) {
		ctx.vector.y = -ctx.vector.y
	}
	// bounce from platform bottom
	if ctx.vector.y < 0 && ctx.platform.Contains(ballX, ballY-BallSize) {
		ctx.vector.y = -ctx.vector.y
	}
	// bounce from platform left
	if ctx.vector.x > 0 && ctx.platform.Contains(ballX+BallSize, ballY) {
		ctx.vector.x = -ctx.vector.x
	}
	// bounce from platform right
	if ctx.vector.x < 0 && ctx.platform.Contains(ballX-BallSize, ballY) {
		ctx.vector.x = -ctx.vector.x
	}

	points := [...]Point{
		// left-top corner of the platform
		{x: ctx.platform.x, y: ctx.platform.y},
		// right-top corner of the platform
		{x: ctx.platform.x + ctx.platform.width, y: ctx.platform.y},
		// left-bottom corner of the platform
		{x: ctx.platform.x, y: ctx.platform.y + ctx.platform.height},
		// right-bottom corner of the platform
		{x: ctx.platform.x + ctx.platform.width, y: ctx.platform.y + ctx.platform.height},
	}

	for _, point := range points {
		if ctx.Contains(point) {
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
