package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱我的家乡!"
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Printf("", []byte(s))
	fmt.Println()
	fmt.Printf("%X\n", []byte(s))
	for _, v := range []byte(s) {
		fmt.Printf("%X ", v)
	}
	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Ruune count: ", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d  %c) ", i, ch)
	}

	fmt.Println()
}
