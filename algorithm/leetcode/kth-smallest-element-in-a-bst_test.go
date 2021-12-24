package leetcode

// 230. 二叉搜索树中第K小的元素 https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/

/*
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

输入：root = [3,1,4,null,2], k = 1
输出：1

输入：root = [5,3,6,2,4,null,null,1], k = 3
输出：3

树中的节点数为 n 。
1 <= k <= n <= 104
0 <= Node.val <= 104

思路:
利用二分搜索树中序遍历是有序(ascending order),对遍历的数字进行计数等于k进行返回
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	//中序遍历是有序的,对遍历的数字进行计数等于k进行返回
	stack := []*TreeNode{}
	cur := root
	cnt := 0 // 计数
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		// pop
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cnt++
		if k == cnt {
			return cur.Val
		}
		cur = cur.Right
	}
	return -1 // 未找到
}
