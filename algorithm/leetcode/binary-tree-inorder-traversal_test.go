package leetcode

// 94. 二叉树的中序遍历 https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	leftRes := inorderTraversal(root.Left)
	if leftRes != nil {
		res = append(res, leftRes...)
	}
	res = append(res, root.Val)
	rightRes := inorderTraversal(root.Right)
	if rightRes != nil {
		res = append(res, rightRes...)
	}
	return res
}
