package leetcode

// 49. 字母异位词分组 https://leetcode-cn.com/problems/group-anagrams/

/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母都恰好只用一次。

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

输入: strs = [""]
输出: [[""]]

输入: strs = ["a"]
输出: [["a"]]

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母

思路:
查表法, 给每个字符串创建一个26大小的数组,数组相等表示字母异位词,
将26大小的数组当作key存入 strSlice作为val存入map
*/

func groupAnagrams(strs []string) [][]string {
	tab := make(map[[26]int][]string, len(strs))

	var emptyStrs []string
	for _, str := range strs {
		if str == "" {
			emptyStrs = append(emptyStrs, "")
		} else {
			var asciiArr [26]int
			for j := 0; j < len(str); j++ {
				asciiArr[str[j]-'a']++
			}
			tab[asciiArr] = append(tab[asciiArr], str)
		}
	}

	res := make([][]string, 0, len(strs)+1)
	if len(emptyStrs) > 0 { // 添加""串
		res = append(res, emptyStrs)
	}
	for _, v := range tab {
		res = append(res, v)
	}
	return res
}
