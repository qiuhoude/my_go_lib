package leetcode

// 454. 四数相加 II https://leetcode-cn.com/problems/4sum-ii/

/*

给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：

0 <= i, j, k, l < n
nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0

输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
输出：2
解释：
两个元组如下：
1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0

输入：nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
输出：1

n == nums1.length
n == nums2.length
n == nums3.length
n == nums4.length
1 <= n <= 200
-228 <= nums1[i], nums2[i], nums3[i], nums4[i] <= 228

思路:
使用查找表，将其中两个数组(nums3,nums4)的所有情况的和放到一个map中, 来减少遍历的次数，可以将暴力解法的 O(n^4) 降低到 O(n^2)

*/

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	tabSum := make(map[int]int) // <sum,count>
	res := 0
	// nums3,nums4的合放到map中
	for i := range nums3 {
		for j := range nums4 {
			tabSum[nums3[i]+nums4[j]]++
		}
	}
	for i := range nums1 {
		for j := range nums2 {
			target := 0 - (nums1[i] + nums2[j])
			if cnt, ok := tabSum[target]; ok {
				res += cnt
			}
		}
	}
	return res
}
