package main

import "github.com/life4/gweb/canvas"

type Brick struct {
	Rectangle
	context canvas.Context2D
	cost    int
	removed bool
}

func (brick *Brick) Collide(ball *Ball, bounce bool) bool {
	if brick.removed {
		return false
	}

	// quick checks of ball position
	if ball.x-BallSize > brick.x+brick.width { // ball righter
		return false
	}
	if ball.x+BallSize < brick.x { // ball lefter
		return false
	}
	if ball.y+BallSize < brick.y { // ball upper
		return false
	}
	if ball.y-BallSize > brick.y+brick.height { // ball downer
		return false
	}

	points := [...]Point{
		// bottom of brick collision
		{x: ball.x, y: ball.y - BallSize},
		// top of brick collision
		{x: ball.x + brick.width, y: ball.y + BallSize},
		// left of brick collision
		{x: ball.x, y: ball.y + BallSize},
		// right of brick collision
		{x: ball.x + brick.width, y: ball.y - BallSize},
	}

	for _, point := range points {
		if brick.Contains(point) {
			if bounce {
				ball.BounceFromPoint(point)
			}
			return true
		}
	}

	points = [...]Point{
		// left-top corner of the brick
		{x: brick.x, y: brick.y},
		// right-top corner of the brick
		{x: brick.x + brick.width, y: brick.y},
		// left-bottom corner of the brick
		{x: brick.x, y: brick.y + brick.height},
		// right-bottom corner of the brick
		{x: brick.x + brick.width, y: brick.y + brick.height},
	}

	for _, point := range points {
		if ball.Contains(point) {
			if bounce {
				ball.BounceFromPoint(point)
			}
			return true
		}
	}

	return false
}

func (brick *Brick) Draw(color string) {
	brick.context.SetFillStyle(color)
	brick.context.Rectangle(brick.x, brick.y, brick.width, brick.height).Filled().Draw()
	brick.removed = false
}

func (brick *Brick) Remove() {
	brick.context.SetFillStyle(BGColor)
	brick.context.Rectangle(brick.x, brick.y, brick.width, brick.height).Filled().Draw()
	brick.removed = true
}
