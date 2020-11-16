package main

import "fmt"

func length0fNonRepeatingSubStr(s string) int {
	last0ccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {

		if lastI, ok := last0ccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		last0ccurred[ch] = i
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
