package liskov_substitution

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof"
}

func MakeAnimalSpeak(a Animal) {
	fmt.Println(a.Speak())
}

func runGood() {
	var animals []Animal

	animals = append(animals, Bird{})
	animals = append(animals, Dog{})

	for _, animal := range animals {
		MakeAnimalSpeak(animal)
	}
}
