package interface_segregation

import "math"

type shapeGood interface {
	area() float64
}

type object interface {
	shapeGood
	volume() float64
}

func (s square) areaGood() float64 {
	return math.Pow(s.sideLen, 2)
}

func (c cubeGood) volume() float64 {
	return math.Pow(c.sideLen, 3)
}
