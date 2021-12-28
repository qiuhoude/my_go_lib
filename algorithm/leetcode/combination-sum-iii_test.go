package leetcode

// 216. 组合总和 III https://leetcode-cn.com/problems/combination-sum-iii/

/*

找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

说明：

所有数字都是正整数。
解集不能包含重复的组合。

输入: k = 3, n = 7
输出: [[1,2,4]]

输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]

思路:
递归回溯 与 40 思路也基本一致
每次都从 [cur:9] 选取一个数, 递归的选取
*/

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	queue := []int{}
	var recursionFn = func(cur, sum int) {}
	recursionFn = func(cur, sum int) {
		if len(queue) == k {
			if sum == n {
				res = append(res, append([]int(nil), queue...))
			}
			return
		}
		if n-sum < cur {
			return
		}
		for i := cur; i <= 9; i++ {
			queue = append(queue, i)
			recursionFn(i+1, sum+i) // i+1 为去除本身
			queue = queue[:len(queue)-1]
		}

	}
	recursionFn(1, 0)
	return res
}

//func Test_combinationSum3(t *testing.T) {
//	//res := combinationSum3(3, 7)
//	//res := combinationSum3(3, 9)
//	res := combinationSum3(4, 1)
//	t.Logf("%v\n", res)
//}
