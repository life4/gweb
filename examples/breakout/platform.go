package main

import (
	"github.com/life4/gweb/canvas"
	"github.com/life4/gweb/web"
)

type Platform struct {
	Rectangle
	context canvas.Context2D
	element web.Canvas
	// movement
	mouseX int
	// borders
	windowWidth  int
	windowHeight int
}

func (platform Platform) Contains(x, y int) bool {
	if y < platform.y { // ball upper
		return false
	}
	if y > platform.y+platform.height { // ball downer
		return false
	}
	if x > platform.x+platform.width { // ball righter
		return false
	}
	if x < platform.x { // ball lefter
		return false
	}
	return true

}

func (ctx *Platform) changePosition() {
	path := ctx.mouseX - (ctx.x + ctx.width/2)
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
	if ctx.x+path <= 0 {
		ctx.x = 0
		return
	}
	if ctx.x+path >= ctx.windowWidth-ctx.width {
		ctx.x = ctx.windowWidth - ctx.width
		return
	}

	ctx.x += path
}

func (platform *Platform) handleMouse(event web.Event) {
	platform.mouseX = event.Get("clientX").Int()
}

func (ctx *Platform) handleFrame() {
	// clear out previous render
	ctx.context.SetFillStyle(BGColor)
	ctx.context.Rectangle(ctx.x, ctx.y, ctx.width, ctx.height).Filled().Draw()

	// change platform coordinates
	ctx.changePosition()

	// draw the platform
	ctx.context.SetFillStyle(PlatformColor)
	ctx.context.Rectangle(ctx.x, ctx.y, ctx.width, ctx.height).Filled().Draw()
}
