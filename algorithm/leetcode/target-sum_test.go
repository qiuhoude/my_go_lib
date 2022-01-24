package leetcode

// 494. 目标和 https://leetcode-cn.com/problems/target-sum/

/*
给你一个整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：

例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3

输入：nums = [1], target = 1
输出：1

1 <= nums.length <= 20
0 <= nums[i] <= 1000
0 <= sum(nums[i]) <= 1000
-1000 <= target <= 1000

思路:
可以将问题转换成 0-1背包问题

原问题等同于： 找到nums一个正子集P和一个负子集N(P和N存的都是正整数) sum(nums) = sum(P) + sum(N)
使得总和等于target。即sum(P) - sum(N) == target，

sum(N) + sum(P) - sum(N) == target+ sum(N)  两边都加上sum(N)
sum(nums) - sum(N) = target+ sum(N) => 2 * sum(N) = sum(nums) - target
则问题转换为：存在多少个子集N， sum(N) = (sum(nums) - target)/2 其中 sum(nums) - target 必须时偶数

dp[i][j]表示前i个元素有多少个目标和为j的子集个数。dp[0][0] = 1
    1. dp[i][j] = dp[i-1][j]
    2. 如果nums[0...i-2]存在目标和为j-nums[i-1]的子集，则dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
*/

func findTargetSumWays(nums []int, target int) int {
	length := len(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	if diff < 0 || diff%2 != 0 {
		return 0
	}
	nega := diff / 2 // sum(子集N)

	/*dp := make([][]int, length+1)
	for i := range dp {
		dp[i] = make([]int, nega+1)
	}
	dp[0][0] = 1
	for i := 1; i <= length; i++ {
		for j := 0; j <= nega; j++ {
			if j-nums[i-1] >= 0 {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[length][nega]*/

	//因为只依赖于前一行所以可以优化成单数组
	dp := make([]int, nega+1)
	dp[0] = 1
	for i := 1; i <= length; i++ {
		for j := nega; j >= 0; j-- {
			if j-nums[i-1] >= 0 {
				dp[j] = dp[j] + dp[j-nums[i-1]]
			}
		}
	}
	return dp[nega]

}
