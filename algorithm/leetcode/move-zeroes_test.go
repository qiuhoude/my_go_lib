package leetcode

import (
	"testing"
)

// 283. 移动零 https://leetcode-cn.com/problems/move-zeroes/

/*
26,  80
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]

思路1: 可以借鉴冒泡排序,遇到0就与后面的非0元素进行交换,
先统计0元素的总数 zeroN
每次循环都会将一个0元素转移到对应的位置, 循环的区间:[0,len-i] ,i表示第几次循环

思路2(优化): 使用双指针,指针i记录当前0元素的位置([0...i)之间是非0元素),指针j用于表示整个遍历进度,
j只要发现非0元素就和i进行替换,然后i++
遍历完成后 [i,len)区间的值赋值成0

*/
func moveZeroes2(nums []int) {
	l := len(nums)
	i := 0 // [0,i) 是非0元素
	for j := 0; j < l; j++ {
		if nums[j] != 0 { // 非0替换
			if i != j {
				nums[j], nums[i] = nums[i], nums[j]
			}
			i++
		}
	}
	//for k := i; k < l; k++ {
	//	nums[k] = 0
	//}

}
func moveZeroes(nums []int) {
	// 统计0元素的个数
	zeroN := 0 // 0 元素的个数
	for _, e := range nums {
		if e == 0 {
			zeroN++
		}
	}
	//if zeroN == 0 { // 没有0元素直接返回
	//	return
	//}
	l := len(nums)
	for i := 0; i < l && zeroN > 0; i++ {
		for j := 0; j < l-1-i; j++ { // (l-1)-i 避免j溢出
			if nums[j] == 0 && nums[j+1] != 0 {
				nums[j], nums[j+1] = nums[j+1], nums[j] // swap
			}
		}
		zeroN--
	}
}

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{2, 1}},
	}
	for _, tt := range tests {
		//moveZeroes(tt.input)
		moveZeroes2(tt.input)
		if !eqSliceInt(tt.input, tt.expected) {
			t.Errorf("expected=%v, got=%v", tt.expected, tt.input)
		}
	}

}
func eqSliceInt(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
