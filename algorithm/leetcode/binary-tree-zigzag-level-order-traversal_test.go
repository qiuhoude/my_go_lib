package leetcode

// 103. 二叉树的锯齿形层序遍历 https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

/*
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层序遍历如下：

[
  [3],
  [20,9],
  [15,7]
]

思路:
与102号问题一致, 只需要在遍历时判断层数,进行颠倒添加

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	que := []*TreeNode{root}
	isReverse := true
	for len(que) > 0 {
		var list []int
		var tmpQue []*TreeNode
		for i := 0; i < len(que); i++ {
			if isReverse {
				list = append(list, que[len(que)-1-i].Val)
			} else {
				list = append(list, que[i].Val)
			}
			v := que[i]
			if v.Left != nil {
				tmpQue = append(tmpQue, v.Left)
			}
			if v.Right != nil {
				tmpQue = append(tmpQue, v.Right)
			}
		}
		res = append(res, list)
		que = tmpQue
		isReverse = !isReverse
	}
	return res
}
