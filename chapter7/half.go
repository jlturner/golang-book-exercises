package main
import "fmt"

func half (input int) (output int, even bool) {
	output = input / 2
	if input % 2 == 0 {
		even = true
	} else {
		even = false
	}
	return
}

func main () {
	for i := 0; i <= 10; i++ {
		output, even := half(i)
		fmt.Println(output, even)
	}
}
