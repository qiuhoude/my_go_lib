package leetcode

// 41. 缺失的第一个正数 https://leetcode-cn.com/problems/first-missing-positive/

/*
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

输入：nums = [1,2,0]
输出：3

输入：nums = [3,4,-1,1]
输出：2

输入：nums = [7,8,9,11,12]
输出：1

1 <= nums.length <= 5 * 10^5
-2^31 <= nums[i] <= 2^31 - 1

思路:
看别人详解后做出
总体思路使用hash表, 但是hash表的的空间复杂度不满足题意,
退而求其次,可以将数组下标当作hash表使用, 题目是要求得是最小的正整数值,可以将原数组中[1,len]范围内得值,映射成数组下标,进行替换操作,不在范围类得就不管
映射公式: val = index+1
然后下标从小到大遍历数组,
如果有一个映射值不等于下标值此值就是未出现得得最小值
如果数组中全部映射了说明未出现得最小值是 len+1

*/
func firstMissingPositive(nums []int) int {
	n := len(nums)
	// 将[1,len]进行替换
	for i := 0; i < n; i++ {
		// 确保i位置放着得满足条件得, 要用for不能使用if
		for nums[i] >= 1 && nums[i] <= n && nums[i] != nums[nums[i]-1] { //在[1,len]范围
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return n + 1
}

//func Test_firstMissingPositive(t *testing.T) {
//	res := firstMissingPositive([]int{3, 4, -1, 1})
//	t.Logf("%v\n", res)
//}
