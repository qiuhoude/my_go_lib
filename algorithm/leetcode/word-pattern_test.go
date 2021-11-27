package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

// 290. 单词规律 https://leetcode-cn.com/problems/word-pattern/
/**
给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。
这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

输入: pattern = "abba", str = "dog cat cat dog"
输出: true

输入:pattern = "abba", str = "dog cat cat fish"
输出: false

输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false

输入: pattern = "abba", str = "dog dog dog dog"
输出: false

你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。

思路:
将pattern的字符与s字符串切割成的字符串数组做一个映射关系,检测后面出现相同的字符与对应的字符是否匹配
*/

func wordPattern(pattern string, s string) bool {
	pSize := len(pattern)
	strArr := strings.Split(s, " ")
	sSize := len(strArr)
	if pSize != sSize {
		return false
	}
	tabC := make(map[byte]string)
	tabS := make(map[string]byte)
	for i := 0; i < sSize; i++ {
		c, str := pattern[i], strArr[i]

		ss, ok := tabC[c]
		if ok {
			if ss != str { // 检测不相等
				return false
			}
		} else { // 不包含进行填充
			tabC[c] = str
		}

		cc, ok := tabS[str]
		if ok {
			if cc != c { // 检测不相等
				return false
			}
		} else { // 不包含进行填充
			tabS[str] = c
		}
	}
	return true
}

func Test_wordPattern(t *testing.T) {
	tests := []struct {
		expected bool
		arg1     string
		arg2     string
	}{
		{true, "abba", "dog cat cat dog"},
		{false, "abba", "dog cat cat fish"},
		{false, "aaaa", "dog cat cat dog"},
		{false, "abba", "dog dog dog dog"},
	}
	for _, tt := range tests {
		res := wordPattern(tt.arg1, tt.arg2)
		t.Run(fmt.Sprintf("%v <-> %v", tt.arg1, tt.arg2), func(t *testing.T) {
			if res != tt.expected {
				t.Logf("expected:%v but got:%v", tt.expected, res)
			}
		})

	}

}
