package liskov_substitution

import "fmt"

func (b Bird) Fly() string {
	return "I can fly"
}

type Penguin struct{}

func (p Penguin) Speak() string {
	return "Squawk"
}

func MakeBirdFly(b Bird) {
	fmt.Println(b.Fly())
}

func runBad() {
	var birds []Bird

	birds = append(birds, Bird{})
	birds = append(birds, Penguin{}) // エラー: PenguinはFlyメソッドを持たないためBirdの代わりにならない

	for _, bird := range birds {
		MakeBirdFly(bird)
	}
}
