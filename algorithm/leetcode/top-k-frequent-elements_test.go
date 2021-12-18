package leetcode

// 347. 前 K 个高频元素 https://leetcode-cn.com/problems/top-k-frequent-elements/

/*
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]

输入: nums = [1], k = 1
输出: [1]

1 <= nums.length <= 105
k 的取值范围是 [1, 数组中不相同的元素的个数]
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的

进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。

思路1:
堆排序,使用优先级队列进行处理

思路2:
使用桶排序
*/

func topKFrequent(nums []int, k int) []int {
	freqTab := map[int]int{}
	maxFreq := 0
	for _, v := range nums {
		freqTab[v]++
		if freqTab[v] > maxFreq {
			maxFreq = freqTab[v]
		}
	}
	// 建立桶
	bucket := make([][]int, maxFreq+1)
	for ch, cnt := range freqTab {
		bucket[cnt] = append(bucket[cnt], ch)
	}
	var res []int
	cnt := k
out:
	for i := len(bucket) - 1; i >= 0; i-- {
		arr := bucket[i]
		if len(arr) > 0 {
			for _, n := range arr {
				res = append(res, n)
				cnt--
				if cnt == 0 {
					break out
				}
			}
		}
	}
	return res
}
