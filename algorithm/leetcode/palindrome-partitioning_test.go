package leetcode

import "testing"

// 131. 分割回文串 https://leetcode-cn.com/problems/palindrome-partitioning/
/*
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
回文串 是正着读和反着读都一样的字符串。


输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]

输入：s = "a"
输出：[["a"]]

1 <= s.length <= 16
s 仅由小写英文字母组成

思路:
与93题思路, 将问题转换成树形问题, 把字符串切割位置当作节点,通过条件剪枝优化程序,其实在树上来找符合条件的路径

第一个节点 s[:1],s[:2],...,s[:len-1], 检测这些否是回文串,是的就继续递归下去
可以使用hash表记录检测过的字符串是否是回文串
*/

func palindrome_partition(s string) [][]string {
	n := len(s)
	// 记录 开始下标都结束下标的字符字符串是否是回文串 ,0 表示尚未搜索，1 表示是回文串，-1 表示不是回文串
	table := make([][]int8, n)
	for i := range table {
		table[i] = make([]int8, n)
	}
	var isPalindromeFn func(start, end int) int8
	isPalindromeFn = func(start, end int) int8 {
		if start >= end {
			return 1
		}
		if table[start][end] != 0 {
			return table[start][end]
		}
		table[start][end] = -1
		if s[start] == s[end] {
			table[start][end] = isPalindromeFn(start+1, end-1)
		}
		return table[start][end]

	}
	var res [][]string
	recursionPalindrome(s, nil, 0, &res, isPalindromeFn)
	return res
}

func recursionPalindrome(s string, path []string, startIndex int, res *[][]string, isPalindromeFn func(start, end int) int8) {
	if startIndex == len(s) {
		// 结束记录结果
		*res = append(*res, append([]string(nil), path...))
		return
	}
	for endIndex := startIndex; endIndex < len(s); endIndex++ {
		if isPalindromeFn(startIndex, endIndex) > 0 {
			subStr := s[startIndex : endIndex+1]
			recursionPalindrome(s, append(path, subStr), endIndex+1, res, isPalindromeFn)
		}
	}
}

func Test_palindrome_partition(t *testing.T) {
	//res := palindrome_partition("ababbbabbaba")
	res := palindrome_partition("aba")
	t.Logf("%q\n", res)
	/*
		//splits := []string{}
		//var dfs func(int)
		//dfs = func(startIndex int) {
		//	if startIndex == n {
		//		res = append(res, append([]string(nil), splits...))
		//		return
		//	}
		//	for endIndex := startIndex; endIndex < n; endIndex++ {
		//		if isPalindromeFn(startIndex, endIndex) > 0 {
		//			splits = append(splits, s[startIndex:endIndex+1])
		//			dfs(endIndex + 1)
		//			splits = splits[:len(splits)-1]
		//		}
		//	}
		//}
		//dfs(0)
	*/
}
