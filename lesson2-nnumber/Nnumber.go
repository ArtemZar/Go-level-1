package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	var N int64
	fmt.Println("Введите число от 2 до 1000: ")
	fmt.Scanln(&N)
	if N < 2 || N > 1000 {
		fmt.Println("не корректный ввод")
		os.Exit(1)
	}
	var i int64
	for i = 2; i < N; i++ {

		if big.NewInt(i).ProbablyPrime(0) {
			fmt.Println(i, "- простое число")
		}
	}
}
