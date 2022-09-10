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

/*
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]

			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func main() {
	arr := make([]int, 15)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	fmt.Println(quicksort(arr))
}
*/
