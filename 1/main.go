package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type human struct {
	name     string
	lastname string
}

type action struct {
	human
}

func (h *human) tospeak() {
	fmt.Println("Hi! My name is " + h.name + ". " + "My lastname is " + h.lastname + ".")
}

func main() {
	h := human{name: "Artur", lastname: "Pirojkov"}
	a := action{h}
	a.tospeak()
}
