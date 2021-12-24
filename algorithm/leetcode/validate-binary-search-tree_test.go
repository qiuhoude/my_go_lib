package leetcode

import (
	"math"
	"testing"
)

// 98. 验证二叉搜索树 https://leetcode-cn.com/problems/validate-binary-search-tree/

/*
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

输入：root = [2,1,3]
输出：true

输入：root = [5,1,4,null,null,3,6]
输出：false
解释：根节点的值是 5 ，但是右子节点的值是 4 。

树中节点数目范围在[1, 104] 内
-231 <= Node.val <= 231 - 1

思路:
1. 递归的思路, 当前节点是否有效,子节点是否有效,全部才真的有效,每次判断有效都是个子过程,注意需要判断节点的上下界
2. 利用二分搜索树中序遍历后是有序的性质,可以检测遍历后的数组是否有序
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 中序遍历写法
func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	cur := root
	var prevNode *TreeNode
	for cur != nil || len(stack) > 0 {
		for ; cur != nil; cur = cur.Left {
			stack = append(stack, cur)
		}
		// pop
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if prevNode != nil && cur.Val <= prevNode.Val {
			return false
		}
		prevNode = cur
		cur = cur.Right
	}
	return true
}

// 递归写法
func isValidBST1(root *TreeNode) bool {
	return checkIsValidBSTNode(root, math.MinInt64, math.MaxInt64)
}
func checkIsValidBSTNode(root *TreeNode, lower, upper int64) bool {
	if root == nil {
		return true
	}
	if !(int64(root.Val) > lower && int64(root.Val) < upper) {
		return false
	}
	return checkIsValidBSTNode(root.Left, lower, int64(root.Val)) && checkIsValidBSTNode(root.Right, int64(root.Val), upper)
}

func Test_isValidBST(t *testing.T) {
	tree := BuildCompleteBinaryTree([]int{2, 2, 2})
	want := false
	got := isValidBST(tree)
	if got != want {
		t.Fatalf("want:%v got:%v", want, got)
	}
}
