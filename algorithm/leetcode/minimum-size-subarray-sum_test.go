package leetcode

import (
	"testing"
)

//209. 长度最小的子数组 https://leetcode-cn.com/problems/minimum-size-subarray-sum/

/*
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，
并返回其长度。如果不存在符合条件的子数组，返回 0 。

输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长

输入：target = 4, nums = [1,4,4]
输出：1

输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0

1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 105

思路:
滑动窗口, ph,pt 两个指针
保证窗口满足条件,
扩大窗口 小于target或pt==ph就pt++进
缩小窗口 大于target就ph++
*/

func minSubArrayLen(target int, nums []int) int {
	l := len(nums)
	ph, pt := 0, 0   // 滑动窗口的指针
	wSum := nums[ph] // 窗口的数的总和
	minSeqLen := 0   // 最小子序列长度
	for ph <= pt && pt < l {
		if wSum >= target { // 判断窗口条件满足
			tmpLen := pt - ph + 1
			if minSeqLen == 0 || minSeqLen > tmpLen {
				minSeqLen = tmpLen
			}
			if minSeqLen == 1 { // 直接找到最小的序列,不用继续了
				return minSeqLen
			}
		}

		if pt < l-1 && (pt == ph || wSum <= target) { //  窗口扩大
			pt++
			wSum += nums[pt]
		} else { // 窗口缩小
			wSum -= nums[ph]
			ph++
		}
	}
	return minSeqLen
}

func Test_minSubArrayLen(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{"[2, 3, 1, 2, 4, 3], 7", []int{2, 3, 1, 2, 4, 3}, 7, 2},
		{"[1, 1, 1, 1, 1, 1, 1, 1], 11", []int{1, 1, 1, 1, 1, 1, 1, 1}, 11, 0},
		{"[1, 4, 4], 4", []int{1, 4, 4}, 4, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.target, tt.nums); got != tt.expected {
				t.Errorf("minSubArrayLen() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
