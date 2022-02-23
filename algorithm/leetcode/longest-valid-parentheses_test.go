package leetcode

import (
	"testing"
)

// 32. 最长有效括号 https://leetcode-cn.com/problems/longest-valid-parentheses/

/*
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"

输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"

输入：s = ""
输出：0

0 <= s.length <= 3 * 104
s[i] 为 '(' 或 ')'

思路:
1. 使用stack, stack中存储的是下标值
遇到 '(' 放入下标栈中
遇到 ')' 先弹栈
	检测stack是否是空,空就放入下标
	否则就计算有效长度,对比最大值
*/

func longestValidParentheses(s string) int {
	res := 0
	n := len(s)
	stack := []int{-1} // 初始需要是-1
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				availableSize := i - stack[len(stack)-1] // 有效长度
				if availableSize > res {
					res = availableSize
				}
			}
		}
	}
	return res
}

func Test_longestValidParentheses(t *testing.T) {
	//res := longestValidParentheses("()(()")
	//t.Logf("%v\n", res)

}
