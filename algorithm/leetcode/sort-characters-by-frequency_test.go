package leetcode

// 451. 根据字符出现频率排序 https://leetcode-cn.com/problems/sort-characters-by-frequency/

/*
给定一个字符串，请将字符串里的字符按照出现的频率降序排列。

输入:
"tree"
输出:
"eert"
'e'出现两次，'r'和't'都只出现一次。
因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。

输入:
"cccaaa"
输出:
"cccaaa"
'c'和'a'都出现三次。此外，"aaaccc"也是有效的答案。
注意"cacaca"是不正确的，因为相同的字母必须放在一起。

输入:
"Aabb"
输出:
"bbAa"
此外，"bbaA"也是一个有效的答案，但"Aabb"是不正确的。
注意'A'和'a'被认为是两种不同的字符。

思路
使用hash标记记录字符出现的频率,进行按出现次数分桶,最后把桶中的数据倒出来

*/

func frequencySort(s string) string {
	sRune := []rune(s)
	if len(sRune) < 2 {
		return string(sRune)
	}

	maxFreq := 0                  // 出现最大值
	freqTab := make(map[rune]int) // 字符出现的频率表
	for i := range sRune {
		freqTab[sRune[i]]++
		tmpFreq := freqTab[sRune[i]]
		if tmpFreq > maxFreq {
			maxFreq = tmpFreq
		}
	}
	bucket := make([][]rune, maxFreq+1) // 按照频率分桶,频率最高的字符再最后面
	for ch, cnt := range freqTab {
		bucket[cnt] = append(bucket[cnt], ch)
	}

	// 把桶里的元素倒出来
	resRune := make([]rune, 0, len(sRune))
	for i := maxFreq; i > 0; i-- {
		for _, ch := range bucket[i] { // 每个字符添加i次
			for j := 0; j < i; j++ {
				resRune = append(resRune, ch)
			}
		}
	}
	return string(resRune)
}
