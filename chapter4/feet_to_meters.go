package main
import "fmt"

func feetToMeters (feet float64) (float64) {
	return feet * 0.3048
}

func main () {
	var inputFeet float64
	fmt.Println("Enter feet:")
	fmt.Scanf("%f", &inputFeet)
	fmt.Println(feetToMeters(inputFeet))
}
