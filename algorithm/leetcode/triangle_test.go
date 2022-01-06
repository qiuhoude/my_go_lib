package leetcode

import "math"

// 120. 三角形最小路径和 https://leetcode-cn.com/problems/triangle/

/*
给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。


输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：
   2
  3 4
 6 5 7
4 1 8 3
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

输入：triangle = [[-10]]
输出：-10

1 <= triangle.length <= 200
triangle[0].length == 1
triangle[i].length == triangle[i - 1].length + 1
-104 <= triangle[i][j] <= 104

进阶：
你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？

思路:
先从自顶向下进行思考,要求到 n 层最小路径和,就要求 n-1 层的每个位置,到 n 层路径和,在这些路径和中找最小值
具体做法,使用 n 的数组记录上一层到该层每个点的最小路径和

*/

func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle))
	dp[0] = triangle[0][0] // 初始化

	for i := 1; i < len(triangle); i++ {
		preOldVal := 0 // 前一个原来的值就是上一层 j-1位置的值,此处用于缓存,方便下一步计算
		lastIndex := len(triangle[i]) - 1
		for j := 0; j < lastIndex; j++ {
			if j == 0 {
				preOldVal = dp[0]
				// 到0这个位置,只能是上一步也是0这个位置
				dp[0] += triangle[i][j]
			} else {
				// 到 j 这个位置,只能是上一步 是 j 或j-1的位置才能到达,找到最小值就可以
				minVal := int(math.Min(float64(preOldVal), float64(dp[j])))
				preOldVal = dp[j]
				dp[j] = minVal + triangle[i][j]
			}
		}
		dp[lastIndex] = preOldVal + triangle[i][lastIndex] // 本排最后一步,一定是上一排最后一步才能走到的
	}
	// 在dp中找到最小值
	res := math.MaxInt32
	for _, v := range dp {
		if res > v {
			res = v
		}
	}
	return res
}
