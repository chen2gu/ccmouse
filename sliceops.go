package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v,len(s1)=%d,cap(s1)=%d\n", s, len(s), cap(s))
	//fmt.Printf("len(s)=%d,cap(s)=%d\n",  len(s), cap(s))
}

func main() {
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

}
