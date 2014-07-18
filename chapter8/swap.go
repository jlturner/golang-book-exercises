package main
import "fmt"

func swap(a *int, b *int) {
	newA := *b
	newB := *a
	*a = newA
	*b = newB
	return
}

func main() {
	x := 10
	y := 23
	print := func() {
		fmt.Println(x, y)
	}

	print()
	swap(&x, &y)
	print()
}
