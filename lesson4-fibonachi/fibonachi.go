package main

import (
	"fmt"
)

var mapFibonachi = map[uint32]uint32{
	0: 0,
	1: 1,
	2: 1,
}

func main() {
	var fibonachiNumber uint32
	for true {
		fmt.Scan(&fibonachiNumber)

		_, exist := mapFibonachi[fibonachiNumber]
		if exist == true {
			fmt.Println(mapFibonachi[fibonachiNumber])
		} else {

			fmt.Println(FindFibonachiElement(fibonachiNumber))
			//fmt.Println(mapFibonachi)
		}
	}
}

//рекурсивная функция findFibonachiElement принимает порядковый номер числа Фибоначчи и возвращает само число Фибоначчи
func FindFibonachiElement(fnum uint32) uint32 {
	if fnum == 0 {
		return 0
	} else if fnum == 1 || fnum == 2 {
		return 1
	} else {
		mapFibonachi[fnum] = FindFibonachiElement(fnum-2) + FindFibonachiElement(fnum-1)
		return mapFibonachi[fnum]
	}
}
