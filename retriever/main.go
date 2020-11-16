package main

import (
	"fmt"
	"github.com/chen2guo/ccmouse/retriever/mock"
	"github.com/chen2guo/ccmouse/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")

}

func inspect(r Retriever) {

	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)

	}

}

func main() {

	fmt.Println("1.----->")
	var r Retriever

	r = &mock.Retriever{"This is a fake com."}
	fmt.Printf("%T %v\n", r, r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)
	//fmt.Println(download(r))

	inspect(r)

	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever.")
	}

	fmt.Println("2.----->")
	fmt.Println(download(mock.Retriever{
		"This is a fake com."}))

}
