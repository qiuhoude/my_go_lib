package leetcode

import "math"

// 64. 最小路径和 https://leetcode-cn.com/problems/minimum-path-sum/

/*
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。

输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。

输入：grid = [[1,2,3],[4,5,6]]
输出：12

m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 100

思路:
基本思路和 120 基本一致, 使用二维数组记录到 [i][j] 位置最小路径和的值, 一步步递推
到 [i][j] 这个一步必然是 [i-1][j] 或 [i][j-1] 这两个坐标过来的,
sumMin([i][j]) = min([i-1][j],[i][j-1]) + [i][j]
*/

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[i]))
	}
	// 初始化
	dp[0][0] = grid[0][0] // 第一步
	// 第一排
	for c := 1; c < len(dp[0]); c++ {
		dp[0][c] = dp[0][c-1] + grid[0][c]
	}
	// 第一列
	for r := 1; r < len(dp); r++ {
		dp[r][0] = dp[r-1][0] + grid[r][0]
	}

	for r := 1; r < len(grid); r++ {
		for c := 1; c < len(grid[r]); c++ {
			dp[r][c] = int(math.Min(float64(dp[r-1][c]), float64(dp[r][c-1]))) + grid[r][c]
		}
	}
	n := len(dp)
	m := len(dp[n-1])
	return dp[n-1][m-1]
}
