package main

import (
	"fmt"
	"sort"
)

//Реализовать бинарный поиск встроенными методами языка.

func main() {
	m := []int{1234, 12312, 321, 432, 865, 23, 4321, 543, 24, 32, 57}
	sort.Ints(m)
	a := sort.SearchInts(m, 57)
	fmt.Println(m[a])
}

/*
func binaryS(elem int, arr []int) bool {
	middle := len(arr) / 2
	if elem == arr[middle] {
		return true
	} else if len(arr) == 1 {
		return false
	} else if elem > arr[middle] {
		arr = arr[middle:]
	} else if elem < arr[middle] {
		arr = arr[:middle]
	}
	return binaryS(elem, arr)
}

func main() {
	arr := make([]int, 100)
	for i := range arr {
		arr[i] = i
	}
	fmt.Println(binaryS(32, arr))
	fmt.Println(binaryS(100, arr))
}
*/
