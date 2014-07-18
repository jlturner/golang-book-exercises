package main
import "fmt"

func makeOddsGenerator () (generator func() int) {
	odds := 1
	generator = func() (output int) {
		output = odds
		odds += 2
		return
	}
	return
}

func main () {
	oddsGenerator := makeOddsGenerator()
	for i := 0; i < 10; i++ {
		fmt.Println(oddsGenerator())
	}
	return
}
