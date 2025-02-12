package open_closed

import (
	"fmt"
	"math"
)

func (c calculator) sumAreasBad(shapes ...interface{}) float32 {
	var sum float32

	for _, shape := range shapes {
		switch shape.(type) {
		case circle:
			r := shape.(circle).radius
			sum += math.Pi * r * r
		case square:
			l := shape.(square).sideLen
			sum += l * l
		}
	}

	return sum
}

func runBad() {
	c := circle{radius: 5}
	s := square{sideLen: 2}
	calc := calculator{}
	fmt.Println("total of areas:", calc.sumAreasBad(c, s))
}
