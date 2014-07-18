package main

import "fmt"
import m "golang-book/chapter11/math"

func main() {
	xs := []float64{1,2,3,4}
	avg := m.Average(xs)
	min, max := m.MinMax(xs)
	fmt.Println(avg, min, max)
}
