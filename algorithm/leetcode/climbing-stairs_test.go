package leetcode

import "testing"

// 70. 爬楼梯 https://leetcode-cn.com/problems/climbing-stairs/

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

思路:
1. 使用自顶向下的思路, 要求n层阶梯的解就是 n-1层解题 加上 n-2层阶梯的解, 加上记忆化搜索的方式减少递归次数
2. 使用自顶向上的思路, 也就是动态规划的思路, 要求n层阶梯就得从 1层阶梯解; 2层阶梯得解; 3阶梯得解= 1层阶梯解+2层阶梯解,一步步从小到大地推到n
*/

func climbStairs2(n int) int {
	dp := make([]int, n+1) // 记录每层阶梯需要得步数
	for i := 1; i <= n; i++ {
		if i == 1 {
			dp[i] = 1
			continue
		}
		if i == 2 {
			dp[i] = 2
			continue
		}
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairs(n int) int {
	tables := make([]int, n+1) // 下标表示阶梯,数值表示步数,记忆化搜索
	var dfsFn func(int) int
	dfsFn = func(m int) int {
		if tables[m] != 0 {
			return tables[m]
		}
		if m == 1 {
			return 1
		}
		if m == 2 {
			return 2
		}
		res := dfsFn(m-1) + dfsFn(m-2)
		tables[m] = res
		return res
	}
	return dfsFn(n)
}

func Test_climbStairs(t *testing.T) {

	tests := []struct {
		arg      int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
	}

	for _, tt := range tests {
		if got := climbStairs2(tt.arg); got != tt.expected {
			t.Errorf("climbStairs(%v) got %v, expected %v", tt.arg, got, tt.expected)
		}
	}

}
