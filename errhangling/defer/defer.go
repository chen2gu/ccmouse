package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/chen2gu/ccmouse/functional/fib"
	"os"
)

func tyeDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred.lsz")
	fmt.Println(4)
}

func tyeDefer1() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("error.")
		}
	}
}

func wr1iterFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0644)

	err = errors.New("This is a custom error.")

	if err != nil {
		//panic(err)
		//fmt.Println("Error: ",err.Error())

		if PathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s %s %sy.\n", PathError.Op,
				PathError.Path,
				PathError.Err)
		}
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tyeDefer()
	wr1iterFile("fib.txt")
	//tyeDefer1()
}
