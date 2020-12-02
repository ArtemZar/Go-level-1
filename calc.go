package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var numberA, numberB, resalt float64
	var operator string
	fmt.Println("Введите два числа через пробел: ")
	fmt.Scanln(&numberA, &numberB)
	fmt.Println("Укажите оператор (+, -, /, *, ^): ")
	fmt.Scanln(&operator)
	switch operator {
	case "+":
		resalt = numberA + numberB
	case "-":
		resalt = numberA - numberB
	case "/":
		if numberB == 0 {
			fmt.Println("На 0 делить нельзя")
			os.Exit(1)
		} else {
			resalt = numberA / numberB
		}
	case "*":
		resalt = numberA * numberB
	case "^":
		resalt = math.Pow(numberA, numberB)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %f\n", resalt)

}
