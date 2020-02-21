package main

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/life4/gweb/canvas"
	"github.com/life4/gweb/web"
)

const (
	BGColor       = "#ecf0f1"
	BallColor     = "#27ae60"
	PlatformColor = "#2c3e50"
	TextColor     = "#2c3e50"
)

const PlatformWidth = 120
const PlatformHeight = 20
const PlatformMaxSpeed = 40

const BallSize = 20

const (
	BrickHeight     = 20
	BrickRows       = 8
	BrickCols       = 14
	BrickMarginLeft = 120 // pixels
	BrickMarginTop  = 10  // pixels
	BrickMarginX    = 5   // pixels
	BrickMarginY    = 5   // pixels
)

const (
	TextWidth  = 90
	TextHeight = 20
	TextLeft   = 10
	TextTop    = 10
	TextMargin = 5

	TextBottom = TextTop + (TextHeight+TextMargin)*3
	TextRight  = TextLeft + TextWidth
)

func sign(n float64) float64 {
	if n >= 0 {
		return 1
	}
	return -1
}

type Point struct{ x, y int }
type Vector struct{ x, y float64 }
type Rectangle struct{ x, y, width, height int }

type Ball struct {
	Point
	vector Vector

	windowWidth  int
	windowHeight int

	context  canvas.Context2D
	platform *Platform
}

func (ball Ball) Contains(x, y int) bool {
	hypotenuse := math.Pow(float64(BallSize), 2)
	cathetus1 := math.Pow(float64(x-ball.x), 2)
	cathetus2 := math.Pow(float64(y-ball.y), 2)
	return cathetus1+cathetus2 < hypotenuse
}

func (ball *Ball) BounceFromPoint(x, y int) {
	katX := float64(x - ball.x)
	katY := float64(y - ball.y)
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

	// bounce from left-top corner of the platform
	if ctx.Contains(ctx.platform.x, ctx.platform.y) {
		ctx.BounceFromPoint(ctx.platform.x, ctx.platform.y)
		return
	}
	// bounce from right-top corner of the platform
	if ctx.Contains(ctx.platform.x+ctx.platform.width, ctx.platform.y) {
		ctx.BounceFromPoint(ctx.platform.x+ctx.platform.width, ctx.platform.y)
		return
	}
	// bounce from left-bottom corner of the platform
	if ctx.Contains(ctx.platform.x, ctx.platform.y+PlatformHeight) {
		ctx.BounceFromPoint(ctx.platform.x, ctx.platform.y+PlatformHeight)
		return
	}
	// bounce from right-bottom corner of the platform
	if ctx.Contains(ctx.platform.x+ctx.platform.width, ctx.platform.y+PlatformHeight) {
		ctx.BounceFromPoint(ctx.platform.x+ctx.platform.width, ctx.platform.y+PlatformHeight)
		return
	}
}

func (ctx *Ball) handle() {
	// clear out previous render
	ctx.context.SetFillStyle(BGColor)
	ctx.context.BeginPath()
	ctx.context.Arc(ctx.x, ctx.y, BallSize+1, 0, math.Pi*2)
	ctx.context.Fill()
	ctx.context.ClosePath()

	ctx.changeDirection()

	// move the ball
	ctx.x += int(math.Round(ctx.vector.x))
	ctx.y += int(math.Round(ctx.vector.y))

	// draw the ball
	ctx.context.SetFillStyle(BallColor)
	ctx.context.BeginPath()
	ctx.context.Arc(ctx.x, ctx.y, BallSize, 0, math.Pi*2)
	ctx.context.Fill()
	ctx.context.ClosePath()
}

