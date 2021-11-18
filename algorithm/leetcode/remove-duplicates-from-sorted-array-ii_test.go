package leetcode

import (
	"testing"
)

// 80. 删除有序数组中的重复项 II https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/

/*

给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 最多出现两次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

输入：nums = [1,1,1,2,2,3]
输出：5, nums = [1,1,2,2,3]
解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。 不需要考虑数组中超出新长度后面的元素。

输入：nums = [0,0,1,1,1,1,2,3,3]
输出：7, nums = [0,0,1,1,2,3,3]
解释：函数应返回新长度 length = 7, 并且原数组的前五个元素被修改为 0, 0, 1, 1, 2, 3, 3 。 不需要考虑数组中超出新长度后面的元素。

1 <= nums.length <= 3 * 104
-104 <= nums[i] <= 104
nums 已按升序排列

思路1: 和26 思路一样都是分区[0,k)满足条件区域, 只是判断重复时加上了条件
思路2: 双指针快慢指针思想,快慢指针间隔最少间隔2,判断
*/

func removeDuplicatesii(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}

	k := 1          // [0,k)区域是非重复元素
	curR := nums[0] // 当前重复的元素的值
	rCnt := 0       // 记录元素重复的次数
	for i := 1; i < l; i++ {
		if nums[i] != curR { // 非重复元素
			curR = nums[i]
			nums[k] = curR
			k++
			rCnt = 0 // 重复次数清除
		} else { // 重复
			rCnt++
			if rCnt < 2 {
				nums[k] = curR // 注意 k++ 和 nums[k] = curR 操作是绑定再一起的,k++意义是满足重复的条件
				k++
			}
		}
	}
	return k
}

func Test_removeDuplicatesii(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	expected := []int{0, 0, 1, 1, 2, 3, 3}
	n := removeDuplicatesii2(nums)
	got := nums[:n]
	if !eqSliceInt(expected, got) {
		t.Errorf("expected=%v, got=%v", expected, got)
	}
}

func removeDuplicatesii2(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}
	fast, slow := 2, 2
	for fast < l {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
