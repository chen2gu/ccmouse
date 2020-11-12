package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func converToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result

}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
		time.Sleep(time.Second * 3)
	}

}

func main() {
	fmt.Println(
		converToBin(5),  //11
		converToBin(13), //1011 -->1101
		converToBin(111111),
		converToBin(0),
	)

	printFile("go.mod")

	s := `
		asda
		asda
		asfas       fff
		aaa
		`

	printFileContents(strings.NewReader(s))

	//forever()

}
