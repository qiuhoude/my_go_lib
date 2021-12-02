package leetcode

// 219. 存在重复元素 II https://leetcode-cn.com/problems/contains-duplicate-ii/

/*
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。

输入: nums = [1,2,3,1], k = 3
输出: true

输入: nums = [1,0,1,1], k = 1
输出: true

输入: nums = [1,2,3,1,2,3], k = 2
输出: false

思路1 :
使用查表法,key为数组中的数字,val存储最近一次出现该数字的下标,
在遍历数组的过程中发现map中有值并且最近一次出现的下标与当前下标的差值 <=k 直接返回

思路2 :
使用滑动窗口+查找表, 窗口的大小  width=k
l,r=l+k 是窗口头尾指针,每次在 [l,r]之间使用hash表存储,发现重复就返回
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	n := len(nums)
	if n < 2 || k < 1 {
		return false
	}
	tab := make(map[int]bool)

	for r := 0; r < n; r++ {
		if _, ok := tab[nums[r]]; ok {
			return true
		}
		l := r - k
		if l >= 0 {
			delete(tab, nums[l])
		}
		tab[nums[r]] = true
	}
	return false
}

func containsNearbyDuplicate1(nums []int, k int) bool {
	if len(nums) < 2 || k < 1 {
		return false
	}
	tab := make(map[int]int)
	for i, v := range nums {
		if lastI, ok := tab[v]; ok && i-lastI <= k {
			return true
		}
		tab[v] = i
	}
	return false
}
