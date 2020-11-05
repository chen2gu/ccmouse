package main

import "fmt"

func main() {
	//fmt.Println("Hello,World!")
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{4, 5, 6, 7, 8}
	var grid1 [4][5]int
	var grid2 [4][5]bool

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(grid1)
	fmt.Println(grid2)

	fmt.Println("------->")
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}


	fmt.Println("------->")
	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	fmt.Println("------->")
	for _,k := range arr3 {
		fmt.Println(k)
	}

	fmt.Println("------->")
	for i,k := range arr3 {
		fmt.Println(i,k)
	}


}
