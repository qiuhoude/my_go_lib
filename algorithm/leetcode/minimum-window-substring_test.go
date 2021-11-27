package leetcode

import (
	"math"
	"strings"
	"testing"
)

// 76. 最小覆盖子串 https://leetcode-cn.com/problems/minimum-window-substring/

/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"

输入：s = "a", t = "a"
输出："a"

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。

1 <= s.length, t.length <= 105
s 和 t 由英文字母组成

思路:
使用滑动窗口 加 字频统计的方式

先统计t串和s[0,len(t)]的字频,s[start,end+1]就形成滑动窗口(start=0,end=len(t)-1)
比较是否满足窗口条件,可以只用比较t串字符的字的字频
如果s[start,end+1]窗口满足条件则让start++进行缩小窗口大小,不满足条件就进行end++扩大窗口,最小长度的位置起始结束位置,
如果end+1-start等于len(t)说明是的子串直接返回
*/

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	freqT := [256]int{}   //  t串的字频率
	freqS := [256]int{}   //  s子串的字频率
	var tIdInAscii []byte // t串字的在ascii表索引,用于比较freqT和freqS使用
	for i := 0; i < len(t); i++ {
		if freqT[t[i]] == 0 { // 该字再t串中首次出现
			tIdInAscii = append(tIdInAscii, t[i])
		}
		freqT[t[i]]++
		freqS[s[i]]++
	}
	resStart, minWinLen := -1, math.MaxInt32 // 记录最小窗口
	start, end := 0, len(t)-1                // 滑动窗口指针

	for end < len(s) {
		for cmpFreq(freqT, freqS, tIdInAscii) && start <= end {
			curLen := end - start + 1 //当前窗口长度
			// 记录最小子串位置
			if curLen < minWinLen {
				minWinLen = curLen
				resStart = start
				if curLen == len(t) { // 已经找到最小直接返回
					return s[resStart : resStart+minWinLen]
				}
			}
			freqS[s[start]]-- // 缩小窗口
			start++
		}

		end++
		if end < len(s) {
			// 新增的字符
			freqS[s[end]]++
		}
	}
	if resStart == -1 { // 无效
		return ""
	}
	return s[resStart : resStart+minWinLen]
}

// 满足条件返回true
func cmpFreq(freqT, freqS [256]int, tIdInAscii []byte) bool {
	for _, c := range tIdInAscii {
		if freqT[c] > freqS[c] {
			return false
		}
	}
	return true
}

func Test_minWindow(t *testing.T) {
	tests := []struct {
		expected string
		arg1     string
		arg2     string
	}{
		{"BANC", "ADOBECODEBANC", "ABC"},
		{"a", "a", "a"},
		{"", "a", "aa"},
	}
	for _, tt := range tests {
		res := minWindow(tt.arg1, tt.arg2)
		if !strings.EqualFold(tt.expected, res) {
			t.Logf("expected:%v but got:%v", tt.expected, res)
		}
	}
}
