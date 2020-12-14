package main

import (
	"fmt"
)

func main() {
	var fibonachiNumber uint32
	fmt.Scan(&fibonachiNumber)
	mapFibonachi := map[uint32]uint32{
		0: 0,
		1: 1,
		2: 1,
		3: 2,
		4: 3,
		5: 5,
		6: 8,
	}
	_, exist := mapFibonachi[fibonachiNumber]
	if exist == true {
		fmt.Println(mapFibonachi[fibonachiNumber])
	} else {
		mapFibonachi[fibonachiNumber] = findFibonachiElement(fibonachiNumber)
		fmt.Println(mapFibonachi[fibonachiNumber])
	}
}

//рекурсивная функция findFibonachiElement принимает порядковый номер числа Фибоначчи и возвращает само число Фибоначчи
func findFibonachiElement(fnum uint32) uint32 {
	if fnum == 0 {
		return 0
	} else if fnum == 1 || fnum == 2 {
		return 1
	} else {
		return findFibonachiElement(fnum-2) + findFibonachiElement(fnum-1)
	}
}
