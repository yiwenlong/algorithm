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
	if i < 0 || i > bondI-1 || j < 0 || j > bondJ-1 {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	// current logic
	grid[i][j] = 0
	// drill down
	setZero(grid, i-1, j, bondI, bondJ)
	setZero(grid, i+1, j, bondI, bondJ)
	setZero(grid, i, j-1, bondI, bondJ)
	setZero(grid, i, j+1, bondI, bondJ)
}

func numIslands2(grid [][]byte) int {
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

type UnionFind struct {
	count  int
	parent []int
}

func (u *UnionFind) Make(n int) {
	u.count = n
	u.parent = make([]int, n)
	for i := 0; i < n; i++ {
		u.parent[i] = i
	}
}

func (u *UnionFind) Find(x int) int {
	for u.parent[x] != x {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u *UnionFind) Union(p, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	if rootP == rootQ {
		return
	}
	u.parent[rootP] = rootQ
	u.count--
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	uf := UnionFind{}
	rows := len(grid)
	cols := len(grid[0])
	uf.Make(rows*cols + 1)
	nigs := [][]int{{1, 0}, {0, 1}}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			index := i*cols + j
			if grid[i][j] == 1 {
				for _, nig := range nigs {
					r := i + nig[0]
					c := j + nig[1]
					if r >= 0 && r < rows && c >= 0 && c < cols && grid[r][c] == 1 {
						uf.Union(index, r*cols+c)
					}
				}
			} else {
				uf.Union(index, rows*cols)
			}
		}
	}
	return uf.count - 1
}

func main() {
	grid := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	fmt.Printf("numIslands: %d\n", numIslands(grid))
	//setZero(grid, 0, 0, len(grid), len(grid[0]))
	//for _, line := range grid{
	//	fmt.Printf("%v\n", line)
	//}
	//fmt.Printf("\n")
	//grid = [][]byte {
	//	{1,1,0,0,0},
	//	{1,1,0,1,0},
	//	{0,0,1,0,0},
	//	{0,0,0,1,1},
	//}
	//setZero(grid, 0, 0, len(grid), len(grid[0]))
	//for _, line := range grid{
	//	fmt.Printf("%v\n", line)
	//}
	//fmt.Printf("\n")
	//grid = [][]byte {
	//	{1,1,1},
	//	{0,1,0},
	//	{1,1,1},
	//}
	//setZero(grid, 0, 0, len(grid), len(grid[0]))
	//for _, line := range grid{
	//	fmt.Printf("%v\n", line)
	//}
}
