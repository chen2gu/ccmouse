package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v,len(s)=%d,cap(s)=%d\n", s, len(s), cap(s))
	//fmt.Printf("len(s)=%d,cap(s)=%d\n",  len(s), cap(s))
}

func main() {
	fmt.Println("Init slice.")
	var s []int
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}

	fmt.Println(s)

	fmt.Println("S1 ------->")
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	fmt.Println("S2 ------->")
	s2 := make([]int, 16)
	printSlice(s2)

	fmt.Println("S3 ------->")
	s3 := make([]int, 10, 32)
	printSlice(s3)

	fmt.Println("Copying slice------->")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements slice------->")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front------->")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back------->")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)
}
