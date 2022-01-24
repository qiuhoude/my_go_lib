package leetcode

import (
	"math"
	"testing"
)

// 322. 硬币问题 https://leetcode-cn.com/problems/coin-change/
/*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1

输入：coins = [2], amount = 3
输出：-1

输入：coins = [1], amount = 0
输出：0

输入：coins = [1], amount = 1
输出：1

输入：coins = [1], amount = 2
输出：2

1 <= coins.length <= 12
1 <= coins[i] <= 2^31 - 1
0 <= amount <= 10^4

思路:
可以将该问题转成金典的动态规划背包问题,在一堆物品中选出一些物品填满背包让其数量最小

自顶向下的思考:
币数虽然是无限的,但是总金额大小是有限的,其实将币数也变成有限
	定义函数 f(amount)   amount:总金额 返回最小数量
	min(for(f(amount-coins[i]))  // 每次都尝试每种币的金额,只要有解进行最小值比较

动态规划思路
dp[i] = min(dp[i-coin]+1, dp[i]) 当 i-coin >= 0

*/

// 动态规划思路
func coinCharge(coins []int, amount int) int {
	/*dp := make([]int, amount+1) //dp存储的是 在某个金额个数的最小值
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		min := math.MaxInt32
		for _, coin := range coins {
			if coin <= i && // 当前币种金额小于当前金额才有效 防止 i-coin <0
				dp[i-coin] != -1 && // 有解
				dp[i-coin]+1 < min { // 比较大小,找最小值
				min = dp[i-coin] + 1
			}
		}
		if min == math.MaxInt32 {
			dp[i] = -1
		} else {
			dp[i] = min
		}
	}*/

	dp := make([]int, amount+1) // 到该金额最小组成的最小值
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp[0] = 0
	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if i-coin >= 0 {
				// 选择i：金币个数+1,  不选择i
				dp[i] = min(dp[i-coin]+1, dp[i])
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

// 自顶向下的思考 leetcode 超时
func coinCharge2(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	memory := make([]int, amount+1)
	/*
		定义函数 f(amount)   amount:总金额 返回最小数量
		min(for(f(amount-coins[i]))
	*/
	var fn func(remain int) int
	fn = func(remain int) int {
		if remain < 0 {
			return -1
		}
		if remain == 0 {
			return 0
		}
		if memory[remain-1] != 0 {
			return memory[remain-1]
		}
		min := math.MaxInt32
		for i := range coins {
			r := fn(remain - coins[i]) // 选择 coins[i]
			if r != -1 && r < min {
				// 回溯
				min = r + 1 // 有解, 个数+1
			}
		}
		if min == math.MaxInt32 {
			min = -1
		}
		memory[remain-1] = min
		return min
	}
	return fn(amount)
}

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"[1,2,5] => 11", args{[]int{1, 2, 5}, 11}, 3},
		{"[2] => 3", args{[]int{2}, 3}, -1},
		{"[2] => 4", args{[]int{2}, 4}, 2},
		{"[1,3,5] => 11", args{[]int{1, 3, 5}, 11}, 3},
		{"[3,7,405,436] => 11", args{[]int{3, 7, 405, 436}, 8839}, 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinCharge(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinCharge() = %v, want %v", got, tt.want)
			}
		})
	}
}
