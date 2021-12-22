package leetcode

import "testing"

// 129. 求根节点到叶节点数字之和 https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/
/*
给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：

例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。

叶节点 是指没有子节点的节点。

输入：root = [1,2,3]
输出：25
解释：
从根到叶子节点路径 1->2 代表数字 12
从根到叶子节点路径 1->3 代表数字 13
因此，数字总和 = 12 + 13 = 25

输入：root = [4,9,0,5,1]
输出：1026
解释：
从根到叶子节点路径 4->9->5 代表数字 495
从根到叶子节点路径 4->9->1 代表数字 491
从根到叶子节点路径 4->0 代表数字 40
因此，数字总和 = 495 + 491 + 40 = 1026

树中节点的数目在范围 [1, 1000] 内
0 <= Node.val <= 9
树的深度不超过 10

思路:
使用深度优先遍历,每次递归都会将当前值*10
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return sumNumbersAddDfs(root, 0)
}
func sumNumbersAddDfs(root *TreeNode, prevSum int) int {
	sum := prevSum*10 + root.Val
	if root.Left == nil && root.Right == nil { // 叶子节点
		return sum
	}
	res := 0
	if root.Left != nil {
		res += sumNumbersAddDfs(root.Left, sum)
	}
	if root.Right != nil {
		res += sumNumbersAddDfs(root.Right, sum)
	}

	return res
}

func sumNumbers1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	allPath := sumNumbersAdd(root, root.Val)
	sum := 0
	for _, v := range allPath {
		sum += v
	}
	return sum
}

func sumNumbersAdd(root *TreeNode, num int) []int {
	if root.Left == nil && root.Right == nil { // 叶子节点
		return []int{num}
	}
	var allPath []int
	if root.Left != nil {
		allPath = append(allPath, sumNumbersAdd(root.Left, num*10+root.Left.Val)...)
	}
	if root.Right != nil {
		allPath = append(allPath, sumNumbersAdd(root.Right, num*10+root.Right.Val)...)
	}
	return allPath
}

func Test_sumNumbers(t *testing.T) {
	root := BuildCompleteBinaryTree([]int{4, 9, 0, 5, 1})
	res := sumNumbers(root)
	want := 1026
	if res != want {
		t.Fatalf("want: %v, but got: %v", want, res)
	}
}
