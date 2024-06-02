package interface_segregation

type square struct {
	sideLen float64
}

type cubeBad struct {
	sideLen float64
}

type cubeGood struct {
	square
}
