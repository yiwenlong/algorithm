package merge_intervals

//56. 合并区间
//给出一个区间的集合，请合并所有重叠的区间。
//
//示例 1:
//
//输入: [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//示例 2:
//
//输入: [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	for i := 0; i < len(intervals)-1; i++ {
		min := i
		for j := i + 1; j < len(intervals); j++ {
			if intervals[min][0] > intervals[j][0] {
				min = j
			}
		}
		intervals[i], intervals[min] = intervals[min], intervals[i]
	}
	var ret [][]int
	l, r := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if r >= intervals[i][0] {
			if r < intervals[i][1] {
				r = intervals[i][1]
			}
		} else {
			ret = append(ret, []int{l, r})
			l, r = intervals[i][0], intervals[i][1]
		}
	}
	ret = append(ret, []int{l, r})
	return ret
}
