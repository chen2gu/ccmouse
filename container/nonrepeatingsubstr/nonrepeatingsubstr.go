package main

import "fmt"

//var last0ccurred = make([]int, 0xffff)

func length0fNonRepeatingSubStr(s string) int {
	//
	last0ccurred := make(map[rune]int)
	for i := range last0ccurred {
		last0ccurred[i] = 0
	}

	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {

		//if lastI, ok := last0ccurred[ch]; ok && lastI >= start {
		if lastI := last0ccurred[ch]; lastI >= start {
			start = lastI
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		last0ccurred[ch] = i + 1
	}
	return maxLength
}

func main() {
	fmt.Println(length0fNonRepeatingSubStr("abcabcaaacbb"))
	fmt.Println(length0fNonRepeatingSubStr("bbbbbbbbb"))
	fmt.Println(length0fNonRepeatingSubStr("pwwkew"))
	fmt.Println(length0fNonRepeatingSubStr(""))
	fmt.Println(length0fNonRepeatingSubStr("Yes我爱我的家乡!"))
	fmt.Println(length0fNonRepeatingSubStr("Yes我爱我家乡!"))
	fmt.Println(length0fNonRepeatingSubStr("Yes我爱家乡!"))

}
