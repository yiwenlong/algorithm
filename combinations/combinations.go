package main

//
//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
//示例:
//输入: n = 4, k = 2
//输出:
//[
//[2,4],
//[3,4],
//[2,3],
//[1,2],
//[1,3],
//[1,4],
//]
//
func combine(n int, k int) [][]int {
	if k == 0 {
		return [][]int{}
	}
	path := make([]int, 0)
	ans := make([][]int, 0)
	dfs(1, n, k, &path, &ans)
	return ans
}

func dfs(level, n, k int, path *[]int, ans *[][]int) {
	if len(*path) == k {
		tmp := make([]int, k)
		copy(tmp, *path)
		*ans = append(*ans, tmp)
		return
	}
	if level == n+1 {
		return
	}
	if len(*path)+(n-level)+1 < k {
		return
	}
	dfs(level+1, n, k, path, ans)
	*path = append(*path, level)
	dfs(level+1, n, k, path, ans)
	*path = (*path)[:len(*path)-1]
}

func main() {
	combine(5, 3)
}
