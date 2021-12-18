package leetcode

import "math"

// 111. 二叉树的最小深度 https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

/*
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。

输入：root = [3,9,20,null,null,15,7]
输出：2

输入：root = [2,null,3,null,4,null,5,null,6]
输出：5


树中节点数的范围在 [0, 105] 内
-1000 <= Node.val <= 1000

思路1:
使用层级遍历,遇到节点左右都为null的直接返回,此处就是最小深度

思路2:
与104题有点类似,递归子问题解决,注意递归陷阱
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left == nil {
		return 1
	}
	minLeft := minDepth(root.Left)
	minRight := minDepth(root.Right)
	if root.Left == nil || root.Right == nil { // 只要有一边是空就该下面的子节点就是最小层数就是另一边的
		return minLeft + minRight + 1 // minLeft 和 minRight其中一个必然时0
	}
	return int(math.Min(float64(minLeft), float64(minRight))) + 1
}

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	queue := []*TreeNode{root}
out:
	for len(queue) > 0 {
		res++
		nodes := queue
		queue = nil
		for _, node := range nodes {
			if node.Left == nil && node.Right == nil { // 叶子节点
				break out
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}
