package main
import "fmt"

func main () {
	defer func () {
		str := recover()
		fmt.Println("Error:", str)
	}()
	panic("PANIC")
}
