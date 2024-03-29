package leetcode

//437. 路径总和 III https://leetcode-cn.com/problems/path-sum-iii/
/*
给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。


输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
输出：3
解释：和等于 8 的路径有 3 条，如图所示。
示例 2：

输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：3

二叉树的节点个数的范围是 [0,1000]
-109 <= Node.val <= 109
-1000 <= targetSum <= 1000

思路:
1. 深度优先,遍历每个节点为起点, 调用一个子过程找到路径和
2. 使用前缀和,把路径上之前的数加起来,

*/

// 前缀和的方式
func pathSum(root *TreeNode, targetSum int) int {
	prefix := make(map[int]int) // <节点的前缀和,拥有路满足条件径数量>
	prefix[0] = 1               // 防止只有一个根节点
	return pathSumDfs(root, prefix, targetSum, 0)
}

// curSum 当前节点的和
func pathSumDfs(root *TreeNode, prefix map[int]int, targetSum, curSum int) int {
	if root == nil {
		return 0
	}

	curSum += root.Val
	res := prefix[curSum-targetSum]
	prefix[curSum]++ //
	res += pathSumDfs(root.Left, prefix, targetSum, curSum)
	res += pathSumDfs(root.Right, prefix, targetSum, curSum)
	prefix[curSum]-- //
	return res
}

func pathSumiii(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	cnt := findNodeSum(root, sum)
	cnt += pathSumiii(root.Left, sum)
	cnt += pathSumiii(root.Right, sum)
	return cnt
}

func findNodeSum(node *TreeNode, num int) int {
	if node == nil {
		return 0
	}

	ret := 0
	if node.Val == num {
		ret += 1 // 找到了+1
	}
	// 还需要
	ret += findNodeSum(node.Left, num-node.Val)
	ret += findNodeSum(node.Right, num-node.Val)
	return ret
}
