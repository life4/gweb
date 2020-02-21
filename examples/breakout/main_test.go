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
