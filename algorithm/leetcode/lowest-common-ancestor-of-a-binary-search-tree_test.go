package leetcode

// 235. 二叉搜索树的最近公共祖先 https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

/*
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，
满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]

输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6
解释: 节点 2 和节点 8 的最近公共祖先是 6。

输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。


所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉搜索树中。


思路:
因为是二分搜索树, 可以每次比较p,q两个节点在cur节点一边,如果在一边可以继续往下找,如在两边说明就是最近的祖先
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	cur := root
	for {
		if cur.Val > p.Val && cur.Val > q.Val { // 全在左边
			cur = cur.Left
		} else if cur.Val < p.Val && cur.Val < q.Val { // 全在右边
			cur = cur.Right
		} else {
			break
		}
	}
	return cur
}
