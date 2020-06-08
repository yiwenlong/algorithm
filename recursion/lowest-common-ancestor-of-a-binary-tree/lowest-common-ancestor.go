package lowest_common_ancestor_of_a_binary_tree

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：后序遍历二叉树. 左 → 右 → 根
// 1. 用返回数字的结果代表结果，1代表已经找到了p , 2 代表已经找到了 q, 3 代表所有结果都找到了。
// 2. 结束条件。如果已经找到了，则不再继续遍历。直接返回给父节点。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, r := bsf(root, p, q)
	return r
}

func bsf(node, p, q *TreeNode) (int, *TreeNode) {
	if node == nil {
		return 0, nil
	}
	rl, rn := bsf(node.Left, p, q)
	if rl == 3 {
		// Left 是最近公共祖先
		return rl, rn
	}
	rr, rn := bsf(node.Right, p, q)
	if rr == 3 {
		// Right 是最近公共祖先
		return rr, rn
	}
	cr := 0
	if node.Val == p.Val {
		cr = 1 + rl + rr
	} else if node.Val == q.Val {
		cr = 2 + rl + rr
	} else {
		cr = rl + rr
	}
	if cr == 3 {
		return cr, node
	} else {
		return cr, nil
	}
}
