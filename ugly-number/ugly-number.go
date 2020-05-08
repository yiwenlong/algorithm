package main


//我们把只包含因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
//
//
//
//示例:
//
//输入: n = 10
//输出: 12
//解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
//说明:
//
//1 是丑数。
//n 不超过1690。
//注意：本题与主站 264 题相同：https://leetcode-cn.com/problems/ugly-number-ii/

import "fmt"

func min(nums ... int) int {
	minNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < minNum {
			minNum = nums[i]
		}
	}
	return minNum
}

func nthUglyNumber(n int) int {
	var res []int
	res = append(res, 1)
	p2, p3, p5 := 0, 0, 0
	m := 0
	for i := 1; i < n; i++ {
		res = append(res, min(res[p2] * 2, res[p3] * 3, res[p5] * 5))
		if res[p2] * 2 == res[i - m] { p2++ }
		if res[p3] * 3 == res[i - m] { p3++ }
		if res[p5] * 5 == res[i - m] { p5++ }
		r := min(p2, p3, p5)
		res = res[min(p2, p3, p5):]
		p2 -= r
		p3 -= r
		p5 -= r
		m += r
	}
	return res[len(res) - 1]
}

func main() {
	fmt.Printf("nthUglyNumber(%d) = %d\n",10, nthUglyNumber(10))
}