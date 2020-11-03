package main

import (
	"fmt"
)

func enal(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation: " + op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b

	//q = a/b
	//r = a%b
	//return
}

func main() {
	fmt.Println(enal(3, 4, "+"))
	q, r := div(13, 4)
	fmt.Println(q, r)
}
