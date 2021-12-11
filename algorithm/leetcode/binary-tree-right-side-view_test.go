package leetcode

// 199. 二叉树的右视图 https://leetcode-cn.com/problems/binary-tree-right-side-view/

/*
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。


输入: [1,2,3,null,5,null,4]
输出: [1,3,4]


输入: [1,null,3]
输出: [1,3]

输入: []
输出: []

二叉树的节点个数的范围是 [0,100]
-100 <= Node.val <= 100

思路:
和 102题一样都是层级遍历, 只是每次获取每层的最后一个元素

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	que := []*TreeNode{root}
	for len(que) > 0 {
		var tmpQue []*TreeNode
		for _, v := range que {

			if v.Left != nil {
				tmpQue = append(tmpQue, v.Left)
			}
			if v.Right != nil {
				tmpQue = append(tmpQue, v.Right)
			}
		}
		res = append(res, que[len(que)-1].Val)
		que = tmpQue
	}
	return res
}
