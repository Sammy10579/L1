package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке
//уникальные (true — если уникальные, false etc). Функция проверки должна быть
//регистронезависимой.
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func main() {
	fmt.Println(CountLetters("abcde"))
	fmt.Println(CountLetters("abcdee"))
	fmt.Println(CountLetters("abCDe"))
	fmt.Println(CountLetters("AbcDEe"))
}

func CountLetters(str string) (string, bool) {
	s := strings.ToLower(str)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return s, false
			}
		}
	}
	return s, true
}
