package leetcode

//3. 无重复字符的最长子串 https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

输入: s = ""
输出: 0

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成

思路:
只有ASCII 无需转rune,
使用滑动窗口, 保证窗口内的数据满足条件(无重复字符), 扫描整个数组就可以找到
判断重复的两种方式
1 每次判断新加入的元素是否与窗口中的数据有重复,有重复 ph头指针直接移动到 重复位置+1
2 判断是否重复: 创建一个ascii表记录窗口中字符出现次数
*/

func lengthOfLongestSubstring1(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	maxLen := 1
	tmpLen := 1
	star := 0
	for end := 1; end < len(s); end++ {
		tmpc := s[end]
		for i := star; i < end; i++ {
			if tmpc == s[i] {
				star = i + 1 // 重复元素的位置+1
				tmpLen = end - star
				break
			}
		}
		tmpLen++
		if tmpLen > maxLen {
			maxLen = tmpLen
		}
	}
	return maxLen
}
func lengthOfLongestSubstring2(s string) int {
	size := len(s)
	if size <= 1 {
		return size
	}
	freq := [256]int{} // 记录窗口中字符出现的频率

	ph, pt := 0, -1 // 滑动窗口指针 [ph,pt]
	maxSubLen := 1  // 最长子串的长度
	for ph < size {
		if pt < size-1 && freq[s[pt+1]] == 0 { // pt条件是之前没有出现过
			pt++
			freq[s[pt]]++
		} else {
			freq[s[ph]]--
			ph++
		}
		tmpLen := pt - ph + 1
		if tmpLen > maxSubLen {
			maxSubLen = tmpLen
		}
	}

	return maxSubLen
}
