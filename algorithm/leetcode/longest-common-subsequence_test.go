package leetcode

import (
	"fmt"
)

// 1143. 最长公共子序列 https://leetcode-cn.com/problems/longest-common-subsequence/
/*
给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace" ，它的长度为 3 。

输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc" ，它的长度为 3 。

输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0 。

1 <= text1.length, text2.length <= 1000
text1 和 text2 仅由小写英文字符组成。

思路:

使用动态规划的思路解决
dp存储的到该位置之前最长子序列
当text1[i-1] != text2[j-1] 时 dp[i][j] = max(dp[i-1][j],dp[i][j-1])
当text1[i-1] == text2[j-1] 时 dp[i][j] = dp[i-1][j-1]+1
*/

func longestCommonSubsequence(text1 string, text2 string) int {
	maxFn := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n, m := len(text1), len(text2)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxFn(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	//printLCS(text1, text2, dp)
	return dp[n][m]
}

// 通过dp表找到最长子序列实际的值
func printLCS(text1 string, text2 string, dp [][]int) {
	n, m := len(text1), len(text2)
	subLen := dp[n][m]
	subArr := make([]byte, subLen)
	k, i, j := subLen-1, n, m
	for k >= 0 {
		if dp[i][j] == dp[i-1][j] {
			i--
		} else if dp[i][j] == dp[i][j-1] {
			j--
		} else { // 相等的情况
			subArr[k] = text1[i-1] // or subArr[k] = text2[j-1]
			i--
			j--
			k--
		}
	}
	fmt.Printf("%s\n", string(subArr))
}

//func Test_longestCommonSubsequence(t *testing.T) {
//	longestCommonSubsequence("abcde", "ace")
//}
