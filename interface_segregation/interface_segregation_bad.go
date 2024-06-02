package interface_segregation

import "math"

type shapeBad interface {
	area() float64
	volume() float64
}

func (s square) areaBad() float64 {
	return s.sideLen * s.sideLen
}

func (s square) volume() float64 {
	return 0
}

func (c cubeBad) area() float64 {
	return math.Pow(c.sideLen, 2)
}

func (c cubeBad) volume() float64 {
	return math.Pow(c.sideLen, 3)
}
