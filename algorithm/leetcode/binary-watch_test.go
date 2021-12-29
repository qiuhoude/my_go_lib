package leetcode

import (
	"fmt"
	"testing"
)

// 401. 二进制手表 https://leetcode-cn.com/problems/binary-watch/

/*
二进制手表顶部有 4 个 LED 代表 小时（0-11），底部的 6 个 LED 代表 分钟（0-59）。每个 LED 代表一个 0 或 1，最低位在右侧。
例如，下面的二进制手表读取 "3:25" 。

给你一个整数 turnedOn ，表示当前亮着的 LED 的数量，返回二进制手表可以表示的所有可能时间。你可以 按任意顺序 返回答案。

小时不会以零开头：

例如，"01:00" 是无效的时间，正确的写法应该是 "1:00" 。
分钟必须由两位数组成，可能会以零开头：

例如，"10:2" 是无效的时间，正确的写法应该是 "10:02" 。

输入：turnedOn = 1
输出：["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]

输入：turnedOn = 9
输出：[]

0 <= turnedOn <= 10

思路:
花里胡哨的题目，看第一眼一脸懵逼，仔细想一想其实是个组合问题，
手表上面4个灯（小时）下面6个灯（分钟）共10个灯， 有n个灯亮 可以转换成 在 1~10 数组中挑选n个数的组合有哪些，
然后把这些组合翻译成时间，不符合时间规则的组合去掉

*/
func readBinaryWatch(turnedOn int) []string {
	var numsToTimeFn func([]int) string // 将数字转成时间字符串
	numsToTimeFn = func(nums []int) string {
		// nums 中的数字只能由 0~9组成,最大长度是10
		// 数字0~3表示小时, 数字4~9表示分钟
		hour, minute := 0, 0
		for _, v := range nums {
			if 0 <= v && v <= 3 {
				lsh := uint(v)
				hour += 1 << lsh
			} else if 4 <= v && v <= 9 {
				lsh := uint(v - 4)
				minute += 1 << lsh
			}
		}
		// 检测时间是否合规
		if (hour < 0 || hour > 11) || (minute < 0 || minute > 59) {
			return ""
		}
		return fmt.Sprintf("%d:%02d", hour, minute)
	}
	// 此处逻辑和 77 题思路一样
	var res []string
	var dfsFn func(int)
	var que []int
	dfsFn = func(start int) {
		if len(que) == turnedOn {
			if timeStr := numsToTimeFn(que); timeStr != "" {
				res = append(res, timeStr)
			}
			return
		}
		end := 9 - (turnedOn - len(que)) + 1
		for i := start; i <= end; i++ {
			que = append(que, i)
			dfsFn(i + 1)
			que = que[:len(que)-1]
		}
	}
	dfsFn(0)
	return res
}

func Test_readBinaryWatch(t *testing.T) {
	res := readBinaryWatch(2)
	t.Logf("%v\n", res)
}
