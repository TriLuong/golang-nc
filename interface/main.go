package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	Width  float64
	Height float64
}

func (r Rect) Area() float64 {
	return 300
}

func (r Rect) Perimeter() float64 {
	return 400
}

type Circle struct {
	Diameter float64
}

func (r Circle) Area() float64 {
	return 100
}

func (r Circle) Perimeter2() float64 {
	return 600
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

func PrintArea2(i interface{}) {
	s, ok := i.(Shape)
	if ok {
		fmt.Println(s.Area())
	} else {
		fmt.Println("Failed")

	}

}

func main() {

	aRect := Rect{Height: 100, Width: 200}

	PrintArea(aRect)

	aCircle := Circle{Diameter: 50}
	PrintArea2(aCircle)

}
