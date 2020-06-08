package main

import "fmt"

//
//在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
//示例:
//输入:
//1 0 1 0 0
//1 0 1 1 1
//1 1 1 1 1
//1 0 0 1 0
//输出: 4
//

func square(i, j, level int, matrix [][]byte, res *int) {
	// terminal
	if i+level >= len(matrix) || j+level >= len(matrix[i]) {
		return
	}
	for c := j; c <= j+level; c++ {
		if matrix[i+level][c] == 0 {
			return
		}
	}
	for r := i; r <= i+level; r++ {
		if matrix[r][j+level] == 0 {
			return
		}
	}
	// current logic
	*res = (level + 1) * (level + 1)
	// dill down
	square(i, j, level+1, matrix, res)
}

func maximalSquare(matrix [][]byte) int {
	max := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 1 {
				var tmp int
				square(i, j, 0, matrix, &tmp)
				if tmp > max {
					max = tmp
				}
			}
		}
	}
	return max
}

func main() {
	//matrix := [][]byte {
	//	{1, 0, 1, 0, 0},
	//	{1, 0, 1, 1, 1},
	//	{1, 1, 1, 1, 1},
	//	{1, 0, 0, 1, 0},
	//}
	matrix := [][]byte{
		{0, 0, 0, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
		{0, 1, 1, 1}}

	res := maximalSquare(matrix)
	for i := 0; i < len(matrix); i++ {
		fmt.Printf("%v\n", matrix[i])
	}
	fmt.Printf("res: %d\n", res)
}
