package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в
//канал, а с другой стороны канала — читать. По истечению N секунд программа
//должна завершаться.

func readCannel(ch chan int) {
	for data := range ch {
		fmt.Printf("data printing %d\n", data)
	}
}

func writeChannel(ch chan int) {
	for {
		data := rand.Intn(100)
		ch <- data
	}
}

func main() {
	var seconds = 10

	ch := make(chan int)
	defer close(ch)

	go writeChannel(ch)
	go readCannel(ch)

	time.Sleep(time.Duration(seconds) * time.Second)
}
