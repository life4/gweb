package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBounceFromPoint(t *testing.T) {
	ball := Ball{
		Point:  Point{x: 10, y: 10},
		vector: Vector{x: 5, y: 0},
		radius: 20,
	}

	// bounce from the right
	ball.BounceFromPoint(Point{x: 30, y: 10})
	assert.InDelta(t, ball.vector.x, -5, 0.0001, "bounce from the right: x")
	assert.InDelta(t, ball.vector.y, 0, 0.0001, "bounce from the right: y")

	// bounce from the left
	ball.vector = Vector{x: -5, y: 0}
	ball.BounceFromPoint(Point{x: -10, y: 10})
	assert.InDelta(t, ball.vector.x, 5, 0.0001, "bounce from the left: x")
	assert.InDelta(t, ball.vector.y, 0, 0.0001, "bounce from the left: y")

	// bounce from the bottom
	ball.vector = Vector{x: 0, y: 5}
	ball.BounceFromPoint(Point{x: 10, y: 30})
	assert.InDelta(t, ball.vector.x, 0, 0.0001, "bounce from the bottom: x")
	assert.InDelta(t, ball.vector.y, -5, 0.0001, "bounce from the bottom: y")
}
