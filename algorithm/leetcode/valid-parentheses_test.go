package leetcode

// 20. 有效的括号 https://leetcode-cn.com/problems/valid-parentheses/

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

输入：s = "()"
输出：true

输入：s = "()[]{}"
输出：true

输入：s = "(]"
输出：false

输入：s = "([)]"
输出：false

输入：s = "{[]}"
输出：true

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成

思路:
遍历字符串, 使用栈保存入栈的值(栈中只有左边括号的值), 每次判断栈顶与对应的值是否匹配
*/

func isValid(s string) bool {
	var stack []byte
	for _, v := range []byte(s) {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else if v == ')' || v == ']' || v == '}' {
			if len(stack) == 0 {
				return false
			}
			var matchC byte
			switch v {
			case ')':
				matchC = '('
			case ']':
				matchC = '['
			case '}':
				matchC = '{'
			default:
				matchC = '_'
			}
			c := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if c != matchC {
				return false
			}
		}
	}
	return len(stack) == 0
}
