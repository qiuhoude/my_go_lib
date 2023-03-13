package leetcode

import (
	"github.com/bmizerany/assert"
	"testing"
)

// 410. 分割数组的最大值
/*
给定一个非负整数数组 nums 和一个整数m ，你需要将这个数组分成m个非空的连续子数组。
设计一个算法使得这m个子数组各自和的最大值最小。

示例 1：

输入：nums = [7,2,5,10,8], m = 2
输出：18
解释：
一共有四种方法将 nums 分割为 2 个子数组。
其中最好的方式是将其分为 [7,2,5] 和 [10,8] 。
因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。

示例 2：

输入：nums = [1,2,3,4,5], m = 2
输出：9
示例 3：

输入：nums = [1,4,4], m = 3
输出：4

提示：
1 <= nums.length <= 1000
0 <= nums[i] <= 106
1 <= m <= min(50, nums.length)


思路：
和 1062题（vip题） 思路有些类似, 将问题转换成 二分搜索 + 条件函数的方式
1. 使用二分搜索 + 条件函数
本地的答案在  max(nums[]) 到 sum(nums[]) 之间进查找,
条件函数就是, 遍历数组 将数值进行相加 ,当sum大于target 就将切分份数 cnt++, 最后比较参数的份数 和 实际份数  cnt<=paramCnt
*/

func splitArray(nums []int, k int) int {
	left, right := 0, 0
	for _, num := range nums {
		right += num
		if num > left {
			left = num
		}
	}
	for left < right {
		mid := left + ((right - left + 1) >> 1)
		if checkSplitArray(nums, mid, k) { // 说明满足条件，但要求最小值，需要继续从右往左找, 如果最大值则是从从左往右找
			right = mid
		} else {
			left = mid + 1
		}

	}
	return right
}

// target: 表示目标值, k 分割的数量, true表示满足分割的条件
func checkSplitArray(nums []int, target int, k int) bool {
	tmpSum, cnt := 0, 1
	for _, num := range nums {
		if tmpSum+num > target {
			cnt++
			tmpSum = num // 清除重新计算
		} else {
			tmpSum += num
		}
	}
	return cnt <= k
}

func Test_splitArray(t *testing.T) {
	ret := splitArray([]int{7, 2, 5, 10, 8}, 2)
	assert.Equal(t, 18, ret)
}
