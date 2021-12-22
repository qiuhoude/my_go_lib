package leetcode

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// 257. 二叉树的所有路径 https://leetcode-cn.com/problems/binary-tree-paths/
/*
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

叶子节点 是指没有子节点的节点。

输入：root = [1,2,3,null,5]
输出：["1->2->5","1->3"]

输入：root = [1]
输出：["1"]

思路:
递归深度优先进行遍历, 和 126题的 dfs遍历出所有路径类似
1. 利用递归返回值进行
2. 利用 queue 在遇到叶子节点时进行大于全部
*/

func Test_binaryTreePaths(t *testing.T) {
	lr := TreeNode{Val: 5}
	l := TreeNode{Val: 2, Right: &lr}
	r := TreeNode{Val: 3}
	root := TreeNode{Val: 1, Left: &l, Right: &r}
	fmt.Println(binaryTreePaths(&root))
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var res []string
	que := []*TreeNode{}
	dfsBinaryTreePaths(root, que, &res)
	return res
}

func dfsBinaryTreePaths(root *TreeNode, que []*TreeNode, res *[]string) {
	if root.Left == nil && root.Right == nil { // 叶子节点,添加所有路径
		var sb strings.Builder
		if len(que) > 0 {
			for _, n := range que {
				sb.WriteString(strconv.Itoa(n.Val) + "->")
			}
		}
		sb.WriteString(strconv.Itoa(root.Val))
		*res = append(*res, sb.String())
		return
	}
	if root.Left != nil {
		dfsBinaryTreePaths(root.Left, append(que, root), res)
	}
	if root.Right != nil {
		dfsBinaryTreePaths(root.Right, append(que, root), res)
	}
}

//利用递归返回值进行
func binaryTreePaths1(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil { // 叶子节点
		return []string{strconv.Itoa(root.Val)}
	} else {
		var ret []string
		if root.Left != nil {
			c := strconv.Itoa(root.Val) + "->"
			sub := binaryTreePaths(root.Left)
			for _, v := range sub {
				ret = append(ret, c+v)
			}
		}
		if root.Right != nil {
			c := strconv.Itoa(root.Val) + "->"
			sub := binaryTreePaths(root.Right)
			for _, v := range sub {
				ret = append(ret, c+v)
			}
		}
		return ret
	}
}
