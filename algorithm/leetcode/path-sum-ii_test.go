package leetcode

// 113. 路径总和 II https://leetcode-cn.com/problems/path-sum-ii/

/*
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
叶子节点 是指没有子节点的节点。


输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]

输入：root = [1,2,3], targetSum = 5
输出：[]

输入：root = [1,2], targetSum = 0
输出：[]

树中节点总数在范围 [0, 5000] 内
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000

思路:
使用递归回溯的方式,每次在经过的路径都减去当前节点的值,
看到叶子节点时结果是否是0
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSumii(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil { // 叶子节点
		if targetSum-root.Val == 0 { // 递归出口
			return [][]int{{root.Val}}
		} else {
			return nil
		}
	}
	var res [][]int
	if root.Left != nil {
		path := pathSumii(root.Left, targetSum-root.Val)
		if path != nil {
			for _, v := range path {
				res = append(res, append([]int{root.Val}, v...))
			}
		}
	}
	if root != nil {
		path := pathSumii(root.Right, targetSum-root.Val)
		if path != nil {
			for _, v := range path {
				res = append(res, append([]int{root.Val}, v...))
			}
		}
	}
	return res
}
