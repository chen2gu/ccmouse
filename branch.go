package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	const filename = "go.mod"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	//fmt.Printf("%s\n",contents)   //error

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println("------->")

}
