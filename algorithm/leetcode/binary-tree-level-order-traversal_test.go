package leetcode

// 102. 二叉树的层序遍历 https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

/*
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

思路:
使用 queue 的数据结构,首先将根节点添加, 将队列中的数据全部取出来,然后把他们的子节点全部添加到队列中,每次就是一层
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	que := []*TreeNode{root}
	for len(que) > 0 {
		var list []int
		var tmpQue []*TreeNode
		for _, v := range que {
			list = append(list, v.Val)
			if v.Left != nil {
				tmpQue = append(tmpQue, v.Left)
			}
			if v.Right != nil {
				tmpQue = append(tmpQue, v.Right)
			}
		}
		res = append(res, list)
		que = tmpQue
	}
	return res
}
