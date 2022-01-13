package leetcode

// 337. 打家劫舍 III https://leetcode-cn.com/problems/house-robber-iii/

/*
在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。
除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。

输入: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \
     3   1

输出: 7
解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.

输入: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \
 1   3   1

输出: 9
解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.

思路：
将树分成各个小树
给一个 node 节点分成两种情况，1.选择node 2.不选择node
f(node) 表示选择node, g(node)不选择node的值
1. 选择node:		f(node) = node.val + g(node.left) + g(node.right)
2. 不选择node: 	g(node) = max(f(node.left),g(node.left))  + max(f(node.right),g(node.right))
3. 最后比较 max(f(node),g(node))
不选择node = max(选择node.left,不选择node.left) + max(选择node.right,不选择node.right)

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func robiii(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxFn := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	/*
	f(node) 表示选择node, g(node)不选择node的值
	1. 选择node:		f(node) = node.val + g(node.left) + g(node.right)
	2. 不选择node: 	g(node) = max(f(node.left),g(node.left))  + max(f(node.right),g(node.right))
	3. 最后比较 max(f(node),g(node))
	不选择node = max(选择node.left,不选择node.left) + max(选择node.right,不选择node.right)
	*/
	// dfsFn 表示 r 节点 返回的是 (选择,不选择的)的值
	var dfsFn func(r *TreeNode) (selected, notSelected int)
	dfsFn = func(r *TreeNode) (selected, notSelected int) {
		if r == nil {
			return
		}
		ls, ln := dfsFn(r.Left)
		rs, rn := dfsFn(r.Right)
		selected = r.Val + ln + rn                  // 选择r
		notSelected = maxFn(ls, ln) + maxFn(rs, rn) // 不选择r
		return
	}
	return maxFn(dfsFn(root))
}
