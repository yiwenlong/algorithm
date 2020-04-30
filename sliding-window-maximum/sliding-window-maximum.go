package sliding_window_maximum

import "fmt"

//给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//返回滑动窗口中的最大值。
//进阶：
//你能在线性时间复杂度内解决此题吗？
//
//示例:
//
//输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3,3,5,5,6,7]
//解释:
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
//1 [3  -1  -3] 5  3  6  7       3
//1  3 [-1  -3  5] 3  6  7       5
//1  3  -1 [-3  5  3] 6  7       5
//1  3  -1  -3 [5  3  6] 7       6
//1  3  -1  -3  5 [3  6  7]      7
// 
//提示：
//
//1 <= nums.length <= 10^5
//-10^4 <= nums[i] <= 10^4
//1 <= k <= nums.length
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sliding-window-maximum

func enqueue(queue *[]int, k, num int)  {
	// 如果队列为空直接添加新元素
	// 如果新元素小于队首元素，并且队列还有位置，则直接添加新元素
	if len(*queue) == 0 || (len(*queue) < k && (*queue)[0] > num){
		*queue = append(*queue, num)
		return
	}
	// 如果队首位置的元素小于新元素，直接清空队列，然后添加新元素
	if (*queue)[0] <= num {
		*queue = []int{ num }
		return
	}
	// 如果队列满了，队首出队
	if len(*queue) >= k {
		*queue = (*queue)[len(*queue) - k + 1:]
	}
	*queue = append(*queue, num)
	// 调整队列，最大值左侧的所有元素出队
	maxIndex := 0
	for j := 0; j < len(*queue); j++ {
		if (*queue)[j] >= (*queue)[maxIndex] {
			maxIndex = j
		}
	}
	*queue = (*queue)[maxIndex:]
}
// 思路：
// 使用一个最大长度为 K 的队列维护窗口中的值。
// 始终保持队首为窗口中的最大值。
// 窗口滑动时，使队列最大值左侧的数全部出队。
// 每滑动一次窗口，取队首的值添加到结果集合。
func maxSlidingWindow(nums []int, k int) []int {
	var queue, res []int
	for i := 0; i < len(nums); i++ {
		enqueue(&queue, k, nums[i])
		fmt.Printf("queue: %v\n", queue)
		if i >= (k - 1) {
			res = append(res, queue[0])
		}
	}
	return res
}