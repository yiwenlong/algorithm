package n_ary_tree_level_order_traversal

import "github.com/yiwenlong/algorithm/common"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *common.Node) [][]int {
	var (
		res   [][]int
		queue []*common.Node
	)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		level := len(queue)
		subres := make([]int, level)
		for i := 0; i < level; i++ {
			subres = append(subres, queue[i].Val)
			if len(queue[i].Children) > 0 {
				queue = append(queue, queue[i].Children...)
			}
		}
		res = append(res, subres)
		queue = queue[level:]
	}
	return res
}
