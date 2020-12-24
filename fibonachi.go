package main

import (
	"fmt"
	"os"
)

var mapFibonachi = map[uint32]uint32{
	0: 0,
	1: 1,
	2: 1,
}

func main() {
	var fibonachiNumber uint32
	for true {
		fmt.Println("Введите номер числа Фибоначчи")
		_, err := fmt.Scan(&fibonachiNumber) // проверяем ошибку соответсвея типу вводимых данных
		if err != nil {
			fmt.Println("Проверьте типы входных параметров")
			os.Exit(1)
		}

		_, exist := mapFibonachi[fibonachiNumber]
		if exist == true {
			fmt.Println(mapFibonachi[fibonachiNumber])
		} else {

			fmt.Println(findFibonachiElement(fibonachiNumber))
			//fmt.Println(mapFibonachi)
		}
	}
}

//рекурсивная функция findFibonachiElement принимает порядковый номер числа Фибоначчи и возвращает само число Фибоначчи
func findFibonachiElement(fnum uint32) uint32 {
	if fnum == 0 {
		return 0
	} else if fnum == 1 || fnum == 2 {
		return 1
	} else {
		mapFibonachi[fnum] = findFibonachiElement(fnum-2) + findFibonachiElement(fnum-1)
		return mapFibonachi[fnum]

	}
}
