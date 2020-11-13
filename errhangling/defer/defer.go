package main

import (
	"bufio"
	"fmt"
	"github.com/chen2gu/ccmouse/functional/fib"
	"os"
)

func tyeDefer()  {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred.lsz")
	fmt.Println(4)
}

func tyeDefer1() {
	for i := 0;i<100; i++ {
		defer fmt.Println(i)
		if i==30 {
			panic("error.")
		}
	}
}

func writerFile(filename string) {
	file, err:=os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i :=0; i<20;i++{
		fmt.Fprintln(writer,f())
	}
}

func main() {
	//tyeDefer()
	//writerFile("fib.txt")
	tyeDefer1()
}
