package leetcode

import (
	"testing"
)

// 78. 子集 https://leetcode-cn.com/problems/subsets/

/*

给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。


输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

输入：nums = [0]
输出：[[],[0]]

1 <= nums.length <= 10
-10 <= nums[i] <= 10
nums 中的所有元素 互不相同

思路1:
循环+递归回溯的方式, 使用循环查找 递归回溯 n选1,n选2 n选{0...n}的子集
其中递归回溯和 77 题一致

思路2:
使用二进制表示法, 以为题目中集合的数量是[1:10]个, int32位足够,
例如 {1,2,3} 可以转换成3位二进制 000,001,010,...111,
000表示{}, 111表示{1,2,3}, 010表示{2}, 每个bit 0表不存在子集中


*/
func subsets4(nums []int) [][]int {
	var res [][]int
	var dfsFn func(int)
	var queue []int
	// 回溯法,和思路1类似
	dfsFn = func(start int) {
		res = append(res, append([]int(nil), queue...)) // 每次进来
		for i := start; i < len(nums); i++ {
			if i != start && nums[i] == nums[i-1] {
				continue
			}
			queue = append(queue, nums[i])
			dfsFn(i + 1)
			queue = queue[:len(queue)-1]
		}
	}
	dfsFn(0)
	return res
}

func subsets3(nums []int) [][]int {
	var dfs func(int)
	var res [][]int
	var set []int
	dfs = func(cur int) {
		if cur == len(nums) {
			res = append(res, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return res
}

func subsets2(nums []int) [][]int {
	/*
		思路
		使用二进制表示法, 以为题目中集合的数量是[1:10]个, int32位足够,
	例如 {1,2,3} 可以转换成3位二进制 000,001,010,...111,
	000表示{}, 111表示{1,2,3}, 010表示{2}, 每个bit 0表不存在子集中
	*/
	n := 1 << uint(len(nums))
	var res [][]int
	for i := 0; i < n; i++ {
		var set []int
		c := uint(i)
		for j := 0; c > 0; j++ {
			if c&1 == 1 {
				set = append(set, nums[j])
			}
			c = c >> 1
		}
		res = append(res, set)
	}
	return res
}

func subsets(nums []int) [][]int {
	res := [][]int{{}} // 先添加一个空集合
	var dfs = func(selectCnt, start int, queue []int) {}
	dfs = func(selectCnt, start int, queue []int) {
		if len(queue) == selectCnt {
			res = append(res, append([]int(nil), queue...))
			return
		}
		remain := selectCnt - len(queue)
		end := len(nums) - remain + 1 // 剪枝
		//fmt.Printf("start:%v ,end:%v \n", start, end)
		for i := start; i < end; i++ {
			queue = append(queue, nums[i])
			dfs(selectCnt, i+1, queue)
			queue = queue[:len(queue)-1]
		}
	}
	for i := 1; i < len(nums); i++ {
		dfs(i, 0, nil)
	}
	res = append(res, append([]int(nil), nums...)) // 添加一个全部集合
	return res
}

func Test_subsets(t *testing.T) {
	res := subsets4([]int{1, 2, 2})
	t.Logf("%v\n", res)
}
