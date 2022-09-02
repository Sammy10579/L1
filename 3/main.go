package main

import (
	"fmt"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
//квадратов(22+32+42….) с использованием конкурентных вычислений.

func sumSq(s []int) (sum int) {
	mych := make(chan int)

	sq := func(q int, mych chan<- int) {
		mych <- q * q
	}

	for _, v := range s {
		go sq(v, mych)
	}

	for i := 1; i <= len(s); i++ {
		sum += <-mych
	}
	return
}

func main() {
	s := []int{2, 4, 6, 8, 10}
	fmt.Println(sumSq(s))
}
