package main
import "fmt"

func fib (n int) (output int) {
	switch n {
	case 0: output = 0
	case 1: output = 1
	default: output = fib(n - 1) + fib(n - 2)
	}
	return
}

func main () {
	for i := 0; i < 20; i++ {
		fmt.Println(fib(i))
	}
}
