package open_closed

import (
	"fmt"
	"math"
)

type shape interface {
	area() float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float32 {
	return s.sideLen * s.sideLen
}

func (t triangle) area() float32 {
	return t.base * t.height / 2
}

func (c *calculator) sumAreasGood(shapes ...shape) float32 {
	var sum float32

	for _, shape := range shapes {
		sum += shape.area()
	}
	return sum
}

func runGood() {
	c := circle{radius: 5}
	s := square{sideLen: 2}
	t := triangle{height: 10, base: 5}

	calc := calculator{}
	fmt.Println(calc.sumAreasGood(c, s, t))
}
