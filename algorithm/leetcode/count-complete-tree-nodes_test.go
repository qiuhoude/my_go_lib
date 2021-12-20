package leetcode

import (
	"testing"
)

// 222. 完全二叉树的节点个数 https://leetcode-cn.com/problems/count-complete-tree-nodes/
/*
给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，
并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

输入：root = [1,2,3,4,5,6]
输出：6

输入：root = []
输出：0

输入：root = [1]
输出：1


树中节点的数目范围是[0, 5 * 104]
0 <= Node.val <= 5 * 104
题目数据保证输入的树是 完全二叉树

进阶：遍历树来统计节点是一种时间复杂度为 O(n) 的简单解决方案。你可以设计一个更快的算法吗？

思路1:
总数量 = 左子树数量+左子树数量 + 1

思路2:
利用完全二叉树的定义, 先计算二叉树的高度 maxLevel, 那么最后一排节点下标编号的区间段 [2^(maxLevel-1),(2^maxLevel)),
后面可以使用二分查找在该区间找出最后一个节点的位置,利用该节点是否存在进行查找
如何判断序号节点是否存在?
使用二进制表示树节点的下标,
例如:下标 9 二进制 1001, 首位时1,0表示左边1表示右边,路径:右右左就找到下标9的节点
下标12, 二进制 1100 路径:左右右
取出下标的每一位形成路径,跟着路径找看是否能找到节点,能找到说明存在反之不存在

*/

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var maxLevel uint = 0
	for e := root; e != nil; e = e.Left {
		maxLevel++
	}

	existsFn := func(node *TreeNode, level uint, index int) bool {
		bits := 1 << (level - 2)
		cur := node
		for ; cur != nil && bits > 0; bits = bits >> 1 {
			if index&bits == 0 {
				cur = cur.Left
			} else {
				cur = cur.Right
			}
		}
		return cur != nil
	}

	// [ 2^(maxLevel-1),2^(maxLevel) )之间进行二分查找 2^maxLevel<= x <2^(maxLevel)
	low, high := 1<<(maxLevel-1), 1<<maxLevel
	for low < high {
		mid := low + ((high - low) >> 1)
		if existsFn(root, maxLevel, mid) { // 存在
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low - 1
}

func Test_countNodes2(t *testing.T) {
	//root := BuilderCompleteBinaryTree([]int{1, 2, 3, 4, 5, 6})
	root := BuilderCompleteBinaryTree([]int{1})
	res := countNodes2(root)
	t.Log(res)
}

func BuilderCompleteBinaryTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	return createNode(arr, 0)
}

func createNode(arr []int, index int) *TreeNode {
	var node *TreeNode
	if index < len(arr) {
		node = &TreeNode{
			Val:   arr[index],
			Left:  createNode(arr, index*2+1),
			Right: createNode(arr, index*2+2),
		}
	}
	return node

}
