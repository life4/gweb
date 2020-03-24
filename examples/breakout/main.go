package main

import (
	"sync"
	"time"

	"github.com/life4/gweb/web"
)

func main() {
	window := web.GetWindow()
	doc := window.Document()
	doc.SetTitle("Breakout")
	body := doc.Body()

	// create canvas
	h := window.InnerHeight() - 40
	w := window.InnerWidth() - 40
	canvas := doc.CreateCanvas()
	canvas.SetHeight(h)
	canvas.SetWidth(w)
	body.Node().AppendChild(canvas.Node())

	context := canvas.Context2D()

	// draw background
	context.SetFillStyle(BGColor)
	context.BeginPath()
	context.Rectangle(0, 0, w, h).Filled().Draw()
	context.Fill()
	context.ClosePath()

	// make handlers
	rect := Rectangle{
		x:      w / 2,
		y:      h - 60,
		width:  PlatformWidth,
		height: PlatformHeight,
	}
	platformCicrle := CircleFromRectangle(rect)
	platform := Platform{
		rect:         &rect,
		circle:       &platformCicrle,
		context:      context,
		element:      canvas,
		mouseX:       w / 2,
		windowWidth:  w,
		windowHeight: h,
	}
	block := TextBlock{context: context, updated: time.Now()}
	ballCircle := Circle{
		x:      platform.circle.x,
		y:      platform.rect.y - BallSize - 5,
		radius: BallSize,
	}
	ball := Ball{
		context:     context,
		vector:      Vector{x: 5, y: -5},
		Circle:      ballCircle,
		windowWidth: w, windowHeight: h,
		platform: &platform,
	}
	bricks := Bricks{
		context:      context,
		windowWidth:  w,
		windowHeight: h,
		ready:        false,
		text:         &block,
	}
	go bricks.Draw()

	// register mouse movement handler
	body.EventTarget().Listen(web.EventTypeMouseMove, platform.handleMouse)

	// register frame updaters
	handler := func() {
		wg := sync.WaitGroup{}
		wg.Add(4)
		go func() {
			// update FPS
			block.handle()
			wg.Done()
		}()
		go func() {
			// update platform position
			platform.handleFrame()
			wg.Done()
		}()
		go func() {
			// check if the ball should bounce from a brick
			bricks.Handle(&ball)
			wg.Done()
		}()
		go func() {
			// check if the ball should bounce from border or platform
			ball.handle()
			wg.Done()
		}()
		wg.Wait()
	}
	window.RequestAnimationFrame(handler, true)

	// prevent ending of the program
	select {}
}
