package leetcode

// 145. 二叉树的后序遍历 https://leetcode-cn.com/problems/binary-tree-postorder-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	leftRes := postorderTraversal(root.Left)
	if leftRes != nil {
		res = append(res, leftRes...)
	}
	rightRes := postorderTraversal(root.Right)
	if rightRes != nil {
		res = append(res, rightRes...)
	}
	res = append(res, root.Val)
	return res
}
