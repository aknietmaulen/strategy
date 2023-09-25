package main

import "fmt"

func main() {
	arr := []int{7, 6, 5, 4}

	bubble_sort := BubbleSortStrategy{}
	quick_sort := QuickSortStrategy{}

	context := NewSortingContext(bubble_sort)
	sortedArr := context.Sort(arr)
	fmt.Println("Sorted by Bubble Sort: ", sortedArr)

	context.SetStrategy(quick_sort)
	sortedArr = context.Sort(arr)
	fmt.Println("Sorted by Quick Sort: ", sortedArr)
}
