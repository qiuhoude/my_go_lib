package leetcode

import "testing"

// 62. 不同路径 https://leetcode-cn.com/problems/unique-paths/

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径？

输入：m = 3, n = 7
输出：28

输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右
3. 向下 -> 向右 -> 向下

输入：m = 7, n = 3
输出：28

输入：m = 3, n = 3
输出：6

1 <= m, n <= 100
题目数据保证答案小于等于 2 * 109

思路:
这个题和 64 最小路径和 题思路有点类似
[i][j]路径 = [i][j-1]路径 + [i-1][j]路径, [0][i]和[i][0]得路径都是1

*/
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	// 第一排
	for c := 1; c < n; c++ {
		dp[0][c] += dp[0][c-1]
	}
	// 第一列
	for r := 1; r < m; r++ {
		dp[r][0] += dp[r-1][0]
	}
	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			dp[r][c] = dp[r-1][c] + dp[r][c-1]
		}
	}
	return dp[m-1][n-1]
}

func Test_uniquePaths(t *testing.T) {
	tests := []struct {
		arg1     int
		arg2     int
		expected int
	}{
		{7, 3, 28},
		{3, 3, 6},
	}
	for _, tt := range tests {
		if got := uniquePaths(tt.arg1, tt.arg2); got != tt.expected {
			t.Errorf("uniquePaths(%v,%v) => got=%v  expected=%v", tt.arg1, tt.arg2, got, tt.expected)
		}
	}
}
