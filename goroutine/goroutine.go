package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				//fmt.Printf("Hello From goroutine %d\n", i)
				//fmt.Printf("Hello From goroutine %d\n", a)
				//fmt.Println(a)
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println(a)
}
