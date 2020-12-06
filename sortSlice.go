package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(arr []int) []int {
	swappend := true
	for swappend {
		swappend = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swappend = true
			}
		}
	}
	return arr
}

func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i
		for j >= 1 && arr[j-1] > x {
			arr[j] = arr[j-1]
			j = j - 1
		}
		arr[j] = x

	}

	return arr
}

func createSlice() []int {
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
	sliceForBubbleSort := createSlice()
	fmt.Println("Исходный слайс: ", sliceForBubbleSort)
	fmt.Println("Алгоритм сортировки пузырьком: ", bubbleSort(sliceForBubbleSort))

	sliceForInsertSort := createSlice()
	fmt.Println("Исходный слайс: ", sliceForInsertSort)
	fmt.Println("Алгоритм сортировки вставками: ", insertSort(sliceForInsertSort))
}
