package binary_tree_inorder_traversal

import "github.com/yiwenlong/algorithm/common"

func inorderTraversal(root *common.TreeNode) []int {
	var res []int
	var stack []*common.TreeNode
	for 0 < len(stack) || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		index := len(stack) - 1
		res = append(res, stack[index].Val)
		root = stack[index].Right
		stack = stack[:index]
	}
	return res
}
