package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const helpMessage = "command format: 'cmd num1 num2'"

func sum(a, b float64) (float64, error) {
	return a + b, nil
}

func subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("На 0 делить нельзя")
	}
	return a / b, nil
}

func elevate(a, b float64) (float64, error) {
	return math.Pow(a, b), nil
}

var cmdMap = map[string]func(float64, float64) (float64, error){
	"sum":      sum,
	"subtract": subtract,
	"multiply": multiply,
	"divide":   divide,
	"elevate":  elevate,
}

func processCmd(cmd string, a, b float64) (float64, error) {
	if cmdFunc, ok := cmdMap[cmd]; ok {
		return cmdFunc(a, b)
	}
	return 0, fmt.Errorf("команда не реализована")
}

func main() {
	fmt.Printf("Введите параметры команды калькулятора в формате 'cmd num1 num2' (варианты cmd приведены ниже)\nsum\nsubtract\nmultiply\ndivide\nelevate\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		tokens := []string{}

		for _, token := range strings.Split(input, " ") {
			token = strings.TrimSpace(token)
			if token != "" {
				tokens = append(tokens, token)
			}
		}
		if len(tokens) != 3 {
			fmt.Println(helpMessage)
			continue
		}
		cmd := tokens[0]
		num1, err1 := strconv.ParseFloat(tokens[1], 64)
		num2, err2 := strconv.ParseFloat(tokens[2], 64)
		if err1 != nil || err2 != nil {
			fmt.Println(helpMessage)
			continue
		}
		res, err := processCmd(cmd, num1, num2)
		if err != nil {
			fmt.Printf("%s: %v\n", input, err)
			continue
		}
		fmt.Println("Ответ: ", res)

	}
}
