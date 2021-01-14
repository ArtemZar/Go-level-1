package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/ArtemZar/Go-level-1/lesson3-sortslise/kind_of_sorting"
)



func createRandomSlice() []int {
	count := 5
	sliceForSort := make([]int, 0, count)
	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(1000)
		sliceForSort = append(sliceForSort, n)
	}
	return sliceForSort
}

func main() {
	sliceForBubbleSort := createRandomSlice()
	fmt.Println("Исходный слайс: ", sliceForBubbleSort)
	fmt.Println("Алгоритм сортировки пузырьком: ", kind_of_sorting.BubbleSort(sliceForBubbleSort))

	sliceForInsertSort := createRandomSlice()
	fmt.Println("Исходный слайс: ", sliceForInsertSort)
	fmt.Println("Алгоритм сортировки вставками: ", kind_of_sorting.InsertSort(sliceForInsertSort))
}
