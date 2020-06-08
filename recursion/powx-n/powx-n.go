package powx_n
//
// 实现 pow(x, n) ，即计算 x 的 n 次幂函数。
//
// 示例 1:
//
// 输入: 2.00000, 10
// 输出: 1024.00000
// 示例 2:
//
// 输入: 2.10000, 3
// 输出: 9.26100
// 示例 3:
//
// 输入: 2.00000, -2
// 输出: 0.25000
// 解释: 2-2 = 1/22 = 1/4 = 0.25
// 说明:
// -100.0 < x < 100.0
// n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/powx-n
//
// 思路：
// 使用分治法，把本需要 N 次循环拆成 LOG(N)
func pow(x float64, n int) float64 {
	if n == 0 { return 1.0 }
	if n == 1 { return x }
	subRes := pow(x, n / 2)
	if n % 2 == 1 {
		return subRes * subRes * x
	} else {
		return subRes * subRes
	}
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / pow(x, -n)
	}
	return pow(x, n)
}

