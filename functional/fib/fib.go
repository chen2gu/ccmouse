package fib

type intGen func() int

func (i intGen) Read(p []byte) (n int, err error) {
	panic("implement me")
}

func Fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
