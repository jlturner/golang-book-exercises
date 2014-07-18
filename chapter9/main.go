package main
import "fmt"
import "math"

type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, shape := range m.shapes {
		area += shape.area()
	}
	return area
}
func (m *MultiShape) perimeter() float64 {
	var perimeter float64
	for _, shape := range m.shapes {
		perimeter += shape.perimeter()
	}
	return perimeter
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() (area float64) {
	area = math.Pi * c.r * c.r
	return
}
func (c *Circle) perimeter() (perimeter float64) {
	perimeter = 2.0 * math.Pi * c.r
	return
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() (area float64) {
	area = (r.x2 - r.x1) * (r.y2 - r.y1)
	return
}
func (r *Rectangle) perimeter() (perimeter float64) {
	perimeter = (r.x2 - r.x1) * 2 +  (r.y2 - r.y1) * 2
	return
}

func main() {
	c := Circle{
		x: 0,
		y: 0,
		r: 5,
	}
	fmt.Println(c.x, c.y, c.r, c.area(), c.perimeter())
	r := Rectangle{
		x1: 10,
		y1: 20,
		x2: 20,
		y2: 30,
	}
	fmt.Println(r.area(), r.perimeter())
	m := MultiShape{[]Shape{&c, &r}}
	fmt.Println(m.perimeter())
}
