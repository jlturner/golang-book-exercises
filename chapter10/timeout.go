package main
import(
	"fmt"
	"time"
)

func main() {
	go func() {
		timed_out := false
		for ; timed_out == false ; {
			select {
			case <- time.After(time.Second * 2):
				fmt.Println("timeout!")
				timed_out = true
//			default:
//				fmt.Println("waiting...")
			}
		}
		fmt.Println("done")
	}()

	var input string
	fmt.Scanln(&input)
}
