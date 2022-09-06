package main

import (
	"fmt"
	"sort"
)

func main() {
	m := []int{25, 8756, 324, 3, 12, 432, 543, 7}
	sort.Ints(m)
	fmt.Println(m)
}
