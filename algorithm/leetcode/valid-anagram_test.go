package leetcode

//242. 有效的字母异位词 https://leetcode-cn.com/problems/valid-anagram/
/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false

1 <= s.length, t.length <= 5 * 104
s 和 t 仅包含小写字母

进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

思路
用hash表或字母表统计字符出现次数, s串加 t串减, 最后看表的数据是否全是0
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	tab := [26]int{}
	for i := 0; i < len(s); i++ {
		tab[s[i]-'a']++
		tab[t[i]-'a']--
	}
	for i := range tab {
		if tab[i] != 0 {
			return false
		}
	}
	return true
}
