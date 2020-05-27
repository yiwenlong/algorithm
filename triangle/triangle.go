package main

//120. 三角形最小路径和
//给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
//
//相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
//
//例如，给定三角形：
//[
//[2],
//[3,4],
//[6,5,7],
//[4,1,8,3]
//]
//自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
//
//说明：
//如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。

func minVal(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < i; j++ {
			temp := triangle[i-1][j] + triangle[i][j]
			if j-1 >= 0 {
				temp = minVal(temp, triangle[i-1][j-1]+triangle[i][j])
			}
			triangle[i][j] = temp
		}
		triangle[i][i] += triangle[i-1][i-1]
	}
	temp := triangle[len(triangle)-1][0]
	for _, v := range triangle[len(triangle)-1] {
		if v < temp {
			temp = v
		}
	}
	return temp
}
