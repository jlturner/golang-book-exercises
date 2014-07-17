package main
import "fmt"

func main () {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83

	total := 0.0
	for i := 0; i < 5; i++ {
		total += y[i]
	}
	fmt.Println(total / 5)

	y[4] = 120
	var newTotal float64 = 0
	for i := 0; i < len(y); i++ {
		newTotal += y[i]
	}
	fmt.Println(newTotal / float64(len(y)))

	y[1] = 999
	
	var anotherTotal float64 = 0
	for _, value := range y {
		anotherTotal += value
	}
	fmt.Println(anotherTotal / float64(len(y)))

	g := [5]float64 { 22, 99, 44, 88, 55 }
	j := [5]int {
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(g,j)

	t := j[1:len(j)-1]
	fmt.Println(t)
	fmt.Println(j[:])
	fmt.Println(append(j[2:4], 100, 200, 300, 400, 600, 900))

	slice1 := j[0:4]
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

	myMap := make(map[string]string)
	myMap["key"] = "10"
	myMap["name"] = "james"
	fmt.Println(myMap)
	delete(myMap, "key")
	fmt.Println(myMap)
}
