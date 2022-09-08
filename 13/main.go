package main

import "fmt"

//Поменять местами два числа без создания временной переменной.

func main() {
	a, b := 1, 2
	fmt.Printf("a - %d,  b - %d\n", a, b)
	a, b = b, a
	fmt.Printf("a - %d,  b - %d\n", a, b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a - %d , b - %d \n", a, b)
}
