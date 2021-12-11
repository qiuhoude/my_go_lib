package leetcode

// 144. 二叉树的前序遍历 https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

/*

给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

输入：root = [1,null,2,3]
输出：[1,2,3]

输入：root = []
输出：[]

输入：root = [1]
输出：[1]

输入：root = [1,2]
输出：[1,2]

输入：root = [1,null,2]
输出：[1,2]

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100

进阶：递归算法很简单，你可以通过迭代算法完成吗？

*/

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */
func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	leftRes := preorderTraversal(root.Left)
	if leftRes != nil {
		res = append(res, leftRes...)
	}
	rightRes := preorderTraversal(root.Right)
	if rightRes != nil {
		res = append(res, rightRes...)
	}
	return res
}

func preorderTraversal_iteration(root *TreeNode) []int {
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
			if cmd.Node.Left != nil {
				stack = append(stack, command{false, cmd.Node.Left})
			}
			stack = append(stack, command{true, cmd.Node})
		}
	}
	return res
}
