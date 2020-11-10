package main

import (
	"fmt"

	"github.com/chen2gu/ccmouse/queue"
)

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEMpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEMpty())

}
