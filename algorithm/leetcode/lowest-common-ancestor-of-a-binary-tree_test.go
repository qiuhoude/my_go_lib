package leetcode

// 236. 二叉树的最近公共祖先 https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，
最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”


输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。

输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出：5
解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。

输入：root = [1,2], p = 1, q = 2
输出：1

树中节点数目在范围 [2, 105] 内。
-109 <= Node.val <= 109
所有 Node.val 互不相同 。
p != q
p 和 q 均存在于给定的二叉树中。

思路:
235也是找公共祖先,235题是BST,所以很简单.
1. 使用两个数组分别记录查找p 和q的路径, 然后对比两个数组找公共节点
2.
*/
/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */
func lowestCommonAncestorii(root, p, q *TreeNode) *TreeNode {
	var res *TreeNode
	findLCA(root, p, q, &res)
	return res
}

// 最外层调用此函数 p,q必须在tree中
func findLCA(root, p, q *TreeNode, res **TreeNode) bool {
	if root == nil {
		return false
	}
	if root == p || root == q {
		*res = root
		return true
	}
	// 检测 p 或 q 其中一个是否在root.left中
	left := findLCA(root.Left, p, q, res)
	// 检测 p 或 q 其中一个是否在root.right中
	right := findLCA(root.Right, p, q, res)
	if left && right { // 两边都存在说明 p,q 分别 root两侧,而不是一侧
		*res = root
	}
	return left || right
}
