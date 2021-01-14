package kind_of_sorting

func BubbleSort(arr []int) []int {
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
	return arr
}

func InsertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		el := arr[i]
		j := i
		for j >= 1 && arr[j-1] > el {
			arr[j] = arr[j-1]
			j = j - 1
		}
		arr[j] = el

	}

	return arr
}
