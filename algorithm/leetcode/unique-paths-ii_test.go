package leetcode

// 63. 不同路径 II https://leetcode-cn.com/problems/unique-paths-ii/

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
网格中的障碍物和空位置分别用 1 和 0 来表示。

输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
输出：2
解释：
3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右

输入：obstacleGrid = [[0,1],[0,0]]
输出：1

m == obstacleGrid.length
n == obstacleGrid[i].length
1 <= m, n <= 100
obstacleGrid[i][j] 为 0 或 1

思路: 63 题基本一致, 遇到障碍物体设置成 0

[i][j]路径 = [i][j-1]路径 + [i-1][j]路径, [0][i]和[i][0] 如果自身时障碍物则跳过
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 { //起点终点都是障碍物就永远没有路径
		return 0
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	// 第一排
	for c := 1; c < n; c++ {
		if obstacleGrid[0][c] == 0 {
			dp[0][c] += dp[0][c-1]
		}
	}
	// 第一列
	for r := 1; r < m; r++ {
		if obstacleGrid[r][0] == 0 {
			dp[r][0] += dp[r-1][0]
		}
	}
	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			if obstacleGrid[r][c] == 0 {
				dp[r][c] = dp[r-1][c] + dp[r][c-1]
			}
		}
	}
	return dp[m-1][n-1]
}
