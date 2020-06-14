package main

import "fmt"

//面试题40. 最小的k个数
//输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
//
//示例 1：
//输入：arr = [3,2,1], k = 2
//输出：[1,2] 或者 [2,1]
//
//示例 2：
//输入：arr = [0,1,2,1], k = 1
//输出：[0]
//
//限制：
//0 <= k <= arr.length <= 10000
//0 <= arr[i] <= 10000

// quick sort algorithm.
func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == k {
		return arr
	}
	lastNumbers(arr, k, 0, len(arr)-1)
	return arr[:k]
}

func lastNumbers(arr []int, k, begin, end int) {
	if k == 0 {
		return
	}
	pivot := partition(arr, begin, end)
	if k == pivot-begin || k == pivot-begin+1 {
		return
	} else if k > pivot-begin+1 {
		lastNumbers(arr, k-(pivot-begin+1), pivot+1, end)
	} else if k < pivot-begin {
		lastNumbers(arr, k, begin, pivot-1)
	}
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
	return counter
}

func main() {
	//fmt.Printf("%v\n", getLeastNumbers([]int{1, 2, 3, 4, 5}, 3))
	//fmt.Printf("%v\n", getLeastNumbers([]int{0,0,1,3,4,5,0,7,6,7}, 9))
	fmt.Printf("%v\n", getLeastNumbers([]int{0, 1, 2, 3, 4, 0, 3, 3, 8, 1, 4, 6, 2, 8, 8, 15, 10, 0, 9, 9, 1, 2, 17, 8, 17, 25, 18, 18, 16, 13, 18, 29, 2, 3, 32, 2, 26, 23, 18, 8, 34, 8, 11, 36, 36, 39, 46, 30, 21, 25, 21, 14, 41, 10, 31, 55, 45, 16, 33, 47, 4, 52, 59, 60, 1, 43, 42, 10, 12, 56, 12, 27, 22, 52, 38, 12, 41, 42, 71, 5, 42, 76, 8, 3, 31, 65, 11, 29, 28, 68, 33, 50, 73, 87, 22, 68, 31, 1, 38, 89}, 60))
}
