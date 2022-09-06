package main

import (
	"fmt"
	"sort"
)

func main() {
	m := []int{1234, 12312, 321, 432, 865, 23, 4321, 543, 24, 32, 57}
	sort.Ints(m)
	a := sort.SearchInts(m, 57)
	fmt.Println(m[a])
}
