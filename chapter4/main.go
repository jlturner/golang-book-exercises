package main
import "fmt"

var z int = 123

func main() {
	var x string = "Hello!"
	var y = 32
	fmt.Println(x)
	fmt.Println(y)
	y = 3930
	fmt.Println(y)
	g := "oh no"
	fmt.Println(g)
	g = "yes"
	fmt.Println(g)
	wow := 23423
	fmt.Println(wow)
	dogsName := "Lady"
	fmt.Println("My dog's name is", dogsName)
	fmt.Println(z)
	f()
	f()
	f()
	const q string = "frozen"
	fmt.Println(q)
	var (
		name = "James"
		age = 26
		height = "6'"
	)
	fmt.Println("Name:", name, "- Age:", age, "- Height:", height)
	const (
		IPhone = "iPhone"
		IPad = "iPad"
		Android = "Android"
	)
	fmt.Println(IPhone, IPad, Android)
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}

func f() {
	fmt.Println(z)
}
