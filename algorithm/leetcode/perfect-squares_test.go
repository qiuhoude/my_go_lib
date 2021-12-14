package leetcode

import "math"

// 279. 完全平方数 https://leetcode-cn.com/problems/perfect-squares/

/*
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。


输入：n = 12
输出：3
解释：12 = 4 + 4 + 4

输入：n = 13
输出：2
解释：13 = 4 + 9

1 <= n <= 104

思路1:
对问题进行建模,转换成图论问题
从n到0,每个数字表示一个顶点了如果两个数字x到y相差一个完全平方数,则连接一条边,得到一个无权图.
问题转换成求整个无权图中n到0的最短路径.
使用广度优先遍历寻最短路径

思路2:
动态规划, f(i)=1+min(i-j*j), j=1..i, dp[i]中存储的是步数.
一步步求出 0~n每个数字的最小步数

*/

func numSquaresDP(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minStep := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			a := i - j*j // i - j*j 表示上一步
			if minStep > dp[a] {
				minStep = dp[a]
			}
		}
		dp[i] = minStep + 1
	}
	return dp[n]
}

func numSquares(n int) int {
	// 求解n点到0点最短路径, 每相差一个完全平方数就是一个顶点
	que := make([][2]int, 0)      // [2]int , [0]表示节点, [1]表示步数
	visited := make(map[int]bool) // 访问过的顶点
	visited[n] = true
	que = append(que, [2]int{n, 0})

	// 广度优先
	for len(que) > 0 {
		vertexN := que[0][0]
		step := que[0][1]
		// pop que
		que = que[1:]
		// 当前节点到下一个可到达节点所有情况
		for i := 1; true; i++ {
			nextVertexN := vertexN - i*i // 当前顶点 - 完全平方数
			if nextVertexN < 0 {
				break
			} else if nextVertexN == 0 { // 到达0顶点
				return step + 1
			}
			if !visited[nextVertexN] {
				que = append(que, [2]int{nextVertexN, step + 1})
				visited[nextVertexN] = true
			}
		}
	}
	panic("No solution")
}
