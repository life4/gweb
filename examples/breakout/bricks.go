package main

import (
	"fmt"

	"github.com/life4/gweb/canvas"
)

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
				context:   bricks.context,
				Rectangle: Rectangle{x: x, y: y, width: width, height: BrickHeight},
				cost:      cost,
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
