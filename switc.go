package main

import "fmt"

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))

	}
	return g
}

func main() {
	fmt.Println(
		grade(10),
		grade(70),
		grade(80),
		grade(30),
		grade(100),
		grade(101),
		grade(-101),
	)

}
