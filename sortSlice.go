package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(arr []int) ([]int, error) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}
	return arr, nil
}

func insertSort(arr []int) ([]int, error) {
	for i := 1; i < len(arr); i++ {
		el := arr[i]
		j := i
		for j >= 1 && arr[j-1] > el {
			arr[j] = arr[j-1]
			j = j - 1
		}
		arr[j] = el

	}

	return arr, nil
}

func createRandomSlice() ([]int, error) {
	count := 5
	sliceForSort := make([]int, 0, count)
	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(1000)
		sliceForSort = append(sliceForSort, n)
	}
	return sliceForSort, nil
}

func main() {
	sliceForBubbleSort := createRandomSlice()
	fmt.Println("Исходный слайс: ", sliceForBubbleSort)
	fmt.Println("Алгоритм сортировки пузырьком: ", bubbleSort(sliceForBubbleSort))

	sliceForInsertSort := createRandomSlice()
	fmt.Println("Исходный слайс: ", sliceForInsertSort)
	fmt.Println("Алгоритм сортировки вставками: ", insertSort(sliceForInsertSort))
}
