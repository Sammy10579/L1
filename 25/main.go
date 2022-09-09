package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

func sleep1(s time.Duration) {
	<-time.After(s)
}

func sleep2(s time.Duration) {
	<-time.NewTimer(s).C
}

func main() {
	sleep1(time.Second * 3)
	fmt.Println("it's sleep1")
	sleep2(time.Second * 3)
	fmt.Println("it's sleep2")

}
