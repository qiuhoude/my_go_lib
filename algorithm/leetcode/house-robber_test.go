package leetcode

import "testing"

// 198. 打家劫舍 https://leetcode-cn.com/problems/house-robber/

/*
 309
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。

输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。

1 <= nums.length <= 100
0 <= nums[i] <= 400

思路:
动态规划
第n房子最大值 = max((n-2)最大值+n的数字 ,(n-1)最大值)
*/

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)

	maxFn := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		if i == 1 {
			dp[i] = maxFn(dp[0], nums[i-1])
		} else {
			dp[i] = maxFn(dp[i-1], dp[i-2]+nums[i-1])
		}
	}
	return dp[n]
}

func Test_rob(t *testing.T) {
	tests := []struct {
		arg []int

		expected int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{}, 0},
	}
	for _, tt := range tests {
		if got := rob(tt.arg); got != tt.expected {
			t.Errorf("rob(%v) => got=%v  expected=%v", tt.arg, got, tt.expected)
		}
	}
}
