package main
import "fmt"

func fahrenheitToCelsius(fahrenheit float64) (float64) {
	return (fahrenheit - 32) * (5.0 / 9.0)
}

func main() {
	fmt.Println("Enter fahrenheit:")
	var inputFahrenheit float64
	fmt.Scanf("%f", &inputFahrenheit)
	fmt.Println(fahrenheitToCelsius(inputFahrenheit))
}
