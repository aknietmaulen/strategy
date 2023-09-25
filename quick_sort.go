package main

type QuickSortStrategy struct {
}

func (quicksort QuickSortStrategy) Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, element := range arr[1:] {
		if element <= pivot {
			left = append(left, element)
		} else {
			right = append(right, element)
		}
	}

	left = quicksort.Sort(left)
	right = quicksort.Sort(right)

	return append(append(left, pivot), right...)
}
