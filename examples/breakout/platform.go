package main

import (
	"math"

	"github.com/life4/gweb/canvas"
	"github.com/life4/gweb/web"
)

type Platform struct {
	circle *Circle
	rect   *Rectangle

	context canvas.Context2D
	element web.Canvas
	// movement
	mouseX int
	// borders
	windowWidth  int
	windowHeight int
}

func (pl Platform) Contains(point Point) bool {
	return pl.circle.Contains(point) && pl.rect.Contains(point)
}

// Touch returns touch point of platform and ball if any
func (pl Platform) Touch(ball Ball) *Point {
	point := pl.touchInside(ball)
	if point != nil {
		return point
	}
	point = pl.touchUp(ball)
	if point != nil {
		return point
	}
	return pl.touchCorners(ball)
}

func (pl Platform) touchInside(ball Ball) *Point {
	// don't bounce if ball moves up
	if ball.vector.y < -1.0 {
		return nil
	}

	point := &Point{x: ball.x, y: ball.y}
	if pl.Contains(*point) {
		point.y = ball.y + ball.radius
		return point
	}
	return nil
}

func (pl Platform) touchUp(ball Ball) *Point {
	// don't bounce if ball is inside of the platform
	if ball.y > pl.circle.y {
		return nil
	}
	// don't bounce if ball moves up
	if ball.vector.y < -1.0 {
		return nil
	}

	catx := float64(ball.x - pl.circle.x)
	caty := float64(ball.y - pl.circle.y)

	// check if ball is too far from platform circle
	hypotenuse := math.Sqrt(math.Pow(catx, 2) + math.Pow(caty, 2))
	distance := math.Abs(float64(ball.radius + pl.circle.radius))
	if hypotenuse > distance {
		return nil
	}

	// touches the upper side of the platform
	ratio := float64(ball.radius) / float64(pl.circle.radius)
	point := Point{
		x: ball.x - int(catx*ratio),
		y: ball.y - int(caty*ratio),
	}
	if point.y >= pl.rect.y+pl.rect.height {
		return nil
	}
	return &point
}

func (pl Platform) touchCorners(ball Ball) *Point {
	// left
	if ball.vector.x > 0 {
		point := Point{
			x: pl.rect.x,
			y: pl.rect.y + pl.rect.height,
		}
		if ball.Contains(point) {
			return &point
		}
	}

	// right
	if ball.vector.x < 0 {
		point := Point{
			x: pl.rect.x + pl.rect.width,
			y: pl.rect.y + pl.rect.height,
		}
		if ball.Contains(point) {
			return &point
		}
	}
	return nil
}

func (pl Platform) angle() float64 {
	tan := float64(pl.rect.width/2) / float64(pl.circle.radius-pl.rect.height)
	return math.Atan(tan)
}

func (ctx *Platform) changePosition() {
	path := ctx.mouseX - (ctx.rect.x + ctx.rect.width/2)
	if path == 0 {
		return
	}

	// don't move too fast
	if path > 0 && path > PlatformMaxSpeed {
		path = PlatformMaxSpeed
	} else if path < 0 && path < -PlatformMaxSpeed {
		path = -PlatformMaxSpeed
	}

	// don't move out of playground
	if ctx.rect.x+path <= 0 {
		ctx.rect.x = 0
		return
	}
	if ctx.rect.x+path >= ctx.windowWidth-ctx.rect.width {
		ctx.rect.x = ctx.windowWidth - ctx.rect.width
		return
	}

	ctx.rect.x += path
	ctx.circle.x = ctx.rect.x + ctx.rect.width/2
}

func (platform *Platform) handleMouse(event web.Event) {
	platform.mouseX = event.Get("clientX").Int()
}

func (ctx *Platform) handleFrame() {
	// clear out previous render
	ctx.draw(BGColor, 1)

	// change platform coordinates
	ctx.changePosition()

	// draw the platform
	ctx.draw(PlatformColor, 0)
}

func (pl Platform) draw(color string, delta int) {
	// pl.context.SetFillStyle("red")
	// pl.context.Rectangle(pl.circle.x, pl.rect.y+pl.circle.radius, 2, 2).Filled().Draw()
	// pl.context.SetFillStyle("green")
	// pl.context.Rectangle(pl.circle.x, pl.circle.y, 2, 2).Filled().Draw()

	pl.context.SetFillStyle(color)
	pl.context.BeginPath()
	pl.context.Arc(
		pl.circle.x, pl.circle.y,
		pl.circle.radius+delta,
		1.5*math.Pi-pl.angle()-(float64(delta)/math.Pi/2),
		1.5*math.Pi+pl.angle()+(float64(delta)/math.Pi/2),
	)
	pl.context.Fill()
	pl.context.ClosePath()
}
