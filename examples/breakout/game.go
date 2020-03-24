package main

import (
	"sync"
	"time"

	"github.com/life4/gweb/web"
)

type Game struct {
	Width  int
	Height int
	Window web.Window
	Canvas web.Canvas
	Body   web.HTMLElement

	state    *State
	platform Platform
	ball     Ball
	block    TextBlock
	bricks   Bricks
}

func (game *Game) Init() {
	game.state = &State{Stop: SubState{}}
	context := game.Canvas.Context2D()

	// draw background
	context.SetFillStyle(BGColor)
	context.BeginPath()
	context.Rectangle(0, 0, game.Width, game.Height).Filled().Draw()
	context.Fill()
	context.ClosePath()

	// make handlers
	rect := Rectangle{
		x:      game.Width / 2,
		y:      game.Height - 60,
		width:  PlatformWidth,
		height: PlatformHeight,
	}
	platformCicrle := CircleFromRectangle(rect)
	game.platform = Platform{
		rect:         &rect,
		circle:       &platformCicrle,
		context:      context,
		element:      game.Canvas,
		mouseX:       game.Width / 2,
		windowWidth:  game.Width,
		windowHeight: game.Height,
	}
	game.block = TextBlock{context: context, updated: time.Now()}
	ballCircle := Circle{
		x:      game.platform.circle.x,
		y:      game.platform.rect.y - BallSize - 5,
		radius: BallSize,
	}
	game.ball = Ball{
		context:      context,
		vector:       Vector{x: 5, y: -5},
		Circle:       ballCircle,
		windowWidth:  game.Width,
		windowHeight: game.Height,
		platform:     &game.platform,
	}
	game.bricks = Bricks{
		context:      context,
		windowWidth:  game.Width,
		windowHeight: game.Height,
		ready:        false,
		text:         &game.block,
	}
	go game.bricks.Draw()
}

func (game *Game) handler() {
	if game.state.Stop.Requested {
		game.state.Stop.Done()
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		// update FPS
		game.block.handle()
		wg.Done()
	}()
	go func() {
		// update platform position
		game.platform.handleFrame()
		wg.Done()
	}()
	go func() {
		// check if the ball should bounce from a brick
		game.bricks.Handle(&game.ball)
		wg.Done()
	}()
	go func() {
		// check if the ball should bounce from border or platform
		game.ball.handle()
		wg.Done()
	}()
	wg.Wait()

	game.Window.RequestAnimationFrame(game.handler, false)
}

func (game *Game) Register() {
	// register mouse movement handler
	game.Body.EventTarget().Listen(web.EventTypeMouseMove, game.platform.handleMouse)
	// register frame updaters
	game.Window.RequestAnimationFrame(game.handler, false)
}

func (game *Game) Stop() {
	game.state.Stop.Request()
	game.state.Stop.Wait()
}
