package leetcode

import "testing"

// 42. 接雨水 https://leetcode-cn.com/problems/trapping-rain-water/

/*
给定n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：
输入：height = [4,2,0,3,2,5]
输出：9

n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105

思路：
与 11题.盛最多水的容器 思路有些类似，使用对撞双指针的思路,
默认开始 底为 base=0,  h = min(height[i],height[j])
在 i,j之间找 height[n]<h 进行累加， 迭代一次记录最大的curBase然后复制给base.
移动值小的一端的指针

这种方式很慢之后优化吧
1047题待做
*/

func trap(height []int) int {
	n := len(height)
	if n <= 0 {
		return 0
	}
	i, j := 0, n-1
	base := 0  // 基础水平面
	water := 0 // 水量

	var minFn = func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	var maxFn = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i < j {
		if height[i] <= base { // 低于底部
			i++
			continue
		}
		if height[j] <= base {
			j--
			continue
		}
		// 走到此处说明 两边柱子比基础水平面高，才能可能盛水
		curH := minFn(height[i], height[j]) // 当前的柱子
		curSumW := 0
		for k := i + 1; k < j; k++ {
			if height[k] < curH { // 可以盛水
				w := curH - maxFn(base, height[k])
				curSumW += w
			}
		}
		water += curSumW // 加上当前轮的水量
		base = curH
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return water
}

func Test_trap(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
	}
	for _, tt := range tests {
		if got := trap(tt.height); got != tt.want {
			t.Errorf("trap( %v) got %v, want %v", tt.height, got, tt.want)
		}
	}

}
