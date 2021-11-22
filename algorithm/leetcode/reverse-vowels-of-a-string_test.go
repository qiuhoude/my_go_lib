package leetcode

import "unicode"

// 345. 反转字符串中的元音字母 https://leetcode-cn.com/problems/reverse-vowels-of-a-string/

/*
给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现。

输入：s = "hello"
输出："holle"

输入：s = "leetcode"
输出："leotcede"

1 <= s.length <= 3 * 105
s 由 可打印的 ASCII 字符组成

思路:
对撞指针 h,t找到h,t是元音字符的进行对调
*/

func reverseVowels(s string) string {
	runes := []rune(s)
	h, t := 0, len(runes)-1 // 前后指针
	for h < t {
		if !isVowelRuneIgnoreCase(runes[h]) {
			h++
			continue
		}
		if !isVowelRuneIgnoreCase(runes[t]) {
			t--
			continue
		}
		runes[h], runes[t] = runes[t], runes[h]
		h++
		t--
	}
	return string(runes)
}

func isVowelRuneIgnoreCase(r rune) bool {
	rLower := unicode.ToLower(r)
	switch rLower {
	case 'a':
	case 'e':
	case 'i':
	case 'o':
	case 'u':
	default:
		return false
	}
	return true
}
