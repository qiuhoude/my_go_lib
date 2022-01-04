package leetcode

// 52. N皇后 II https://leetcode-cn.com/problems/n-queens-ii/

/*
n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。

输入：n = 4
输出：2

输入：n = 1
输出：1

1 <= n <= 9


思路:
使用二进制的位运算 + 递归回溯的方式, 对比普通的回溯, 这种方式可以使用二进制的位运算进行对不满足条件的位置进行剪枝操作

使用3个数的二级制（题目中 n<=9 int32有32位足够表示），分别表示 列(columns)，左右斜对角(LDiagonals,RDiagonals)占用情况(1表是占用 0表示位占用),
公式:  (1<<n -1)&(~(columns | LDiagonals | RDiagonals)) 1位 表示下一行可以选择的位置
 x&(-x) 获取x二进制最后出现位为1的位置 ,（其中奇数永远是最后一个位置是1）
 x&(x-1) 将x二进制最后出现位1的位 置成0

*/

func totalNQueens(n int) int {
	ret := 0
	var dfs func(index, columns, LDiagonals, RDiagonals int)
	dfs = func(index, columns, LDiagonals, RDiagonals int) {
		if index == n {
			ret++ // 满足条件
			return
		}
		// 下一行可以选择的位置, bit 1表示可以选
		availablePosition := (1<<uint(n) - 1) & (^(columns | LDiagonals | RDiagonals))
		for ; availablePosition > 0; availablePosition &= availablePosition - 1 {
			position := availablePosition & -availablePosition
			dfs(index+1, position|columns, (position|LDiagonals)>>1, (position|RDiagonals)<<1)
		}
	}
	dfs(0, 0, 0, 0)
	return ret
}

//func Test_totalNQueens(t *testing.T) {
//var c int = 23
//t.Logf("%b %b %b", c, c&(-c), c&(c-1))
//t.Logf("%b %b", 3, (^3)+1)
//res := totalNQueens(2)
//t.Log(res)
//}
