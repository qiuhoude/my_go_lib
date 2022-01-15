package leetcode

import "testing"

// 416. 分割等和子集 https://leetcode-cn.com/problems/partition-equal-subset-sum/

/*
给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。

输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。


1 <= nums.length <= 200
1 <= nums[i] <= 100

思路:
一开始拿到题有点懵, 因为是将分成两个子集并且两子集的合还相等, 子集的合必然是 sum/2,
可以将问题转换成背包问题中 在nums中找出一些数的总和 等于sum/2
half = sum/2 ; sum是奇数无解

自顶向下思路
定义函数 f(i,s) i:数组下标初始值是0, s:最初值是half
f(i,s) => f(i+1,s-nums[i]) // 选择 i
		  || f(i+1,s)	  // 不选择 i

动态规划思路:
开辟 二维数组 [len(nums)][halt]bool, 记录 每个物体选择还是不选择的进行标记
优化使用单数组
*/
func canPartition3(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 { // 奇数
		return false
	}
	half := sum / 2
	// 单数组
	dp := make([]bool, half+1)
	dp[0] = true
	if nums[0] <= half {
		dp[nums[0]] = true
	}
	n := len(nums)
	for i := 1; i < n; i++ {
		for s := half - nums[i]; s >= 0; s-- { // 需要倒过来,避免前面的加上之后污染后面的值
			if dp[s] { //选择i
				dp[s+nums[i]] = true
			}
		}
	}
	return dp[half]
}

func canPartition2(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 { // 奇数
		return false
	}
	half := sum / 2
	n := len(nums)
	dp := make([][]bool, len(nums))
	for i := range dp {
		dp[i] = make([]bool, half+1)
	}
	// 第一排
	dp[0][0] = true
	if nums[0] <= half {
		dp[0][nums[0]] = true
	}
	for i := 1; i < n; i++ {
		// 不选择i
		for s := 0; s <= half; s++ {
			if dp[i-1][s] {
				dp[i][s] = true
				if s == half { // 提前返回
					return true
				}
			}
		}
		// 选择i
		for s := 0; s <= half-nums[i]; s++ {
			if dp[i-1][s] {
				dp[i][s+nums[i]] = true
			}
			if s == half { // 提前返回
				return true
			}
		}

	}
	for i := 0; i < n; i++ {
		if dp[i][half] {
			return true
		}
	}
	return false
}

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 { // 奇数
		return false
	}
	memory := make(map[[2]int]bool)

	n := len(nums)
	var fn func(i, s int) bool
	fn = func(i, s int) bool {
		key := [2]int{i, s}
		if v, ok := memory[key]; ok {
			return v
		}
		if i > n-1 || s < 0 {
			return false
		}
		if s == 0 {
			return true
		}
		r1 := fn(i+1, s-nums[i]) // 选择 i
		r2 := fn(i+1, s)         // 不选择 i
		r := r1 || r2
		memory[key] = r
		return r
	}
	return fn(0, sum/2)
}

func Test_canPartition(t *testing.T) {
	tests := []struct {
		arg  []int
		want bool
	}{
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 2, 3, 5}, false},
		{[]int{99, 2, 3, 98}, true},
		{[]int{99, 2, 3, 2}, false},
	}
	for _, tt := range tests {
		if got := canPartition3(tt.arg); got != tt.want {
			t.Errorf("canPartition(%v) => got=%v  want=%v", tt.arg, got, tt.want)
		}
	}
}
