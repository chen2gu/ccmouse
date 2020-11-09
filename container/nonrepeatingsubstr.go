package main

import "fmt"

func length0fNonRepeatingSubStr(s string) int {
	last0ccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {

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

}
