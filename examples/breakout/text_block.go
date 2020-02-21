package main

import (
	"fmt"
	"time"

	"github.com/life4/gweb/canvas"
)

type TextBlock struct {
	context canvas.Context2D
	updated time.Time
}

func (block TextBlock) drawFPS(now time.Time) {
	// calculate FPS
	fps := time.Second / now.Sub(block.updated)
	text := fmt.Sprintf("%d FPS", int64(fps))

	// clear
	block.context.SetFillStyle(BGColor)
	block.context.Rectangle(TextLeft, TextTop, TextWidth, TextHeight+TextMargin).Filled().Draw()

	// write
	block.context.Text().SetFont(fmt.Sprintf("bold %dpx Roboto", TextHeight))
	block.context.SetFillStyle(TextColor)
	block.context.Text().Fill(text, TextLeft, TextTop+TextHeight, TextWidth)
}

func (block *TextBlock) handle() {
	now := time.Now()
	// update FPS counter every second
	if block.updated.Second() != now.Second() {
		block.drawFPS(now)
	}
	block.updated = now
}

func (block TextBlock) drawText(text string, row int) {
	y := TextTop + row*(TextMargin+TextHeight)

	// clear place where previous score was
	block.context.SetFillStyle(BGColor)
	block.context.Rectangle(TextLeft, y, TextRight, TextHeight+TextWidth).Filled().Draw()

	// draw the score
	block.context.SetFillStyle(TextColor)
	block.context.Text().SetFont(fmt.Sprintf("bold %dpx Roboto", TextHeight))
	block.context.Text().Fill(text, TextLeft, y+TextHeight, TextWidth)
}

func (block TextBlock) DrawScore(score int) {
	// make text
	var text string
	if score == 1 {
		text = fmt.Sprintf("%d point", score)
	} else {
		text = fmt.Sprintf("%d points", score)
	}
	block.drawText(text, 1)
}

func (block TextBlock) DrawHits(hits int) {
	// make text
	var text string
	if hits == 1 {
		text = fmt.Sprintf("%d hit", hits)
	} else {
		text = fmt.Sprintf("%d hits", hits)
	}
	block.drawText(text, 2)
}
