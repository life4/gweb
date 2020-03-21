package main

import (
	"math"

	"github.com/life4/gweb/canvas"
)

type Ball struct {
	Circle
	vector Vector

	windowWidth  int
	windowHeight int

	context  canvas.Context2D
	platform *Platform
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
	if ball.x-ball.radius <= TextRight && ball.y < TextBottom {
		ball.vector.x = -ball.vector.x
	}
	// bounce from bottom of the text box
	if ball.x <= TextRight && ball.y-ball.radius < TextBottom {
		ball.vector.y = -ball.vector.y
	}

	// right and left of the playground
	if ball.x > ball.windowWidth-ball.radius {
		ball.vector.x = -ball.vector.x
	} else if ball.x < ball.radius {
		ball.vector.x = -ball.vector.x
	}

	// bottom and top of the playground
	if ball.y > ball.windowHeight-ball.radius {
		ball.vector.y = -ball.vector.y
	} else if ball.y < ball.radius {
		ball.vector.y = -ball.vector.y
	}

	platform := ball.platform
	points := [...]Point{
		// bounce from platform top
		{x: ball.x, y: platform.rect.y},
		// bounce from platform bottom
		{x: ball.x, y: platform.rect.y + platform.rect.height},
		// bounce from platform left
		{x: platform.rect.x, y: ball.y},
		// bounce from platform right
		{x: platform.rect.x + platform.rect.width, y: ball.y},

		// left-top corner of the platform
		{x: platform.rect.x, y: platform.rect.y},
		// right-top corner of the platform
		{x: platform.rect.x + platform.rect.width, y: platform.rect.y},
		// left-bottom corner of the platform
		{x: platform.rect.x, y: platform.rect.y + platform.rect.height},
		// right-bottom corner of the platform
		{x: platform.rect.x + platform.rect.width, y: platform.rect.y + platform.rect.height},
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
	ball.context.Arc(ball.x, ball.y, ball.radius+1, 0, math.Pi*2)
	ball.context.Fill()
	ball.context.ClosePath()

	ball.changeDirection()

	// move the ball
	ball.x += int(math.Round(ball.vector.x))
	ball.y += int(math.Round(ball.vector.y))

	// draw the ball
	ball.context.SetFillStyle(BallColor)
	ball.context.BeginPath()
	ball.context.Arc(ball.x, ball.y, ball.radius, 0, math.Pi*2)
	ball.context.Fill()
	ball.context.ClosePath()
}
