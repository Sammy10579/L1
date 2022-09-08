package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func reverse(sw string) string {
	words := strings.Split(sw, " ")
	var rW []string
	for i := len(words) - 1; i >= 0; i-- {
		rW = append(rW, words[i])
	}
	var result string
	for _, word := range rW {
		result += word + " "
	}
	return result
}
func main() {
	sw := "snow dog sun"
	newString := reverse(sw)
	fmt.Println(newString)
}
