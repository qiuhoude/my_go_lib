package leetcode

// 215. 数组中的第K个最大元素 https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

/*

给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4

思路1: 使用快排 partition 选取 pivot 的特性, 找到pivot对应位置,
查看pivot位置左右元素的个数,来确定往左找还是又右边找

思路2:
构建大堆,弹k个元素出来

*/

func findKthLargest(nums []int, k int) int {
	l := len(nums)
	if l < 2 {
		return nums[0]
	}
	target := l - k
	i, p, r := 0, 0, l-1
	for {
		i = partition(nums, p, r)
		if i == target {
			break
		} else if i > target { // 在目标的右边
			r = i - 1
		} else { // < target
			p = i + 1
		}
	}
	return nums[i]
}

// 对切片分区,返回 pivot的index
func partition(nums []int, p, r int) int {
	pivot := nums[r]
	l := p // 左边部分
	for j := p; j < r; j++ {
		if nums[j] < pivot {
			if l != j {
				nums[l], nums[j] = nums[j], nums[l]
			}
			l++
		}
	}
	nums[l], nums[r] = nums[r], nums[l]
	return l
}
