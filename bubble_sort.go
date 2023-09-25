package main

type BubbleSortStrategy struct {
}

func (bubblesort BubbleSortStrategy) Sort(arr []int) []int {
	l := len(arr)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < l; i++ {
			if arr[i-1] > arr[i] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
	}
	return arr
}
