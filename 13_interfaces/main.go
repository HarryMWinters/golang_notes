package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

func getArea(s Shape) float64 {
	return s.area()
}

type Circle struct {
	x, y, radius float64
}

type Rectange struct {
	width, height float64
}

type Triangle struct {
	height, width float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectange) area() float64 {
	return r.height * r.width
}

func (t Triangle) area() float64 {
	return t.height * t.width * 0.5
}

func main() {
	myCircle := Circle{0, 0, 5}
	fmt.Println("myCircle = ", myCircle)
	fmt.Printf("Circle area = %f\n", getArea(myCircle))

	myRectangle := Rectange{2, 2}
	fmt.Println("myRectangle = ", myRectangle)
	fmt.Printf("Rectangle area = %f\n", getArea(myRectangle))

	fmt.Println("Program Complete")
}
