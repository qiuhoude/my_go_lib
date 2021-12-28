package leetcode

import (
	"testing"
)

// 46. 全排列 https://leetcode-cn.com/problems/permutations/
/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

输入：nums = [0,1]
输出：[[0,1],[1,0]]

输入：nums = [1]
输出：[[1]]

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同

思路:
使用递归回溯法, 想象成一个树,每个数字都可以当作根节点进行找路径
*/

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	if len(nums) == 2 {
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	}

	var result [][]int
	for index, value := range nums {
		var numsCopy = make([]int, len(nums))
		copy(numsCopy, nums)
		// 将numsCopy中index这个元素给剔除掉赋值给numsSubOne
		numsSubOne := append(numsCopy[:index], numsCopy[index+1:]...)
		valueSlice := []int{value}
		newSubSlice := permute(numsSubOne)
		for _, newValue := range newSubSlice {
			result = append(result, append(valueSlice, newValue...))
		}
	}
	return result
}

// 利用元素在本地dfs的进行替换
func permute2(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	var permutation func(start int)
	// dfs 进行替换
	permutation = func(start int) {
		if start == n-1 { // 最后一位
			res = append(res, append([]int(nil), nums...))
			return
		}
		for i := start; i < n; i++ {
			nums[i], nums[start] = nums[start], nums[i] // 与开始	位置进行交换
			permutation(start + 1)
			nums[i], nums[start] = nums[start], nums[i] // 换回来
		}
	}
	permutation(0)
	return res
}

func Test_permute(t *testing.T) {
	nums := []int{1, 2, 3}
	res := permute2(nums)
	t.Logf("%v\n", res)
}
