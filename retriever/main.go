package main

import (
	"fmt"
	"github.com/chen2gu/ccmouse/retriever/mock"
	"github.com/chen2gu/ccmouse/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")

}
func main() {

	fmt.Println("1.----->")
	var r Retriever
	r = mock.Retriever{"This is a fake com."}
	fmt.Println(download(r))

	fmt.Println("2.----->")
	fmt.Println(download(mock.Retriever{
		"This is a fake com."}))

	fmt.Println("3.----->")
	//var r Retriever
	r = mock.Retriever{"This is a fake com."}
	r = real.Retriever{}
	fmt.Println(download(r))

}
