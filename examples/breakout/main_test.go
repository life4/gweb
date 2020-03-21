package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorRotate(t *testing.T) {
	f := func(gx, gy float64, angle float64, ex, ey float64) {
		v := Vector{x: gx, y: gy}
		actual := v.Rotate(angle)
		assert.InDelta(t, actual.x, ex, 0.0001)
		assert.InDelta(t, actual.y, ey, 0.0001)
	}
	f(10, 10, math.Pi, -10, -10)
	f(10, 10, 2*math.Pi, 10, 10)
	f(10, 10, math.Pi/2, -10, 10)
	f(10, 10, math.Pi*3/2, 10, -10)
}

func TestCircleFromRectangleRadius(t *testing.T) {
	f := func(w, h, expected int) {
		circle := CircleFromRectangle(Rectangle{
			width:  w,
			height: h,
		})
		if expected != 0 {
			assert.Equal(t, circle.radius, expected)
		}
		assert.GreaterOrEqual(t, circle.radius, w/2)
	}

	f(10, 5, 5)
	f(80, 30, 41)
	f(10, 5, 0)
	f(5, 10, 0)
}
