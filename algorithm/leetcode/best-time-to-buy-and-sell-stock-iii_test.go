package leetcode

// 123. 买卖股票的最佳时机 III https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/

/*
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

输入：prices = [3,3,5,0,0,3,1,4]
输出：6
解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。

输入：prices = [1,2,3,4,5]
输出：4
解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。

输入：prices = [7,6,4,3,1]
输出：0
解释：在这个情况下, 没有交易完成, 所以最大利润为 0。

输入：prices = [1]
输出：0

1 <= prices.length <= 105
0 <= prices[i] <= 105

思路:
最多完成2笔交易其实就是最多买入2次
从是否持有1只股票的角度来看可以分成2种的状态 持有(买入), 未持有(卖出)
从完成交易次数的角度来看又有两种状态 完成0次交易, 完成1次交易, 完成2次交易
所以组合起来有
0 完成0次交易 未持有
1 完成0次交易 持有
2 完成1次交易 未持有
3 完成1次交易 持有
4 完成2次交易 未持有
这5种状态

初始状态 持有状态
dp[i][1] = -prices[0]
dp[i][3] = -prices[0]

状态转移方程:
dp[i][0] = dp[i-1][0]  	完成0次交易未持有,到这部状态必须前面一个状态也是这个
dp[i][1] = max(dp[i-1][1],dp[i-1][0]-prices[i])    完成0次交易持有,到这个状态前一次是这个状态 或 完成0次交易未持有买入才能到达
dp[i][2] = max(dp[i-1][2],dp[i-1][1]+prices[i])
dp[i][3] = max(dp[i-1][3],dp[i-1][2]-prices[i])
dp[i][4] = max(dp[i-1][4],dp[i-1][3]+prices[i])

最后 max(dp[len-1][0...4])
*/

func maxProfitIII(prices []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(prices)
	dp := make([][5]int, n)
	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	res := 0
	for _, v := range dp[n-1] {
		res = max(res, v)
	}
	return res
}
