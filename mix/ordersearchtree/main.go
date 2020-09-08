package main

/*
Given a binary search tree, rearrange the tree in in-order so that the leftmost node in the tree is now the root of the tree,
and every node has no left child and only 1 right child.
*/

// TreeNode holds node information
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret = TreeNode{Val: 0}
var tmp = &ret

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	increasingBST(root.Left)
	tmp.Right = &TreeNode{
		Val: root.Val,
	}
	tmp = tmp.Right
	increasingBST(root.Right)
	return ret.Right
}