package leetcode

import "testing"

// 121. 买卖股票的最佳时机 https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

/*
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。


输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。

输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 没有交易完成, 所以最大利润为 0。

1 <= prices.length <= 105
0 <= prices[i] <= 104

思路:
动态规划
分为2种状态 持有股票,未持有股票
定义 dp[i][state]  表示第i天持有股票或持有股票的利润,0 表示未持有,1表是持有
状态转移方程:
// 第0天情况
dp[0][0] = 0 			// 未持有
dp[0][1] = -prices[0]	// 持有
// >0天 情况
dp[i][0] = max(dp[i-1][0],dp[i-1][1]+prices[i])	// 未持有, 只能是前一天未持有 或 前一天持有今日卖掉
dp[i][1] = max(dp[i-1][1],-prices[i] ) 			// 只能持有一次,看那一天的持有成本小

*/

func maxProfitI(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(prices)
	/*dp := make([][2]int, n)
	// 第0天情况
	dp[0][0] = 0          // 未持有
	dp[0][1] = -prices[0] // 持有
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]*/
	// 内存优化
	notHave, have := 0, -prices[0]
	for i := 1; i < n; i++ {
		notHave = max(notHave, have+prices[i])
		have = max(have, -prices[i])

	}
	return max(notHave, have)
}

func Test_maxProfitI(t *testing.T) {
	res := maxProfitI([]int{7, 1, 5, 3, 6, 4})
	t.Logf("%v\n", res)
}
