package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками, которые
//представлены в виде структуры Point с инкапсулированными параметрами x,y и
//конструктором.

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) Point {
	return Point{x, y}
}

func Distance(point1 Point, point2 Point) float64 {
	dx := point1.x - point2.x
	dy := point1.y - point2.y
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func main() {
	point1 := NewPoint(10, 12)
	point2 := NewPoint(4, 5)
	fmt.Printf("distance between %v and %v is - %f", point1, point2, Distance(point1, point2))
}
