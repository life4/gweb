package main

import "math"

type Point struct{ x, y int }

type Rectangle struct{ x, y, width, height int }

func (rectangle Rectangle) Contains(point Point) bool {
	if point.y < rectangle.y { // point upper
		return false
	}
	if point.y > rectangle.y+rectangle.height { // point downer
		return false
	}
	if point.x > rectangle.x+rectangle.width { // point righter
		return false
	}
	if point.x < rectangle.x { // point lefter
		return false
	}
	return true

}

type Circle struct{ x, y, radius int }

func (circle Circle) Contains(point Point) bool {
	hypotenuse := math.Pow(float64(circle.radius), 2)
	cathetus1 := math.Pow(float64(point.x-circle.x), 2)
	cathetus2 := math.Pow(float64(point.y-circle.y), 2)
	return cathetus1+cathetus2 < hypotenuse
}

func CircleFromRectangle(rect Rectangle) Circle {
	base := math.Sqrt(math.Pow(float64(rect.width)/2, 2) + math.Pow(float64(rect.height), 2))
	cos := float64(rect.height) / base
	radius := int(base / 2 / cos)

	return Circle{
		x:      rect.x + rect.width/2,
		y:      rect.y + radius,
		radius: radius,
	}
}
