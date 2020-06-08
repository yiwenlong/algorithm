package main

import "fmt"

func QuickSort(arr []int, begin, end int) {
	if end <= begin {
		return
	}
	pivot := partition(arr, begin, end)
	QuickSort(arr, begin, pivot-1)
	QuickSort(arr, pivot+1, end)
}

func partition(arr []int, begin, end int) int {
	pivot, counter := end, begin
	for i := begin; i < end; i++ {
		if arr[i] < arr[pivot] {
			arr[counter], arr[i] = arr[i], arr[counter]
			counter++
		}
	}
	arr[pivot], arr[counter] = arr[counter], arr[pivot]
	fmt.Printf("pivot: %d, arr: %v\n", counter, arr)
	return counter
}

func main() {
	arr := []int{2, 3, 1, 5}
	QuickSort(arr, 0, 3)
	fmt.Printf("%v\n", arr)
}
