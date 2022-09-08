package main

import (
	"fmt"
)

//Разработать программу, которая в рантайме способна определить тип
//переменной: int, string, bool, channel из переменной типа interface{}.

func t(data interface{}) {
	fmt.Printf("%v is %T\n", data, data)
}

func main() {
	t("str")
	t(342343)
	t(333221.33)
	t(make(chan int))
}

// Как вариант, можно создать структуру принимающую дата интерфейс, и в нем с помощью кострукции stich/case t.(type)
// реализовать каждый вариант.
/*func identifier(data interface{}) {
	switch data.(type) {
	case int:
		fmt.Println("it's int!")
	case string:
		fmt.Println("it's string!")
	case float64:
		fmt.Println("it's float64!")
	case rune:
		fmt.Println("it's rune!")
	case bool:
		fmt.Println("it's bool!")
	case chan int:
		fmt.Println("it's chan int!")
	}
}

func main() {
	var data []interface{}
	data = append(data, "str")
	data = append(data, 231233)
	data = append(data, 'd')
	data = append(data, true)
	data = append(data, 34123.4)
	data = append(data, make(chan int))

	for _, val := range data {
		identifier(val)
	}

}*/
