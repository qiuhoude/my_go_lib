package leetcode

import (
	"sort"
	"testing"
)

// 435. 无重叠区间 https://leetcode-cn.com/problems/non-overlapping-intervals/
/*
给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。

注意:
可以认为区间的终点总是大于它的起点。
区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。

输入: [ [1,2], [2,3], [3,4], [1,3] ]
输出: 1
解释: 移除 [1,3] 后，剩下的区间没有重叠。

输入: [ [1,2], [1,2], [1,2] ]
输出: 2
解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。

输入: [ [1,2], [2,3] ]
输出: 0
解释: 你不需要移除任何区间，因为它们已经是无重叠的了。

思路:
题目转换成求最多有多少个不充分的区间

动态规划 或 贪心 结合

动态规划:
对区间开始或结尾进行排序
有点类似于 300.最长上升子序列 的问题, 例如当前比较到 dp[i] = for j in 0~(i-1) max(dp[j]+1,dp[i]), dp存储不重叠个数
然后 max(dp)

贪心:
对区间结尾进行排序, 每次选着结尾越早并且不会与选择的上一个区间重叠 (结尾越早说明留给后面空间就越大)

*/

// 动态规划 (居然超时) ,这是逼着用贪心
func eraseOverlapIntervalsDP(intervals [][]int) int {
	maxFn := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool { // 开始位置升序
		return intervals[i][0] < intervals[j][0]
	})
	dp := make([]int, n) //dp[i] 到i位置的最多不重复区间个数
	dp[0] = 1            // 最少有1个
	res := dp[0]         // 最多有多少个不重复的区间
	for i := 1; i < n; i++ {
		dp[i] = 1 // 最少有1个
		for j := 0; j < i; j++ {
			if intervals[i][0] >= intervals[j][1] { // 两区间不覆盖
				dp[i] = maxFn(dp[j]+1, dp[i])
			}
		}
		res = maxFn(dp[i], res)
	}
	return n - res
}

// 贪心
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}

	// 求最多有多少个子区域不重叠 ,然后 len(intervals) - 最多个数
	sort.Slice(intervals, func(i, j int) bool { // 结束位置升序,每次选择最少结束的区间,给后面留下更大的空间
		return intervals[i][1] < intervals[j][1]
	})

	res := 1      // 只有一个
	preIndex := 0 // 选择的前一个区间的下标
	for i := 1; i < n; i++ {
		if intervals[i][0] >= intervals[preIndex][1] { // 有覆盖
			res++
			preIndex = i
		}
	}
	return n - res
}

// 判断两个区间是否有重复
func isOverlapping(rg1, rg2 []int) bool {
	return (rg1[0] < rg2[0] && rg2[0] < rg1[1]) ||
		(rg1[0] < rg2[1] && rg2[1] < rg1[1]) ||
		(rg2[0] < rg1[0] && rg1[0] < rg2[1]) ||
		(rg2[0] < rg1[1] && rg1[1] < rg2[1]) ||
		rg1[0] == rg2[0] && rg1[1] == rg2[1]
}

func Test_eraseOverlapIntervals(t *testing.T) {
	//arg := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	arg := [][]int{{1, 2}}
	ret := eraseOverlapIntervalsDP(arg)
	t.Log(ret)
}

func Test_isOverlapping(t *testing.T) {
	tests := []struct {
		name string
		arg1 []int
		arg2 []int
		want bool
	}{
		{"[1,2],[1,3]", []int{1, 2}, []int{1, 3}, true},
		{"[1,2],[3,4]", []int{1, 2}, []int{3, 4}, false},
		{"[4,8],[3,4]", []int{4, 8}, []int{3, 4}, false},
		{"[4,8],[3,6]", []int{4, 8}, []int{3, 6}, true},
		{"[1,2],[1,2]", []int{1, 2}, []int{1, 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOverlapping(tt.arg1, tt.arg2); got != tt.want {
				t.Errorf("isOverlapping() = %v, want %v", got, tt.want)
			}
		})
	}
}
