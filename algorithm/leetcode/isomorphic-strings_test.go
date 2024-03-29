package leetcode

// 205. 同构字符串 https://leetcode-cn.com/problems/isomorphic-strings/

/*
给定两个字符串 s 和 t，判断它们是否是同构的。
如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

输入：s = "egg", t = "add"
输出：true

输入：s = "foo", t = "bar"
输出：false
示例 3：

输入：s = "paper", t = "title"
输出：true

可以假设 s 和 t 长度相同。

思路:
使用两个map进行相互映射,进行检测

*/

func isIsomorphic(s string, t string) bool {
	runeS, runeT := []rune(s), []rune(t)
	if len(runeS) != len(runeT) {
		return false
	}
	s2tM, t2sM := make(map[rune]rune, len(s)), make(map[rune]rune, len(s))

	for i := 0; i < len(runeS); i++ {
		sc, tc := runeS[i], runeT[i]

		tcc, ok := s2tM[sc] // sc->tc
		if ok {
			if tcc != tc { // 检测不相等
				return false
			}
		} else { // 不包含进行填充
			s2tM[sc] = tc
		}

		scc, ok := t2sM[tc] // tc->sc
		if ok {
			if scc != sc { // 检测不相等
				return false
			}
		} else { // 不包含进行填充
			t2sM[tc] = sc
		}
	}
	return true
}
