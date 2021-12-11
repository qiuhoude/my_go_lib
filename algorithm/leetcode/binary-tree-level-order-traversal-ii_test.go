package leetcode

// 107. 二叉树的层序遍历 II https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/

/*
给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层序遍历为：

[
  [15,7],
  [9,20],
  [3]
]

思路:
与102题一样,只不过将顺序都过来而已
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
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
	//reverse res
	h, t := 0, len(res)-1
	for h < t {
		res[h], res[t] = res[t], res[h]
		h++
		t--
	}
	return res
}
