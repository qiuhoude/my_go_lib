package leetcode

// 450. 删除二叉搜索树中的节点 https://leetcode-cn.com/problems/delete-node-in-a-bst/

/*
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。

输入：root = [5,3,6,2,4,null,7], key = 3
输出：[5,4,6,2,null,null,7]
解释：给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。
一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。
另一个正确答案是 [5,2,6,null,4,null,7]。


输入: root = [5,3,6,2,4,null,7], key = 0
输出: [5,3,6,2,4,null,7]
解释: 二叉树不包含值为 0 的节点

输入: root = [], key = 0
输出: []

节点数的范围 [0, 104].
-105 <= Node.val <= 105
节点值唯一
root 是合法的二叉搜索树
-105 <= key <= 105

进阶： 要求算法时间复杂度为 O(h)，h 为树的高度。

思路:
待删除以下分几种情况
1. 待删除节点左子树为空的情况
2. 待删除节点右子树为空的情况
3. 左右都有数据
	1. 找到比待删除节点大的最小节点, 即待删除节点右子树的最小节点
	2. 用这个节点顶替待删除节点的位置
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归写法
func deleteNode450(root *TreeNode, key int) *TreeNode {
	return bstDeleteNode(root, key)
}

func bstDeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = bstDeleteNode(root.Left, key)
		return root
	} else if root.Val < key {
		root.Right = bstDeleteNode(root.Right, key)
		return root
	} else { // 找到要删除的节点
		/*
		待删除以下分几种情况
		1. 待删除节点左子树为空的情况 将右边返回
		2. 待删除节点右子树为空的情况 将左边返回
		3. 左右都有数据
			1. 找到比待删除节点大的最小节点, 即待删除节点右子树的最小节点
			2. 用这个节点顶替待删除节点的位置
				还可以将待删除节点左边直接接到该节点的左边
		*/
		if root.Left == nil {
			rt := root.Right
			root.Right = nil
			return rt
		}
		if root.Right == nil {
			rt := root.Left
			root.Left = nil
			return rt
		}
		cur := root.Right

		//  用这个节点顶替待删除节点的位置
		/*
			var prev *TreeNode // cur的父节点
			for ;  cur.Left != nil; cur = cur.Left {
				prev = cur
			}
			if prev != nil {
				prev.Left = cur.Right
				cur.Right = root.Right // 防止 root.Right == cur,自己指向自己
			}
			cur.Left = root.Left
			root.Left = nil
			root.Right = nil
			return cur
		*/
		// 将待删除节点左边直接接到该节点的左边
		for ; cur.Left != nil; cur = cur.Left {
		}
		cur.Left = root.Left
		rt := root.Right
		root.Left = nil
		root.Right = nil
		return rt
	}
}

//func Test_deleteNode(t *testing.T) {
//	//  [5,3,6,2,4,null,7]
//	tree := BuildCompleteBinaryTreeInteger([]*Integer{i2Integer(5), i2Integer(3), i2Integer(6), i2Integer(2), i2Integer(4), nil, i2Integer(7)})
//	afterTree := deleteNode450(tree, 3)
//	_ = afterTree
//}
