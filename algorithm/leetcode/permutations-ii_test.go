package leetcode

import (
	"sort"
	"testing"
)

// 47. 全排列 II https://leetcode-cn.com/problems/permutations-ii/

/*
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。


输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

1 <= nums.length <= 8
-10 <= nums[i] <= 10

思路:
与
*/
func permuteUnique(nums []int) [][]int {
	var res [][]int
	// 先排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	used := make([]bool, len(nums)) // 下标是否被使用过
	var stack []int
	var permuteHelper func(start int)
	permuteHelper = func(start int) {
		n := len(nums)
		if start == n {
			// 出口已经找到
			res = append(res, append([]int(nil), stack...))
			return
		}
		for i := 0; i < n; i++ {
			if !used[i] {
				if i > 0 && nums[i] == nums[i-1] && used[i-1] { // 去重, 本元素与上个元素相等跳过,并且没有使用过
					continue
				}
				used[i] = true // 已经访问
				stack = append(stack, nums[i])
				permuteHelper(start + 1)
				stack = stack[:len(stack)-1] // 恢复原来的路径
				used[i] = false
			}
		}
	}
	permuteHelper(0)
	return res
}

func permuteUnique2(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	if len(nums) == 2 {
		if nums[0] == nums[1] { // 与前一个数相等 只返回一种可能性
			return [][]int{{nums[0], nums[1]}}
		} else {
			return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
		}
	}
	// 先排序
	sort.Ints(nums)
	var result [][]int
	for index, value := range nums {
		var numsCopy = make([]int, len(nums))
		copy(numsCopy, nums)
		if index > 0 && numsCopy[index] == numsCopy[index-1] {
			continue
		}
		// 将numsCopy中index这个元素给剔除掉赋值给numsSubOne
		numsSubOne := append(numsCopy[:index], numsCopy[index+1:]...)
		valueSlice := []int{value}
		newSubSlice := permuteUnique2(numsSubOne)
		for _, newValue := range newSubSlice {
			result = append(result, append(valueSlice, newValue...))
		}
	}
	return result
}

func Test_permuteUnique(t *testing.T) {
	//nums := []int{-1, 2, 0, -1, 1, 0, 1}
	//nums := []int{1, 1, 2, 2, 3, 3, 4}
	nums := []int{1, 1, 2, 2, 3, 3, 4}
	res := permuteUnique(nums)
	t.Logf("%v\n", len(res))

}
