package permutations

func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	used := make(map[int]bool, 0)
	dfs(path, nums, used, &res)
	return res
}

func dfs(path []int, nums []int, used map[int]bool, res *[][]int) {
	if len(nums) == len(path) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	for _, v := range nums {
		if used[v] {
			continue
		}
		used[v] = true
		path = append(path, v)
		dfs(path, nums, used, res)
		used[v] = false
		path = path[:len(path)-1]
	}
}
