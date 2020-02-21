package main

import "math"

type Vector struct {
	x, y float64
}

func (vector *Vector) Rotate(angle float64) Vector {
	sin := math.Sin(angle)
	cos := math.Cos(angle)
	return Vector{
		x: vector.x*cos - vector.y*sin,
		y: vector.x*sin + vector.y*cos,
	}
}

func (vector Vector) Len() float64 {
	return math.Sqrt(math.Pow(vector.x, 2) + math.Pow(vector.y, 2))
}

func (vector Vector) Angle(other Vector) float64 {
	return vector.Dot(other) / (vector.Len() * other.Len())
}

func (vector Vector) Dot(other Vector) float64 {
	return vector.x*other.x + vector.y*other.y
}

func (vector Vector) Sub(other Vector) Vector {
	return Vector{x: vector.x - other.x, y: vector.y - other.y}
}

func (vector Vector) Mul(value float64) Vector {
	return Vector{x: vector.x * value, y: vector.y * value}
}

func (vector Vector) Normalized() Vector {
	value := 1.0 / vector.Len()
	return vector.Mul(value)
}
