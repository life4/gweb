package main

import (
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
	text  *TextBlock
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
		go bricks.text.DrawScore(bricks.score)
		go bricks.text.DrawHits(bricks.hits)

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

func (bricks *Bricks) Count() int {
	count := 0
	for _, brick := range bricks.registry {
		if !brick.removed {
			count += 1
		}
	}
	return count
}

func sign(n float64) float64 {
	if n >= 0 {
		return 1
	}
	return -1
}
