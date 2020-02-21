package main

import "github.com/life4/gweb/canvas"

type Brick struct {
	Rectangle
	context canvas.Context2D
	cost    int
	removed bool
}

func (brick Brick) Contains(x, y int) bool {
	if y < brick.y { // ball upper
		return false
	}
	if y > brick.y+brick.height { // ball downer
		return false
	}
	if x > brick.x+brick.width { // ball righter
		return false
	}
	if x < brick.x { // ball lefter
		return false
	}
	return true
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

	// bottom of brick collision
	if ball.vector.y < 0 && brick.Contains(ball.x, ball.y-BallSize) {
		if bounce {
			ball.vector.y = -ball.vector.y
		}
		return true
	}
	// top of brick collision
	if ball.vector.y > 0 && brick.Contains(ball.x, ball.y+BallSize) {
		if bounce {
			ball.vector.y = -ball.vector.y
		}
		return true
	}
	// left of brick collision
	if ball.vector.x > 0 && brick.Contains(ball.x+BallSize, ball.y) {
		if bounce {
			ball.vector.x = -ball.vector.x
		}
		return true
	}
	// right of brick collision
	if ball.vector.x < 0 && brick.Contains(ball.x-BallSize, ball.y) {
		if bounce {
			ball.vector.x = -ball.vector.x
		}
		return true
	}

	// left-bottom corner of the brick collision
	if ball.Contains(brick.x, brick.y+brick.height) {
		if bounce {
			ball.BounceFromPoint(brick.x, brick.y+brick.height)
		}
		return true
	}
	// right-bottom corner of the brick collision
	if ball.Contains(brick.x+brick.width, brick.y+brick.height) {
		if bounce {
			ball.BounceFromPoint(brick.x+brick.width, brick.y+brick.height)
		}
		return true
	}
	// left-top corner of the brick collision
	if ball.Contains(brick.x, brick.y) {
		if bounce {
			ball.BounceFromPoint(brick.x, brick.y)
		}
		return true
	}
	// right-top corner of the brick collision
	if ball.Contains(brick.x+brick.width, brick.y) {
		if bounce {
			ball.BounceFromPoint(brick.x+brick.width, brick.y)
		}
		return true
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
