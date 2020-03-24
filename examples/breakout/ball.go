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
	// ball.context.SetFillStyle("red")
	// ball.context.Rectangle(point.x, point.y, 2, 2).Filled().Draw()

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
	if ball.vector.x > 0 && ball.x > ball.windowWidth-ball.radius {
		ball.vector.x = -ball.vector.x
	}
	if ball.vector.x < 0 && ball.x < ball.radius {
		ball.vector.x = -ball.vector.x
	}

	// bottom and top of the playground
	if ball.vector.y > 0 && ball.y+ball.radius >= ball.windowHeight {
		ball.vector.y = -ball.vector.y
	}
	if ball.vector.y < 0 && ball.y-ball.radius <= 0 {
		ball.vector.y = -ball.vector.y
	}
	// bounce from platform edges
	point := ball.platform.Touch(*ball)
	if point != nil {
		ball.BounceFromPoint(*point)
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
