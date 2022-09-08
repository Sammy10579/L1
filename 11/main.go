package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

func main() {
	set1 := []int{1, 3, 12, 3, 5, 34, 77}
	set2 := []int{3, 6, 12, 9, 3, 34, 2, 4, 77}
	resultMap := make(map[int]bool)
	result := []int{}
	for _, val1 := range set1 {
		for _, val2 := range set2 {
			if val1 == val2 {
				resultMap[val1] = true
			}
		}
	}
	for key, value := range resultMap {
		if value {
			result = append(result, key)
		}
	}
	fmt.Println(result)
}
