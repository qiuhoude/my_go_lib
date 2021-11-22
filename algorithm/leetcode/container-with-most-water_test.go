package leetcode

// 11. 盛最多水的容器 https://leetcode-cn.com/problems/container-with-most-water/

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
(i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

输入：height = [1,1]
输出：1

输入：height = [1,2,1]
输出：2

思路:
求总容器最大,可以先令容器宽度最大,然后慢慢缩小宽度找出最大值,
使用对撞指针, 谁小就移动随

*/

func maxArea(height []int) int {
	h, t := 0, len(height)-1
	curMax := 0
	for h < t {
		tmp := (t - h) * minI(height[h], height[t])
		if tmp > curMax {
			curMax = tmp
		}
		if height[h] > height[t] {
			t--
		} else {
			h++
		}
	}
	return curMax
}

func minI(a, b int) int {
	if a > b {
		return b
	}
	return a
}
