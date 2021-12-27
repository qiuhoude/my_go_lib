package leetcode

import (
	"strconv"
	"strings"
)

// 93. 复原 IP 地址 https://leetcode-cn.com/problems/restore-ip-addresses/

/*
有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。
你不能重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。

输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]

输入：s = "0000"
输出：["0.0.0.0"]

输入：s = "1111"
输出：["1.1.1.1"]

输入：s = "010010"
输出：["0.10.0.10","0.100.1.0"]

输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

0 <= s.length <= 20
s 仅由数字组成

思路:
和17题的思路类似, 使用递归回溯的方式
把ip看作一个树, 构成ip段的每段数字看作树中的节点, 每个节点的数字最多由3个数字组成,从第一个节点开始开始枚举可以组成第一个数字的所有情况
依次类推进行递归下去,到达最后一个节点最后如数字不够用就不构成ip,不加入到结果列表中
其实也是枚举树的每个路径选出符合的路径
*/

func restoreIpAddresses(s string) []string {
	if len(s) < 4 { // 小于4个不能构成ip
		return nil
	}
	var res []string
	recursionBuildIp(s, nil, 0, 1, &res)
	return res
}

/*
s 原始字符串,
path ip的前缀,
curIndex 当前位置的下标(字符串s),
depth 当前递归深度最多是4,
res 结果列表
*/
func recursionBuildIp(s string, path []string, curIndex, depth int, res *[]string) {
	if depth == 4 || curIndex >= len(s) { // 最后一个数字
		// 判断结果是否满足
		numStr := s[curIndex:]
		if len(numStr) > 1 && numStr[0] == '0' { // len>1 并且 0 开头的不是不合法的
			return
		}
		n, err := strconv.Atoi(numStr)
		if err == nil && n >= 0 && n <= 255 {
			path = append(path, numStr)
			*res = append(*res, strings.Join(path, "."))
		}
		return
	}
	numCnt := 0 // 数字数量
	var sb strings.Builder
	for i := curIndex; i < len(s) && numCnt < 3; i++ {
		sb.WriteByte(s[i])
		numStr := sb.String()
		if len(numStr) > 1 && numStr[0] == '0' { // len>1 并且 0 开头的不是不合法的
			break
		}
		n, err := strconv.Atoi(numStr)
		if err != nil {
			break
		}
		if n < 0 || n > 255 {
			break
		}
		// 递归下去
		recursionBuildIp(s, append(path, numStr), i+1, depth+1, res)
		numCnt++
	}
}

//func Test_restoreIpAddresses(t *testing.T) {
//	ipArr := restoreIpAddresses("101023")
//	t.Logf("%v\n", ipArr)
//}
