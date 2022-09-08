package main

import (
	"fmt"
	"sort"
)

//Реализовать быструю сортировку массива (quicksort) встроенными методами
//языка.

func main() {
	m := []int{25, 8756, 324, 3, 12, 432, 543, 7}
	sort.Ints(m)
	fmt.Println(m)
}
