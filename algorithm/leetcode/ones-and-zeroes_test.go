package leetcode

import (
	"strings"
	"testing"
)

// 474. 一和零 https://leetcode-cn.com/problems/ones-and-zeroes/

/*
给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。
如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。

输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
输出：4
解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 n 的值 3 。

输入：strs = ["10", "0", "1"], m = 1, n = 1
输出：2
解释：最大的子集是 {"0", "1"} ，所以答案是 2 。

1 <= strs.length <= 600
1 <= strs[i].length <= 100
strs[i] 仅由 '0' 和 '1' 组成
1 <= m, n <= 100

思路:

动态规划:
问题转换成多维度条件的背包问题
dp[l][m][n], l:表示长度, m,n双条件, 存储的时最大长度
参数的取值范围 l:[0,len(strs)],m:[0,m],n[0,n]
分两种:
1. 不能选择 strs[l] =  dp[l-1][m][n] (m-zeroCnt<0 | n-oneCnt<0)
2. 能选择 strs[l] =
max(
dp[l-1][m][n],					// 不选择
dp[l-1][m-zeroCnt][n-oneCnt]+1   // 选择
)

优化:长度这一栏每次只用到本排和上一排的数据,所以可以将长度这个维度去掉

*/

// 优化
func findMaxForm2(strs []string, m int, n int) int {
	maxFn := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, s := range strs {
		zeroCnt := strings.Count(s, "0") // '0'  的数量
		oneCnt := len(s) - zeroCnt       // '1'的数量
		for z := m; z >= zeroCnt; z-- {
			for o := n; o >= oneCnt; o-- {
				//if z-zeroCnt >= 0 && o-oneCnt >= 0 {
				dp[z][o] = maxFn(dp[z][o], dp[z-zeroCnt][o-oneCnt]+1)
				//}
			}
		}
	}
	return dp[m][n]
}

func findMaxForm(strs []string, m int, n int) int {
	maxFn := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	size := len(strs)
	dp := make([][][]int, size+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}
	// 长度是0的时候, dp[0][m][n] 必然都是0
	for l := 1; l <= size; l++ {
		zeroCnt := strings.Count(strs[l-1], "0") // '0'  的数量
		oneCnt := len(strs[l-1]) - zeroCnt       // '1'的数量
		for z := 0; z <= m; z++ {
			for o := 0; o <= n; o++ {
				if z-zeroCnt >= 0 && o-oneCnt >= 0 {
					dp[l][z][o] = maxFn(dp[l-1][z][o], dp[l-1][z-zeroCnt][o-oneCnt]+1)
				} else {
					dp[l][z][o] = dp[l-1][z][o] //照抄上一行
				}
			}
		}
	}
	return dp[size][m][n]
}

func Test_findMaxForm(t *testing.T) {
	tests := []struct {
		strs     []string
		m        int
		n        int
		expected int
	}{
		{[]string{"10", "0001", "111001", "1", "0"}, 50, 50, 5},
		{[]string{"10", "0001", "111001", "1", "0"}, 4, 3, 3},
		{[]string{"10", "0001", "111001", "1", "0"}, 1, 1, 2},
		{[]string{"10", "0001", "111001", "1", "0"}, 5, 3, 4},
		{[]string{"10", "0", "1"}, 1, 1, 2},
	}
	for _, tt := range tests {
		if got := findMaxForm2(tt.strs, tt.m, tt.n); got != tt.expected {
			t.Errorf("findMaxForm(%v,%v)  got %v, expected %v", tt.m, tt.n, got, tt.expected)
		}
	}
}
