package main

import "fmt"

//Удалить i-ый элемент из слайса.

func main() {
	a := []string{"A", "B", "C", "D", "E"}
	i := 2

	copy(a[i:], a[i+1:])

	a[len(a)-1] = ""

	a = a[:len(a)-1]

	fmt.Println(a)
}
