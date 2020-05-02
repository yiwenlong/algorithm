package main

import "fmt"
//
//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
//此外，你可以假设该网格的四条边均被水包围。
//
//示例 1:
//输入:
//11110
//11010
//11000
//00000
//输出: 1
//
//示例 2:
//输入:
//11000
//11000
//00100
//00011
//输出: 3
//解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。

// https://leetcode-cn.com/problems/number-of-islands/

func setZero(grid [][]byte, i, j, bondI, bondJ int) {
	// terminal
	if i < 0 || i > bondI - 1 || j < 0 || j > bondJ - 1 { return }
	if grid[i][j] == 0 { return }
	// current logic
	grid[i][j] = 0
	// drill down
	setZero(grid, i - 1, j, bondI, bondJ)
	setZero(grid, i + 1, j, bondI, bondJ)
	setZero(grid, i, j - 1, bondI, bondJ)
	setZero(grid, i, j + 1, bondI, bondJ)
}

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				res += 1
				setZero(grid, i, j, len(grid), len(grid[i]))
			}
		}
	}
	return res
}

func main() {
	grid := [][]byte {
		{1,1,1,1,0},
		{1,1,0,1,0},
		{1,1,0,0,0},
		{0,0,0,0,0},
	}
	setZero(grid, 0, 0, len(grid), len(grid[0]))
	for _, line := range grid{
		fmt.Printf("%v\n", line)
	}
	fmt.Printf("\n")
	grid = [][]byte {
		{1,1,0,0,0},
		{1,1,0,1,0},
		{0,0,1,0,0},
		{0,0,0,1,1},
	}
	setZero(grid, 0, 0, len(grid), len(grid[0]))
	for _, line := range grid{
		fmt.Printf("%v\n", line)
	}
	fmt.Printf("\n")
	grid = [][]byte {
		{1,1,1},
		{0,1,0},
		{1,1,1},
	}
	setZero(grid, 0, 0, len(grid), len(grid[0]))
	for _, line := range grid{
		fmt.Printf("%v\n", line)
	}
}
