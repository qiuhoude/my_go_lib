package leetcode

import "math"

// 110. 平衡二叉树 https://leetcode-cn.com/problems/balanced-binary-tree/
/*
给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一棵高度平衡二叉树定义为：
一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。


输入：root = [3,9,20,null,null,15,7]
输出：true


输入：root = [1,2,2,3,3,null,null,4,4]
输出：false

输入：root = []
输出：true


树中的节点数在范围 [0, 5000] 内
-104 <= Node.val <= 104

思路:  左孩子数量,和右孩子数量如差值

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !(isBalanced(root.Left) && isBalanced(root.Right)) {
		return false
	}
	leftHeight := treeHeight(root.Left)
	rightHeight := treeHeight(root.Right)
	if int(math.Abs(float64(leftHeight-rightHeight))) > 1 {
		return false
	}
	return true
}

func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 可用理解为是后序遍历
	// 在跟父点处比较大小
	leftDepth := treeHeight(root.Left)
	rightDepth := treeHeight(root.Right)
	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
}
