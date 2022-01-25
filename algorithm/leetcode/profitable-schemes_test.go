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
转换成 求解 在group选择子集 G, 在profit中选择子集 P, 条件是 sum(G)<=n && sum(P)>=minProfit,有多少种方案
将问题转换成多维背包问题
group 和 profit 就是背包的物品,
两个条件限制: 人数 n, 最至少利润 minProfit

定义 dp[i][j][k] 表示前i种工作,j表示人数,利润k, 存储利润等于k计算数量
选择i
dp[i][j][k] = dp[i-1][j][k] + dp[i-1][j-group[i-1]][k-profit[i-1]] (j-group[i-1]>=0 && k-profit[i-1]>=0)
不选择
dp[i][j][k] = dp[i-1][j][k]
最后求解 minProfit~sumProfit方案的总和
sum(dp[i][j][minProfit~sumProfit])


优化思路2:
dp[i][j][k] 表示前i种工作,j表示人数, 利润k,存储的是满足条件至少利润为k方案数总数
不满足条件时:
dp[i][j][k] = dp[i-1][j][k]
满足条件, 其中k-profit[i]<=0表示利润至少为k,所以0表示利润至少为k
dp[i][j][k]  =  dp[i-1][j][k] + dp[i-1][j-group[i-1]][max(0,k-profit[i-1])]

*/

func profitableSchemes0(n int, minProfit int, group []int, profit []int) int {
	const mod = 1e9 + 7
	sumProfit := 0
	for _, v := range profit {
		sumProfit += v
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, sumProfit+1)
		dp[i][0] = 1 //  利润为0时,方案数量时1 ,不选任何工作
	}

	size := len(group)
	for i := 1; i <= size; i++ {
		// 人数条件和利润
		m, p := group[i-1], profit[i-1]
		for j := n; j >= m; j-- {
			for k := p; k <= sumProfit; k++ {
				dp[j][k] += dp[j-m][k-p] % mod //选择i
			}
		}
	}
	//最后求解 minProfit~sumProfit方案的总和
	ans := 0
	for i := minProfit; i <= sumProfit; i++ {
		ans = (ans + dp[n][i]) % mod
	}

	return ans
}

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
	for j := 0; j <= n; j++ { //  利润为0时,方案数量时1 ,不选任何工作
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
