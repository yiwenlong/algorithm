package main

import (
	"fmt"
	"github.com/yiwenlong/algorithm/common"
)

//根据一棵树的前序遍历与中序遍历构造二叉树。
//注意:
//你可以假设树中没有重复的元素。
//
//例如，给出
//前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
func buildTree(preorder []int, inorder []int) *common.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &common.TreeNode{Val: preorder[0]}
	}
	rootIndex := 0
	for ; rootIndex < len(preorder); rootIndex++ {
		if inorder[rootIndex] == preorder[0] {
			break
		}
	}
	return &common.TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:rootIndex], inorder[:rootIndex]),
		Right: buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:]),
	}
}

func main() {
	root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Printf("%d\n", root.Val)
	//fmt.Printf("%v\n", []int{3,9,20,15,7}[:1])
}
