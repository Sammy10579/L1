package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две
//числовых переменных a,b, значение которых > 2^20.

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("9743089257348957023457", 22)
	b.SetString("6412304632074609723457", 22)

	div := new(big.Int)
	div.Div(a, b)
	fmt.Printf("Результат - %v\n", div)

	mult := new(big.Int)
	mult.Mul(a, b)
	fmt.Printf("Результат - %v\n", mult)

	sum := new(big.Int)
	sum.Add(a, b)
	fmt.Printf("Результат - %v\n", sum)

	dif := new(big.Int)
	dif.Sub(a, b)
	fmt.Printf("Результат - %v\n", dif)
}
