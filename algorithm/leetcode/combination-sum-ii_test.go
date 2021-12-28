package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

// 40. 组合总和 II https://leetcode-cn.com/problems/combination-sum-ii/
/*
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用一次。
注意：解集不能包含重复的组合。

输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]

输入: candidates = [2,5,2,1,2], target = 5,
输出:
[
[1,2,2],
[5]
]

1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30

思路:
递归回溯,与39题思路一致
每次都从 candidates[cur:] 选取一个数,不能选择重复元素,所有递归时需要i+1排除当前值
先对数组进行排序,可以利用 target-sum(cur) 进行剪枝操作,如后面的数都大于这个数不进行递归

*/

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates) // 排序,用于后面的剪枝
	fmt.Println(candidates)
	queue := []int{}
	var recursion func(curIndex, curSum int)
	recursion = func(curIndex, curSum int) {
		if curSum >= target {
			if curSum == target {
				res = append(res, append([]int(nil), queue...))
			}
			return
		}

		for i := curIndex; i < len(candidates); i++ {
			if target-curSum < candidates[curIndex] { // 后面的数已经超过了
				break
			}
			if i != curIndex && i > 0 && candidates[i] == candidates[i-1] { // 与前面数相等则跳过,i != curIndex 为了cur本身的递归继续下去
				continue
			}
			queue = append(queue, candidates[i])
			recursion(i+1, curSum+candidates[i]) //i+1排除当前值
			queue = queue[:len(queue)-1]
		}
	}
	recursion(0, 0)
	return res
}

func Test_combinationSum2(t *testing.T) {

	//res := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	//res := combinationSum2([]int{2,5,2,1,2}, 5)
	//res := combinationSum2([]int{2, 2, 2, 1, 2, 1, 1}, 5)
	res := combinationSum2([]int{3, 1, 3, 5, 1, 1}, 8)
	t.Logf("%v\n", res)
}
