package leetcode

import (
	"sort"
)

// 39. 组合总和 https://leetcode-cn.com/problems/combination-sum/

/*
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有不同组合 ，
并以列表形式返回。你可以按 任意顺序 返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
对于给定的输入，保证和为 target 的不同组合数少于 150 个。

输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。

输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]

输入: candidates = [2], target = 1
输出: []

输入: candidates = [1], target = 1
输出: [[1]]

输入: candidates = [1], target = 2
输出: [[1,1]]


1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都 互不相同
1 <= target <= 500

思路:
递归回溯
每次都从 candidates[cur:] 选取一个数,因为可以同一个元素可以重复选择,所有递归时不用i+1
先对数组进行排序,可以利用 target-sum(cur) 进行剪枝操作,如后面的数都大于这个数不进行递归
*/

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates) // 排序,用于后面的剪枝

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
			queue = append(queue, candidates[i])
			recursion(i, curSum+candidates[i]) // 因为可以重复选着当前元素,所有是i, 如果不能重复选择则就是i+1
			queue = queue[:len(queue)-1]
		}
	}
	recursion(0, 0)
	return res
}

/*func Test_combinationSum(t *testing.T) {
	//res := combinationSum([]int{2, 3, 6, 7}, 7)
	//res := combinationSum([]int{1}, 1)
	res := combinationSum([]int{2, 3, 5}, 8)
	t.Logf("%v\n", res)
}
*/
