package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func div(a, b int) (q, r int) {
	return a / b, a % b

	//q = a/b
	//r = a%b
	//return
}

func enal(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	//opName :=
	fmt.Printf("calling function %s with args "+
		"(%d, %d)\n", opName, a, b)
	return op(a, b)

}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))

}

func sum(number ...int) int {
	s := 0
	for i := range number {
		s += number[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	if result, err := enal(123, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	fmt.Println("------->")
	fmt.Println(enal(123, 4, "/s"))

	fmt.Println("------->")
	q, r := div(13, 4)
	fmt.Println(q, r)

	fmt.Println("------->")
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println("------->")
	fmt.Println(sum(1, 2, 3, 4))

	fmt.Println("------->")
	a, b := 3, 5
	a, b = swap(a, b)
	fmt.Println(a, b)

}
