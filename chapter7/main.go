package main
import "fmt"

func sumAndAverage (ints []int) (sum int, average float64) {
	sum = 0
	for _, v := range ints {
		sum += v
	}
	average = float64(sum) / float64(len(ints))
	return
}

func sumInts (ints ...int) (sum int) {
	sum = 0
	for _, v := range ints {
		sum += v
	}
	return
}

func averageFloatArray (xs []float64) (average float64) {
	total := float64(0)
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x <= 0 {
		return 1
	} else {
		return x * factorial(x - 1)
	}
}

func main () {
	a := []float64{ 1, 2, 3, 4, 5, 6 }
	defer fmt.Println(averageFloatArray(a))
	defer fmt.Println(sumInts(1, 2, 3, 4, 5, 6, 7))
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1,1))
	nextEven := makeEvenGenerator()
	for i := 0; i < 9; i++ {
		fmt.Println(nextEven())
	}
	fmt.Println(factorial(5))

	b := []int{ 1, 2, 3, 4, 5, 6 }
	sum, average := sumAndAverage(b)
	fmt.Println(sum, average)
}
