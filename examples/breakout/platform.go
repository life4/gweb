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

func (pl Platform) Angle() float64 {
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
	ctx.circle.x += path
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
		1.5*math.Pi-pl.Angle()-(float64(delta)/math.Pi/2),
		1.5*math.Pi+pl.Angle()+(float64(delta)/math.Pi/2),
	)
	pl.context.Fill()
	pl.context.ClosePath()
}
