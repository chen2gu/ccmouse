package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)

	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)

	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func triangle() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))

	fmt.Println(a, b, c)
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt((a*a + b*b)))

	fmt.Println(filename, a, b, c)
}

func enums() {
	const (
		cpp    = 0
		java   = 1
		pyhton = 2
		golang = 3
	)
	fmt.Println(cpp, java, pyhton, golang)
}

func enums1() {
	const (
		cpp = iota
		java
		pyhton
		_
		golang
	)
	fmt.Println(cpp, java, pyhton, golang)
}

func enums2() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
	euler()
	fmt.Println("------->")
	triangle()
	fmt.Println("------->")
	consts()
	fmt.Println("------->")
	enums()
	fmt.Println("------->")
	enums1()

	fmt.Println("------->")
	enums2()
}
