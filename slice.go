package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])
	//fmt.Println("arr[:] = ",arr[:])

	fmt.Println()

	s1 := arr[2:]
	fmt.Println("s1 = ", s1)
	s2 := arr[:]
	fmt.Println("s2 = ", s2)

	fmt.Println("S1 ------->")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println()
	fmt.Println("S2 ------->")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

}
