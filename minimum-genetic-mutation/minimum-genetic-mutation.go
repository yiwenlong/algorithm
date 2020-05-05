package minimum_genetic_mutation

//一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。
//假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。
//例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。
//与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。
//现在给定3个参数 — start, end, bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变化次数。如果无法实现目标变化，请返回 -1。
//注意:
//起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
//所有的目标基因序列必须是合法的。
//假定起始基因序列与目标基因序列是不一样的。
//
//示例 1:
//start: "AACCGGTT"
//end:   "AACCGGTA"
//bank: ["AACCGGTA"]
//返回值: 1
//
//示例 2:
//start: "AACCGGTT"
//end:   "AAACGGTA"
//bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
//返回值: 2
//
//示例 3:
//start: "AAAAACCC"
//end:   "AACCCCCC"
//bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
//
//返回值: 3
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-genetic-mutation
//
// 解题思路
// 1. 构造一个队列用于存放搜索的结果。
// 	  初始值为 start
// 	  找到队列中所有元素的一次基因变化的元素，然后入队。如果找到了目标元素，直接返回结果。
//    如果没有找到，当前搜索后的结果出队，基于新结果重新进行搜索。
// 2. 使用一个数组作为缓存，记录已经遍历过的结果。
func minMutation(start string, end string, bank []string) int {
	used := make([]int, len(bank))
	queue := []string{ start }
	nums := 0
	for len(queue) != 0 {
		nums++
		level := len(queue)
		for i := 0; i < level; i++ {
			for k, v := range bank {
				if v != start && used[k] == 0 {
					diff := 0
					for m := 0; m < 8; m++ {
						if v[m] != queue[i][m] { diff++ }
						if diff >= 2 { break }
					}
					if diff == 1 {
						used[k] = nums
						if bank[k] == end { return nums }
						queue = append(queue, bank[k])
					}
				}
			}
		}
		queue = queue[level:]
	}
	return -1
}