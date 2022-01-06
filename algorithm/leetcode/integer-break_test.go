package leetcode

import "testing"

// 343. 整数拆分 https://leetcode-cn.com/problems/integer-break/

/*

给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
说明: 你可以假设 n 不小于 2 且不大于 58。

思路:
自顶向下,将n数拆解成 (n-1)的最大乘积*1 , (n-2)的最大乘积*2, (n-3)的最大乘积*3... (n-(n-1))最大成绩*(n-1)
求这些的最大值
自底向上,使用动态规划 max(multiMax, maxFn(dp[i-j]*j, (i-j)*j))

*/

func integerBreak(n int) int {
	if n < 2 {
		return 0
	}
	dp := make([]int, n+1) // 下标表示当前值i,value表示到i的最大乘积
	dp[1] = 1
	maxFn := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 2; i <= n; i++ {
		multiMax := 0
		for j := 1; j < i; j++ {
			multiMax = maxFn(multiMax, maxFn(dp[i-j]*j, (i-j)*j))
		}
		dp[i] = multiMax
	}
	return dp[n]
}

func Test_integerBreak(t *testing.T) {
	for n := 2; n <= 10; n++ {
		res := integerBreak(n)
		t.Logf("%v %v\n", n, res)
	}
}
