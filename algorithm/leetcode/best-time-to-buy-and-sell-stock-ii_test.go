package leetcode

// 122. 买卖股票的最佳时机 II https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

/*

给定一个数组 prices ，其中 prices[i] 表示股票第 i 天的价格。
在每一天，你可能会决定购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以购买它，然后在 同一天 出售。
返回 你能获得的 最大 利润 。

输入: prices = [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。

输入: prices = [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。

输入: prices = [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。


1 <= prices.length <= 3 * 104
0 <= prices[i] <= 104


思路:
动态规划思路
分为2种状态 持有股票,未持有股票
定义 dp[i][state]  表示第i天持有股票或持有股票的利润,0 表示未持有,1表是持有
状态转移方程:
// 第0天情况
dp[0][0] = 0 			// 未持有
dp[0][1] = -prices[0]	// 持有
// >0天 情况
dp[i][0] = max(dp[i-1][0],dp[i-1][1]+prices[i]) // 第i天未持有最大值,必然是 前一天未持有 与 前一天持有今天卖出 比较最大值
dp[i][1] = max(dp[i-1][1],dp[i-1][0]-prices[i]) // 第i天持有的最大值,必然是 前一天持有 与 前一天未持有今天买进 比较最大值

最后一天是不能买入的(买入是亏的) 所以答案是求 dp[len-1][0]
*/

func maxProfitII(prices []int) int {
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
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]*/

	// 因为只用到前面一个数字可以进行,内存优化
	notHave, have := 0, -prices[0]
	for i := 1; i < n; i++ {
		newNotHave := max(notHave, have+prices[i])
		newHave := max(have, notHave-prices[i])
		notHave = newNotHave
		have = newHave
	}
	return max(notHave, have)
}
