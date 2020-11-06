package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	scanner := bufio.NewScanner(file)
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
	forever()

}
