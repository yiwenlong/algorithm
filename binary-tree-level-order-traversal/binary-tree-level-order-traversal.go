package binary_tree_level_order_traversal

import (
	"github.com/yiwenlong/algorithm/common"
)

func levelOrder(root *common.TreeNode) [][]int {
	var queue []*common.TreeNode
	var res [][]int
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		levelSize := len(queue)
		var subRes []int
		for i := 0; i < levelSize; i++ {
			subRes = append(subRes, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, subRes)
		queue = queue[levelSize:]
	}
	return res
}
