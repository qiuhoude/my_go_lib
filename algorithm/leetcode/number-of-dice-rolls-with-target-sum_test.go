package leetcode

import "testing"

// 1155. 掷骰子的N种方法 https://leetcode-cn.com/problems/number-of-dice-rolls-with-target-sum/

/*
这里有 d 个一样的骰子，每个骰子上都有 f 个面，分别标号为 1, 2, ..., f。
我们约定：掷骰子的得到总点数为各骰子面朝上的数字的总和。
如果需要掷出的总点数为 target，请你计算出有多少种不同的组合情况（所有的组合情况总共有 f^d 种），模 10^9 + 7 后返回。


输入：d = 1, f = 6, target = 3
输出：1

输入：d = 2, f = 6, target = 7
输出：6

输入：d = 2, f = 5, target = 10
输出：1

输入：d = 1, f = 2, target = 3
输出：0

输入：d = 30, f = 30, target = 500
输出：222616187

1 <= d, f <= 30
1 <= target <= 1000

思路:
有d个骰子, 每个骰子有f个面，取值可以是 1~f,
将问题转换成多维背包，每个骰子就是一个背包，有d个背包
每个背包的问题转换成 0-1背包问题 ,
dp[i][j]  i 骰子个数，j 目标点数，k为每次的点数取值范围在 1~min(f,j-k)
dp[i][j]  += dp[i-1][j-k]

*/
func numRollsToTarget(d int, f int, target int) int {
	dp := make([][]int, d+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	const mod = 1e9 + 7
	dp[0][0] = 1
	for i := 1; i <= d; i++ {
		for j := 1; j <= target; j++ {
			for k := 1; k <= f && k <= j; k++ { // 1~f点
				dp[i][j] += dp[i-1][j-k]
				dp[i][j] = dp[i][j] % mod
			}
		}
	}
	return dp[d][target]
}

func Test_numRollsToTarget(t *testing.T) {

	tests := []struct {
		d      int
		f      int
		target int
		want   int
	}{
		{30, 30, 500, 222616187},
		{1, 6, 3, 1},
		{2, 5, 10, 1},
		{1, 2, 3, 0},
	}
	for _, tt := range tests {
		if got := numRollsToTarget(tt.d, tt.f, tt.target); got != tt.want {
			t.Errorf("numRollsToTarget(%v,%v,%v) = %v, want %v", tt.d, tt.f, tt.target, got, tt.want)
		}

	}

}
