package leetcode

import "testing"

// 300. 最长上升子序列 https://leetcode-cn.com/problems/longest-increasing-subsequence/

/*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

输入：nums = [0,1,0,3,2,3]
输出：4

输入：nums = [7,7,7,7,7,7,7]
输出：1

1 <= nums.length <= 2500
-104 <= nums[i] <= 104

思路:

动态规划思路
定义函数 f(i)  i表示数值的下标, 返回最长子序列,第i个数的最长子序列
f(i) = for (j in i之前的 max(f(j)+1,f(i))

*/

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	maxFn := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ { //比较前面i项最大子序列值
			if nums[j] < nums[i] {
				dp[i] = maxFn(dp[j]+1, dp[i])
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func lengthOfLIS2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// tail 数组的定义：长度为 i + 1 的上升子序列的末尾最小是几
	tail := make([]int, len(nums))
	// 遍历第 1 个数，直接放在有序数组 tail 的开头
	tail[0] = nums[0]
	// end 表示有序数组 tail 的最后一个已经赋值元素的索引
	end := 0
	for i := 1; i < len(nums); i++ {
		// 【逻辑 1】比 tail 数组实际有效的末尾的那个元素还大
		if nums[i] > tail[end] {
			end++
			tail[end] = nums[i]
		} else {
			// 使用二分查找法，在有序数组 tail 中
			// 找到第 1 个大于等于 nums[i] 的元素，尝试让那个元素更小
			left := 0
			right := left
			for left < right {
				// 选左中位数不是偶然，而是有原因的，原因请见 LeetCode 第 35 题题解
				mid := left + (right-left)/2
				if tail[mid] < nums[i] {
					// 中位数肯定不是要找的数，把它写在分支的前面
					left = mid + 1
				} else {
					right = mid
				}
			}
			//走到这里是因为 【逻辑 1】 的反面，因此一定能找到第 1 个大于等于 nums[i] 的元素
			tail[left] = nums[i]
		}
	}
	end++
	return end
}

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		//{"[10,9,2,5,3,7,101,18] => 4", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		//{"[10,9,2,5,3,4] => 3", []int{10, 9, 2, 5, 3, 4}, 3},
		{"[1,3,6,7,9,4,10,5,6] => 6", []int{1, 3, 6, 7, 9, 4, 10, 5, 6}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)

			}
		})
	}
}
