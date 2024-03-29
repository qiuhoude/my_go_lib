package leetcode

import (
	"testing"
	"unicode"
)

//125. 验证回文串 https://leetcode-cn.com/problems/valid-palindrome/

/*
11
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。

输入: "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串

输入: "race a car"
输出: false
解释："raceacar" 不是回文串

1 <= s.length <= 2 * 105
字符串 s 由 ASCII 字符组成

思路: 也是使用对撞指针进行求解, h,t
*/

func isPalindromeii(s string) bool {
	runes := []rune(s)
	h, t := 0, len(runes)-1 // 前后指针
	for h < t {
		hc, tc := runes[h], runes[t]
		if !unicode.IsDigit(hc) && !unicode.IsLetter(hc) { // 不是数字又不是字母
			h++
			continue
		}
		if !unicode.IsDigit(tc) && !unicode.IsLetter(tc) { // 不是数字又不是字母
			t--
			continue
		}
		lhc, ltc := unicode.ToLower(hc), unicode.ToLower(tc)
		if lhc != ltc {
			return false
		}
		h++
		t--
	}
	return true
}

func Test_isPalindromeii(t *testing.T) {
	res := isPalindromeii(`A man, a plan, a canal: Panama`)
	if !res {
		t.Logf("expected:%v, got:%v", true, res)
	}
}
