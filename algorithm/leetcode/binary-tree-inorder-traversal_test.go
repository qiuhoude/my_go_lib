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

func inorderTraversal_iteration(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	// 辅助的数据结构
	type command struct {
		Is   bool // 是否是要遍历的元素
		Node *TreeNode
	}

	stack := make([]command, 0)
	stack = append(stack, command{false, root})
	for len(stack) > 0 {
		cmd := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cmd.Is {
			res = append(res, cmd.Node.Val)
		} else {
			if cmd.Node.Right != nil {
				stack = append(stack, command{false, cmd.Node.Right})
			}
			stack = append(stack, command{true, cmd.Node})
			if cmd.Node.Left != nil {
				stack = append(stack, command{false, cmd.Node.Left})
			}
		}
	}
	return res
}
