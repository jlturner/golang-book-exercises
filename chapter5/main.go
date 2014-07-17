package main
import "fmt"

func main() {
	i := 1
	for i <= 99999 {
		fmt.Println(i)
		i = i + 1
	}
	for g := 1; g <= 10; g++ {
		if g % 2 == 0 {
			fmt.Println(g, "even")
		} else {
			fmt.Println(g, "odd")
		}
		switch g {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		case 5: fmt.Println("Five")
		case 6: fmt.Println("Six")
		default: fmt.Println("I don't know")
		}
	}
	
}
