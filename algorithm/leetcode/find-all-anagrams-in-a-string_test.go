package leetcode

// 438. 找到字符串中所有字母异位词 https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/

/*
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。

1 <= s.length, p.length <= 3 * 104
s 和 p 仅包含小写字母

思路:
使用滑动窗口方式,右移一个格让新增字频+1, 最老的字频-1
如何判断一个词是另一个词异位词?
使用ascii码表统计字符串单个字出现的频率,比较单个字的频率
*/

func findAnagrams(s string, p string) []int {
	var res []int
	if len(p) > len(s) {
		return res
	}

	// 统计p串字的频率,和s子串
	freqS := [26]int{}
	freqP := [26]int{}
	for i := 0; i < len(p); i++ {
		freqS[s[i]-'a']++
		freqP[p[i]-'a']++
	}
	h := 0                                 // h头指针
	for t := len(p) - 1; t < len(s); t++ { // t 尾指针
		if freqP == freqS {
			res = append(res, h)
		}
		if t < len(s)-1 {
			freqS[s[t+1]-'a']++ //新增字频+1,
			freqS[s[h]-'a']--   // 最老的字频-1
			h++
		}
	}
	return res
}