type Brick struct {
	context canvas.Context2D
	x, y    int
	cost    int
	width   int
	height  int
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

type Bricks struct {
	context      canvas.Context2D
	registry     []*Brick
	ready        bool
	windowWidth  int
	windowHeight int

	// stat
	score int
	hits  int
}

func (bricks *Bricks) Draw() {
	bricks.registry = make([]*Brick, BrickCols*BrickRows)
	width := (bricks.windowWidth-BrickMarginLeft)/BrickCols - BrickMarginX
	colors := [...]string{"#c0392b", "#d35400", "#f39c12", "#f1c40f"}
	costs := [...]int{7, 5, 3, 1}
	for i := 0; i < BrickCols; i++ {
		for j := 0; j < BrickRows; j++ {
			x := BrickMarginLeft + (width+BrickMarginX)*i
			y := BrickMarginTop + (BrickHeight+BrickMarginY)*j
			color := colors[(j/2)%len(colors)]
			cost := costs[(j/2)%len(colors)]

			brick := Brick{
				context: bricks.context,
				x:       x,
				y:       y,
				cost:    cost,
				width:   width,
				height:  BrickHeight,
			}
			brick.Draw(color)
			bricks.registry[BrickRows*i+j] = &brick
		}
	}
	bricks.ready = true
}

func (bricks Bricks) drawText(text string, row int) {
	y := TextTop + row*(TextMargin+TextHeight)

	// clear place where previous score was
	bricks.context.SetFillStyle(BGColor)
	bricks.context.Rectangle(TextLeft, y, TextRight, TextHeight+TextWidth).Filled().Draw()

	// draw the score
	bricks.context.SetFillStyle(TextColor)
	bricks.context.Text().SetFont(fmt.Sprintf("bold %dpx Roboto", TextHeight))
	bricks.context.Text().Fill(text, TextLeft, y+TextHeight, TextWidth)
}

func (bricks Bricks) drawScore() {
	// make text
	var text string
	if bricks.score == 1 {
		text = fmt.Sprintf("%d point", bricks.score)
	} else {
		text = fmt.Sprintf("%d points", bricks.score)
	}
	bricks.drawText(text, 1)
}

func (bricks Bricks) drawHits() {
	// make text
	var text string
	if bricks.hits == 1 {
		text = fmt.Sprintf("%d hit", bricks.hits)
	} else {
		text = fmt.Sprintf("%d hits", bricks.hits)
	}
	bricks.drawText(text, 2)
}

func (bricks *Bricks) Handle(ball *Ball) {
	if !bricks.ready {
		return
	}
	changed := false
	for _, brick := range bricks.registry {
		// we bounce the ball only on first collision with a brick in a frame
		if !brick.Collide(ball, !changed) {
			continue
		}
		// if the ball touched the brick, remove the brick and count score
		brick.Remove()
		bricks.score += brick.cost
		bricks.hits += 1
		changed = true
	}
	if changed {
		// re-draw stat
		go bricks.drawScore()
		go bricks.drawHits()

		// speed up ball after some hits
		speedUpHits := [...]int{4, 8, 16, 24, 32, 64}
		for _, hits := range speedUpHits {
			if bricks.hits == hits {
				ball.vector.x += sign(ball.vector.x) * 1
				ball.vector.y += sign(ball.vector.y) * 1
				break
			}
		}
	}
}

type FPS struct {
	context canvas.Context2D
	updated time.Time
}

func (h *FPS) drawFPS(now time.Time) {
	// calculate FPS
	fps := time.Second / now.Sub(h.updated)
	text := fmt.Sprintf("%d FPS", int64(fps))

	// clear
	h.context.SetFillStyle(BGColor)
	h.context.Rectangle(TextLeft, TextTop, TextWidth, TextHeight+TextMargin).Filled().Draw()

	// write
	h.context.Text().SetFont(fmt.Sprintf("bold %dpx Roboto", TextHeight))
	h.context.SetFillStyle(TextColor)
	h.context.Text().Fill(text, TextLeft, TextTop+TextHeight, TextWidth)
}

func (h *FPS) handle() {
	now := time.Now()
	// update FPS counter every second
	if h.updated.Second() != now.Second() {
		h.drawFPS(now)
	}
	h.updated = now
}

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
	platform := Platform{
		Rectangle: Rectangle{
			x:      w / 2,
			y:      h - 60,
			width:  PlatformWidth,
			height: PlatformHeight,
		},
		context:      context,
		element:      canvas,
		mouseX:       w / 2,
		windowWidth:  w,
		windowHeight: h,
	}
	fps := FPS{context: context, updated: time.Now()}
	ball := Ball{
		context: context,
		vector: Vector{x: 5, y: -5},
		Point: Point{x: platform.x, y:           platform.y - BallSize,},
		windowWidth: w, windowHeight: h,
		platform: &platform,
	}
	bricks := Bricks{
		context:      context,
		windowWidth:  w,
		windowHeight: h,
		ready:        false,
	}
	go bricks.Draw()

	// register mouse movement handler
	body.EventTarget().Listen(web.EventTypeMouseMove, platform.handleMouse)

	// register frame updaters
	handler := func() {
		wg := sync.WaitGroup{}
		wg.Add(4)
		go func() {
			fps.handle()
			wg.Done()
		}()
		go func() {
			platform.handleFrame()
			wg.Done()
		}()
		go func() {
			bricks.Handle(&ball)
			wg.Done()
		}()
		go func() {
			ball.handle()
			wg.Done()
		}()
		wg.Wait()
	}
	window.RequestAnimationFrame(handler, true)

	// prevent ending of the program
	select {}
}

