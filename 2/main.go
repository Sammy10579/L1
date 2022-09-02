package main

import (
	"os"
	"strconv"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел
//взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func squares(s []int) {
	var wg sync.WaitGroup

	wg.Add(len(s))

	for _, v := range s {
		go func(i int) {
			os.Stdout.WriteString(strconv.Itoa(i*i) + "\n")
			wg.Done()
		}(v)
	}
	wg.Wait()
}

func main() {
	squares([]int{2, 4, 6, 8, 10})
}
