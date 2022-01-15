package leetcode

import "testing"

// 377. 组合总和 Ⅳ https://leetcode-cn.com/problems/combination-sum-iv/

/*
给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
题目数据保证答案符合 32 位整数范围。

输入：nums = [1,2,3], target = 4
输出：7
解释：
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
请注意，顺序不同的序列被视作不同的组合。


输入：nums = [9], target = 3
输出：0

1 <= nums.length <= 200
1 <= nums[i] <= 1000
nums 中的所有元素 互不相同
1 <= target <= 1000

进阶：如果给定的数组中含有负数会发生什么？问题会产生何种变化？如果允许负数出现，需要向题目中添加哪些限制条件？

思路:
又是一个可以转换动态规划背包问题的题目,和 322. 硬币问题 条件基本一致,只是求的内容不同

自顶向下的思路
	定义函数 f(remain)  remain:剩余数量
    sum(for(f(amount-coins[i])))  //  只要有解,就加上返回值

动态规划的思路
dp[i-nums[j]]合法 ,进行 sum += dp[i-nums[i]] 就可以

*/
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1 // 不选任元素只有1种选择
	for i := 1; i <= target; i++ {
		sum := 0
		for j := range nums {
			if i-nums[j] >= 0 {
				sum += dp[i-nums[j]]
			}
		}
		dp[i] = sum
	}
	return dp[target]

}

// 超时...
func combinationSum4r(nums []int, target int) int {
	memory := make([]int, target+1)
	/*
		定义函数 f(remain)  remain:剩余数量
	    sum(for(f(amount-coins[i])))  //  只要有解,就加上返回值
	*/
	var fn func(remain int) int
	fn = func(remain int) int {
		if remain < 0 { // 没有解直接返回0
			return 0
		}
		if remain == 0 { // 有解
			return 1
		}
		if memory[remain-1] != 0 {
			return memory[remain-1]
		}
		sum := 0
		for i := range nums {
			r := fn(remain - nums[i])
			// 回溯
			sum += r
		}
		memory[remain-1] = sum
		return sum
	}
	return fn(target)
}

func Test_combinationSum4(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 2, 3}, 4, 7},
		{[]int{9}, 3, 0},
		{[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 210, 220, 230, 240, 250, 260, 270, 280, 290, 300, 310, 320, 330, 340, 350, 360, 370, 380, 390, 400, 410, 420, 430, 440, 450, 460, 470, 480, 490, 500, 510, 520, 530, 540, 550, 560, 570, 580, 590, 600, 610, 620, 630, 640, 650, 660, 670, 680, 690, 700, 710, 720, 730, 740, 750, 760, 770, 780, 790, 800, 810, 820, 830, 840, 850, 860, 870, 880, 890, 900, 910, 920, 930, 940, 950, 960, 970, 980, 990, 111},
			999, 1},
	}
	for _, tt := range tests {
		if got := combinationSum4(tt.nums, tt.target); got != tt.want {
			t.Errorf("combinationSum4( %v) = %v, want %v", tt.target, got, tt.want)
		}

	}
}
