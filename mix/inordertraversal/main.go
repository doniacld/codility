package inordertraversal

/*
Given a binary tree, return the inorder traversal of its nodes' values.
Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
*/

// TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorderTraversal
func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	st := make([]*TreeNode, 0)

	for len(st) > 0 || root != nil {
		if root != nil {
			st = append(st, root)
			root = root.Left
		} else {
			last := st[len(st)-1]
			st = st[:len(st)-1]
			ret = append(ret, last.Val)
			root = last.Right
		}
	}
	return ret
}
