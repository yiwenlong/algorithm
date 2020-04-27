package _sum

import "sort"

// https://leetcode-cn.com/problems/3sum/

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
// 注意：答案中不可以包含重复的三元组
//
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
// 满足要求的三元组集合为：
// [
//  [-1, 0, 1],
//  [-1, -1, 2]
// ]
//
// 使用双指针法
// 前提是要对数组进行排序
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	size := len(nums)
	var res [][]int
	var left int
	var right int
	if size < 3 { return res }
	for index := 0; index < size; index++ {
		// nums[index] > 0 意味着 index 右边的所有数字都 > 0，所以不可能出现相加等于0 的情况
		if nums[index] > 0 { break }
		// 为了去掉重复的情况
		if index > 0 && nums[index] == nums[index-1] { continue }
		left = index + 1
		right = size - 1
		for left < right {
			// 此时出现的三个数字：nums[index] nums[left] nums[right] 会有以下三种关系
			// 1、nums[right] + nums[left] = -nums[index]
			// 	此时算找到一组结果，保存结果，left向右移动，right向左，继续寻找下一个
			if nums[right] + nums[left] == -nums[index] {
				res = append(res, []int{nums[index], nums[left], nums[right]})
				// 除去可能产生的重复情况
				for left < right && nums[left] == nums[left + 1] { left++ } // 去重
				for left < right && nums[right] == nums[right - 1] { right-- } // 去重
				left++
				right--
				continue
			}
			// 2、nums[right] + nums[left] > -nums[index]
			// 	此时说明需要更小的数，所以让右边的指针向左移动
			if nums[right] + nums[left] > -nums[index] {
				right--
				continue
			}
			// 3、nums[right] + nums[left] < -nums[index]
			// 	此时说明需要更大的数，所以让左边的数向右移动
			if nums[right] + nums[left] < - nums[index] {
				left++
				continue
			}
		}
	}
	return res
}
