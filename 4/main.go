package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток). Реализовать
//набор из N воркеров, которые читают произвольные данные из канала и
//выводят в stdout. Необходима возможность выбора количества воркеров при
//старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
//способ завершения работы всех воркеров.

func worker(id int, c chan int) {
	for data := range c {
		time.Sleep(time.Millisecond)
		fmt.Printf("[worker %v] Sending result by worker %v\n", id, data)
	}
}

func main() {
	c := make(chan int)
	defer close(c)

	// N - количество воркеров
	N := 10

	for i := 0; i < N; i++ {
		go worker(i, c)
	}

	for {
		data := rand.Intn(100)
		c <- data
	}
}
