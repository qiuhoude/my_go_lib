package leetcode

import (
	"sort"
	"testing"
)

//18. 四数之和 https://leetcode-cn.com/problems/4sum/

/*

给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
0 <= a, b, c, d < n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。

输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]


1 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109

思路:
先排序 加 对撞指针,和15题3数之和一致,加上减少循环前置判断条件,比如前4个数加起来都大于target,这个就基本结果就基本结束了

*/

func Test_fourSum(t *testing.T) {
	res := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	t.Logf("%v", res)
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	size := len(nums)
	// j< i < l < r
	for j := 0; j < size-3; j++ {
		if j > 0 && nums[j] == nums[j-1] {
			continue
		}
		// 减少移动判断
		if nums[j]+nums[j+1]+nums[j+2]+nums[j+3] > target {
			break
		}
		for i := j + 1; i < size-2; i++ {
			if i != j+1 && nums[i] == nums[i-1] { // 重复的去掉
				continue
			}
			// 减少移动判断
			if nums[j]+nums[i]+nums[i+1]+nums[i+2] > target {
				break
			}

			l := i + 1
			r := len(nums) - 1

			for l < r {
				sum := nums[j] + nums[i] + nums[l] + nums[r]
				if sum == target { // 找到目的
					res = append(res, []int{nums[j], nums[i], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] { // 去重
						l++
					}
					for l < r && nums[r] == nums[r-1] { // 去重
						r--
					}
					l++
					r--
				} else if sum < target {
					l++
				} else if sum > target {
					r--
				}
			}
		}
	}
	return res
}
