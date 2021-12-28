package leetcode

// 77. 组合 https://leetcode-cn.com/problems/combinations/

/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。

输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

输入：n = 1, k = 1
输出：[[1]]

1 <= n <= 20
1 <= k <= n

思路:
利用递归回溯的思想,将其转换成树形结构
每次取一个数出来, 然后再剩余的数字再取一个出... 依次类推,直到取k个数
如果已经使用的没有了就不递归下去.
优化:
对树形结构不满足条件的进行剪枝

*/

func combine(n int, k int) [][]int {
	var res [][]int
	var queue []int
	var recursionCombineFn = func(int) {}
	recursionCombineFn = func(cur int) {
		if len(queue) == k {
			res = append(res, append([]int(nil), queue...))
			return
		}
		// 剩余的空位 = k- len(queue)
		// [cur...n] 中至少有多余剩余空位的元素才有意义,也就是求cur结束位置
		// 结束位置 + 剩余的空位 - 1 = n; => 结束位置 = n - 剩余的空位 + 1;
		remain := k - len(queue)
		end := n - remain + 1 // 对剩余的部分进行剪枝, 余量不够填充将队列长度到k
		//fmt.Printf("[%v:%v] : %v\n", cur, end, remain)
		for i := cur; i <= end; i++ {
			queue = append(queue, i)
			recursionCombineFn(i + 1)
			queue = queue[:len(queue)-1]
		}
	}
	recursionCombineFn(1)
	return res
}

//func Test_combine(t *testing.T) {
//	res := combine(5, 3)
//	t.Logf("%v\n", res)
//}
