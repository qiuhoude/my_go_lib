package leetcode

import (
	"fmt"
	"testing"
)

// 51. N皇后  https://leetcode-cn.com/problems/n-queens/
/*
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。

输入：n = 1
输出：[["Q"]]

1 <= n <= 9

思路:
递归回溯, 尝试排把皇后放到第一个排每一列进行尝试, 进行深度优先遍历
难点在于判断皇后 在 横 竖 左右斜 线上

*/

func solveNQueens(n int) [][]string {
	record := make([]int, n) // 记录每行皇后方的位置, index:表示行, value:表示皇后放的列
	var retStr [][]string
	calQueens(0, n, record, &retStr)
	return retStr
}

func calQueens(row, n int, record []int, retStr *[][]string) {
	if row >= n {
		// 找到解
		solve := make([]string, n)
		c := make([]rune, n)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if record[i] == j {
					c[j] = 'Q'
				} else {
					c[j] = '.'
				}
			}
			solve[i] = string(c)
			fmt.Println(string(c))
		}
		*retStr = append(*retStr, solve)
		fmt.Println("-----------")
		return
	}
	for col := 0; col < n; col++ { // 每行都有n种可能
		if isOk1(row, col, n, record) {
			record[row] = col                   // 第 row 行的棋子放到了 column 列
			calQueens(row+1, n, record, retStr) // 如果满足条件就进行尝试下一行
		}
	}
}

/*

检测横竖左右斜是否有其他皇后

方式1:
00 01 02 03
10 11 12 13
20 21 22 23
30 31 32 33

是否在同一个左斜对角上 xIndex + yIndex = 恒定值
是否在同一个右斜对角上 xIndex - yIndex = 恒定值
是否在同一个竖线上 xIndex = curX

方式2:
逐行检测,从当前行的上一行检测开始检测,是否列相同,对角线相同
*/
func isOk1(curY, curX, n int, record []int) bool {
	for y := curY - 1; y >= 0; y-- {
		x := record[y]
		if curX == x || x-y == curX-curY || x+y == curX+curY { //同一个竖线上 左右斜对焦
			return false
		}
	}
	return true
}

func isOk2(row, col, n int, record []int) bool {
	lCol := col - 1 // 左边列
	rCol := col + 1 // 右边列
	// 从当前行的上一行检测开始检测
	for r := row - 1; r >= 0; r-- {
		if record[r] == col { // 纵向有相同的
			return false
		}
		// 左上对角线
		if lCol >= 0 && lCol == record[r] {
			return false
		}
		// 右上对角线
		if rCol < n && rCol == record[r] {
			return false
		}
		lCol--
		rCol++
	}
	return true
}

func TestSolveNQueens(t *testing.T) {
	solveNQueens(4)
}
