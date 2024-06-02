package liskov_substitution

type Bird struct{}

func (b Bird) Speak() string {
	return "Tweet"
}
