package leetcode

import "testing"

// 518. 零钱兑换 II https://leetcode-cn.com/problems/coin-change-2/

/*
给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。

请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。

假设每一种面额的硬币有无限个。

题目数据保证结果符合 32 位带符号整数。

输入：amount = 5, coins = [1, 2, 5]
输出：4
解释：有四种方式可以凑成总金额：
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1

输入：amount = 3, coins = [2]
输出：0
解释：只用面额 2 的硬币不能凑成总金额 3 。

输入：amount = 10, coins = [10]
输出：1

1 <= coins.length <= 300
1 <= coins[i] <= 5000
coins 中的所有值 互不相同
0 <= amount <= 5000

思路:
转换成完全背包问题
此处是求组合问题，不是排列问题，所以外层循环是coins,
dp[i] = accumulate(0~amount, dp[amount-coin])

如果该题要求顺序 那么就和  377. 组合总和 一样了,需要外层循环是target

原因: 如果外层循环是target，dp[i]=+dp[i-num] 到某个金额会同一个coin放在不同位置进行计算，
会有重复的

*/

func change(amount int, coins []int) int {
	dp := make([]int, amount+1) // 存储到该金额可以组成个数
	dp[0] = 1                   // 没有金额，只有不选任何金币这1种情况
	for _, coin := range coins {
		for j := 0; j <= amount; j++ {
			if j-coin >= 0 {
				dp[j] += dp[j-coin]
			}
		}
	}
	return dp[amount]
}

func Test_change(t *testing.T) {
	tests := []struct {
		amount int
		coins  []int
		want   int
	}{
		{5, []int{1, 2, 5}, 4},
		{3, []int{2}, 0},
		{10, []int{10}, 1},
		{0, []int{10}, 1},
	}
	for _, tt := range tests {

		if got := change(tt.amount, tt.coins); got != tt.want {
			t.Errorf("change(%v,%v) = %v, want %v", tt.amount, tt.coins, got, tt.want)
		}

	}
}
