package main

import "fmt"

func main() {
	a, b := 1, 2
	a, b = b, a
	fmt.Printf("a - %d,  b - %d\n", a, b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a - %d , b - %d \n", a, b)
}
