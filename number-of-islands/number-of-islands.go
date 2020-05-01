package main

import "fmt"

func setZero(grid [][]byte, startI, startJ, bondI, bondJ int) {
	grid[startI][startJ] = 0
	next := startI - 1
	if next >= 0 && grid[next][startJ] == 1 {
		setZero(grid, next, startJ, bondI, bondJ)
	}
	next = startI + 1
	if next < bondI && grid[next][startJ] == 1 {
		setZero(grid, next, startJ, bondI, bondJ)
	}
	next = startJ - 1
	if next >= 0 && grid[startI][startJ] == 1 {
		setZero(grid, startI, next, bondI, bondJ)
	}
	next = startJ + 1
	if next < bondJ && grid[startI][next] == 1 {
		setZero(grid, startI, next, bondI, bondJ)
	}
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

//[


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
}
