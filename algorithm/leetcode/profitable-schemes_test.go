package leetcode

// 879. 盈利计划 https://leetcode-cn.com/problems/profitable-schemes/

/*
集团里有 n 名员工，他们可以完成各种各样的工作创造利润。

第 i 种工作会产生 profit[i] 的利润，它要求 group[i] 名成员共同参与。如果成员参与了其中一项工作，就不能参与另一项工作。
工作的任何至少产生 minProfit 利润的子集称为 盈利计划 。并且工作的成员总数最多为 n 。
有多少种计划可以选择？因为答案很大，所以 返回结果模 10^9 + 7 的值。

输入：n = 5, minProfit = 3, group = [2,2], profit = [2,3]
输出：2
解释：至少产生 3 的利润，该集团可以完成工作 0 和工作 1 ，或仅完成工作 1 。
总的来说，有两种计划。

输入：n = 10, minProfit = 5, group = [2,3,5], profit = [6,7,8]
输出：7
解释：至少产生 5 的利润，只要完成其中一种工作就行，所以该集团可以完成任何工作。
有 7 种可能的计划：(0)，(1)，(2)，(0,1)，(0,2)，(1,2)，以及 (0,1,2) 。

1 <= n <= 100
0 <= minProfit <= 100
1 <= group.length <= 100
1 <= group[i] <= 100
profit.length == group.length
0 <= profit[i] <= 100

思路:
将问题转换成多维背包问题
group 和 profit 就是背包的物品,
两个条件限制: 人数 n, 最至少利润 minProfit
dp[i][j][k] i表示任务编号,j表示人数, 利润k,存储的是满足条件至少利润为k方案数总数

初始化
不满足条件
dp[i][j][k] = dp[i-1][j][k]
满足条件, 其中k-profit[i]<=0表示利润至少为k,所以0表示利润至少为k
dp[i][j][k]  =  dp[i-1][j][k] + dp[i-1][j-group[i]][max(0,k-profit[i])]

*/

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	const mod = 1e9 + 7
	maxFn := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	size := len(group)
	dp := make([][][]int, size+1)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, minProfit+1)
		}
	}
	for j := 0; j <= n; j++ {
		dp[0][j][0] = 1
	}
	for i := 1; i <= size; i++ {
		// 人数条件和利润
		m, p := group[i-1], profit[i-1]
		for j := 0; j <= n; j++ {
			for k := 0; k <= minProfit; k++ {
				if j-m >= 0 { // 人数条件满足
					ck := maxFn(0, k-p) //利润条件至少为k; 其中k-p<=0表示利润至少为k,所以0表示利润至少为k
					dp[i][j][k] = (dp[i-1][j][k] + dp[i-1][j-m][ck]) % mod
				} else {
					dp[i][j][k] = dp[i-1][j][k] // 不满足 抄上一行
				}
			}
		}
	}
	return dp[size][n][minProfit]

}
