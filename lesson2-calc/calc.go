package main

import (
	"bufio"
	"fmt"
	"github.com/ArtemZar/Go-level-1/lesson2-calc/operations"
	"os"
	"strconv"
	"strings"
)

const helpMessage = "command format: 'cmd num1 num2'"



var cmdMap = map[string]func(float64, float64) (float64, error){
	"sum":      operations.Sum,
	"subtract": operations.Subtract,
	"multiply": operations.Multiply,
	"divide":   operations.Divide,
	"elevate":  operations.Elevate,
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
