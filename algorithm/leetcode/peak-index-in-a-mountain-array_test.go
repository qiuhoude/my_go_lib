package leetcode

// 852. 山脉数组的峰顶索引 https://leetcode.cn/problems/peak-index-in-a-mountain-array/
/*

符合下列属性的数组 arr 称为 山脉数组 ：
arr.length >= 3
存在 i（0 < i < arr.length - 1）使得：
arr[0] < arr[1] < ... arr[i-1] < arr[i]
arr[i] > arr[i+1] > ... > arr[arr.length - 1]
给你由整数组成的山脉数组 arr ，返回任何满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i 。



示例 1：

输入：arr = [0,1,0]
输出：1
示例 2：

输入：arr = [0,2,1,0]
输出：1
示例 3：

输入：arr = [0,10,5,2]
输出：1
示例 4：

输入：arr = [3,4,5,1]
输出：2
示例 5：

输入：arr = [24,69,100,99,79,78,67,36,26,19]
输出：2


提示：
3 <= arr.length <= 10^4
0 <= arr[i] <= 10^6
题目数据保证 arr 是一个山脉数组


进阶：很容易想到时间复杂度 O(n) 的解决方案，你可以设计一个 O(log(n)) 的解决方案吗？

思路: 使用二分查找的方式
left=1, right=len(arr)-2, 在left~right 之间查找值
如何该值 左，自己，右 单调增 则 left=mid+1, 反之单调减 right=mid-1
正好是 自己最大就是要找到的值

*/

func peakIndexInMountainArray(arr []int) int {
	left, right := 1, len(arr)-2
	for left <= right {
		mid := left + ((right - left) >> 1)
		if arr[mid] > arr[mid+1] && arr[mid] > arr[mid-1] { // 找到了
			return mid
		} else if arr[mid-1] < arr[mid] && arr[mid] < arr[mid+1] { // 爬山的阶段
			left = mid + 1
		} else if arr[mid-1] > arr[mid] && arr[mid] > arr[mid+1] {
			right = mid - 1
		} else {
			return -1 // 数据有误
		}
	}
	return -1
}
