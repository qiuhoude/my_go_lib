package leetcode

import (
	"strings"
	"testing"
)

// 139. 单词拆分 https://leetcode-cn.com/problems/word-break/

/*
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。


输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false


1 <= s.length <= 300
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 20
s 和 wordDict[i] 仅有小写英文字母组成
wordDict 中的所有字符串 互不相同

思路:

动态规划思路:
看上去有点像 322. 硬币问题 , 无限元素的背包问题

n=len(s), 创建 n+1 dp表, 存放着 到该位置是否能拼接成功
dp[0] = true
dp[i]  = for(dp[i-len(word)] && word==s[i-len(word):i])
dp[n] 就是解

*/

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for _, word := range wordDict {
			wl := len(word)             // 单词的长度
			if i-wl >= 0 && dp[i-wl] && // 排除这个单词长度,之前也能拼接成功就继续 比较字符串是否相等
				strings.EqualFold(word, string(s[i-wl:i])) {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

func Test_wordBreak(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		expected bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}
	for _, tt := range tests {
		if got := wordBreak(tt.s, tt.wordDict); got != tt.expected {
			t.Errorf("wordBreak(%v)  got %v, expected %v", tt.s, got, tt.expected)
		}
	}
}
