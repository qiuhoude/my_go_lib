package leetcode

import "testing"

// 17. 电话号码的字母组合 https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

输入：digits = ""
输出：[]

输入：digits = "2"
输出：["a","b","c"]

0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字

思路:
使用递归回溯方式,深度优先, 可以想象字符串的长度是一个tree的深度,每个数字都表示多个字母的节点,然后去遍历整个tree,
列出根节点到接子节点的所有路径, 最后转换成求tree的所有路径的问题

*/

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	letterMap := make(map[rune][]rune)
	letterMap['2'] = []rune{'a', 'b', 'c'}
	letterMap['3'] = []rune{'d', 'e', 'f'}
	letterMap['4'] = []rune{'g', 'h', 'i'}
	letterMap['5'] = []rune{'j', 'k', 'l'}
	letterMap['6'] = []rune{'m', 'n', 'o'}
	letterMap['7'] = []rune{'p', 'q', 'r', 's'}
	letterMap['8'] = []rune{'t', 'u', 'v'}
	letterMap['9'] = []rune{'w', 'x', 'y', 'z'}
	var s []rune
	var res []string
	findCombination(0, []rune(digits), letterMap, s, &res)
	return res
}

func findCombination(index int, text []rune, letterMap map[rune][]rune, s []rune, res *[]string) {
	if index == len(text) {
		// 递归的出口就是解
		*res = append(*res, string(s))
		return
	}
	arr, ok := letterMap[text[index]]
	if !ok {
		// 没有找到就直接继续
		findCombination(index+1, text, letterMap, s, res)
		return
	}
	for i := range arr {
		findCombination(index+1, text, letterMap, append(s, arr[i]), res)
	}
}

func TestLetterCombinations(t *testing.T) {
	res := letterCombinations("239")
	t.Log(res)
}
