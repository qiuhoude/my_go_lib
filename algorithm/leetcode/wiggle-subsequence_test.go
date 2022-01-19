package leetcode

import "testing"

// 376. 摆动序列 https://leetcode-cn.com/problems/wiggle-subsequence/

/*
如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列 。第一个差（如果存在的话）可能是正数或负数
。仅有一个元素或者含两个不等元素的序列也视作摆动序列。

例如， [1, 7, 4, 9, 2, 5] 是一个 摆动序列 ，因为差值 (6, -3, 5, -7, 3) 是正负交替出现的。
相反，[1, 4, 7, 2, 5] 和 [1, 7, 4, 5, 5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。

子序列 可以通过从原始序列中删除一些（也可以不删除）元素来获得，剩下的元素保持其原始顺序。
给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列的长度 。

输入：nums = [1,7,4,9,2,5]
输出：6
解释：整个序列均为摆动序列，各元素之间的差值为 (6, -3, 5, -7, 3) 。

输入：nums = [1,17,5,10,13,15,10,5,16,8]
输出：7
解释：这个序列包含几个长度为 7 摆动序列。
其中一个是 [1, 17, 10, 13, 10, 16, 8] ，各元素之间的差值为 (16, -7, 3, -3, 6, -8) 。

输入：nums = [1,2,3,4,5,6,7,8,9]
输出：2

1 <= nums.length <= 1000
0 <= nums[i] <= 1000

进阶：你能否用 O(n) 时间复杂度完成此题?

思路:

1.动态规划
该题和 300. 最长上升子序列 基本一致

up[i] = up[i-1]				  	// nums[i] <= nums[i-1]
up[i] = max(up[i-1],down[i-1]+1) // nums[i] > nums[i-1]
down[i] = up[i-1]				  // nums[i] >= nums[i-1]
down[i] = max(up[i-1]+1,down[i-1]) // nums[i] < nums[i-1]
max(up[n-],down[n-1])

2.计数思路:
其实就是求波峰波谷的个数 计数就可以
*/
// 动态规划思路
func wiggleMaxLengthDp(nums []int) int {
	maxFn := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(nums)
	up, down := make([]int, n), make([]int, n)
	up[0], down[0] = 1, 1

	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up[i] = maxFn(up[i-1], down[i-1]+1)
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			down[i] = maxFn(up[i-1]+1, down[i-1])
			up[i] = up[i-1]
		} else {
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}

	return maxFn(up[n-1], down[n-1])

}

// 计数思路
func wiggleMaxLength(nums []int) int {
	n := len(nums)

	if n == 1 {
		return 1
	}
	const (
		initState = 0 //未初始化不不到下一个是波峰还是波谷
		up        = 1
		down      = 2
	)

	res := 1
	state := initState // 0为位  1为上升 2为下降
	for i := 1; i < n; i++ {
		if state == initState { // 没有被初始化
			// 进行初始化
			if nums[0] == nums[i] {
				continue
			} else if nums[0] > nums[i] { // 下降
				state = up //下一波要上升
				res++
			} else { //上升
				state = down
				res++
			}
		} else {
			if nums[i] == nums[i-1] {
				continue
			} else if nums[i-1] > nums[i] { // 实际下降
				if state == down { //期望也是下降
					state = up
					res++
				}
			} else { // 实际上升
				if state == up { // 期望也是下降
					state = down
					res++
				}
			}
		}
	}
	return res
}

func Test_wiggleMaxLength(t *testing.T) {

	tests := []struct {
		arg  []int
		want int
	}{
		{[]int{1, 7, 4, 9, 2, 5}, 6},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},
		{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, 7},
	}
	for _, tt := range tests {
		if got := wiggleMaxLength(tt.arg); got != tt.want {
			t.Errorf("wiggleMaxLength(%v) => got=%v  want=%v", tt.arg, got, tt.want)
		}
	}
}
