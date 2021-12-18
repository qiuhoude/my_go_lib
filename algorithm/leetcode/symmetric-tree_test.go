package leetcode

//101. 对称二叉树 https://leetcode-cn.com/problems/symmetric-tree/

/*
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3

进阶：
你可以运用递归和迭代两种方法解决这个问题吗？

思路:
1.递归思路,使用两指针,用同样的方式遍历左边和右边,判断左边是否等于右边
2.遍历思路

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	//层级遍历, 同时遍历边 和右边两个树
	que := []*TreeNode{root.Left, root.Right}
	for len(que) > 1 {
		t1 := que[0]
		t2 := que[1]
		que = que[2:]
		if t1 == nil && t2 == nil {
			continue
		} else if t1 != nil && t2 != nil {
			if t1.Val != t2.Val {
				return false
			}
			que = append(que, t1.Left)
			que = append(que, t2.Right)
			que = append(que, t1.Right)
			que = append(que, t2.Left)
		} else {
			return false
		}
	}
	return true
}
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compareMirrorTree(root.Left, root.Right)
}

func compareMirrorTree(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	} else if a != nil && b != nil {
		return a.Val == b.Val && compareMirrorTree(a.Left, b.Right) && compareMirrorTree(a.Right, b.Left)
	} else { // 一个等于nil一个不等于nil
		return false
	}
}
