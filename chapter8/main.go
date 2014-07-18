package main
import "fmt"

func zero(xPointer *int) {
	*xPointer = 0
}

func numberBuilder(number int) (numberFunc func(*int)) {
	numberFunc = func(intPointer *int) {
		*intPointer = number
	}
	return
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)

	twelve := numberBuilder(12)
	newPointer := new(int)
	twelve(newPointer)
	fmt.Println(*newPointer)
}
