package main

import (
	"fmt"
	"math"
)

var (
	a = 01
	s = "hello"
)

func variableZeroValue() {
	var a int
	var s string

	fmt.Println(a, s)
	fmt.Printf("%d    %q\n", a, s)
}

func varibleInitalValue() {
	var b, c int = 11, 221
	var s1 string = "abc" +
		"hahhaha"
	fmt.Println(b, c, s1)

}

func varibleTypeDedution1() {
	var b, c = 11, 222
	var s1 = "abc" +
		"hahhaha"
	fmt.Println(b, c, s1)

}

func varibleTypeDedution2() {
	b, c := 11, 223
	s1 := "abc" +
		"hahhaha"
	fmt.Println(b, c, s1)
}

func triangle() {
	var a, b int = 3, 4

	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c

}

func varibleTypeDedution3() {
	b, c, s1, t1 := 11, 224, true, "sixsixsxi"
	b = 909
	fmt.Println(b, c, s1, t1)
}

func main() {
	fmt.Println("Hello,World!!")
	fmt.Println("Hello,World!!")
	fmt.Println("Hello,World!!")

	variableZeroValue()

	varibleInitalValue()

	varibleTypeDedution1()

	varibleTypeDedution2()

	varibleTypeDedution3()

	fmt.Println(a, s)

	//	// step1: cd 到项目下
	//
	//	// step2: 移除.idea
	//	$ git rm --cached -r .idea
	//	rm '.idea/misc.xml'
	//rm '.idea/modules.xml'
	//rm '.idea/vcs.xml'
	//rm '.idea/workspace.xml'
	//
	//	// step3:
	//	git add .
	//		git commit -am "rm idea"

}
