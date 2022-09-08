package main

import "fmt"

//Разработать программу, которая переворачивает подаваемую на ход строку
//(например: «главрыба — абырвалг»). Символы могут быть unicode.

func reverse(sw string) string {
	sR := []rune(sw)
	for a, b := 0, len(sw)-1; a < b; a, b = a+1, b-1 {
		sR[a], sR[b] = sR[b], sR[a]
	}
	return string(sR)
}
func main() {
	sw := "glavryba"
	newString := reverse(sw)
	fmt.Println(newString)
}
