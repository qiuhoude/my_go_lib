package leetcode

import (
	"math/bits"
	"testing"
)

// 37. 解数独 https://leetcode-cn.com/problems/sudoku-solver/

/*
编写一个程序，通过填充空格来解决数独问题。
数独的解法需 遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
数独部分空格内已填入了数字，空白格用 '.' 表示。

输入：board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
输出：[["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
解释：输入的数独如上图所示，唯一有效的解决方案如下所示：

board.length == 9
board[i].length == 9
board[i][j] 是一位数字或者 '.'
题目数据 保证 输入数独仅有一个解

思路:
使用递归回溯的方式, 每经过"."的位置都会进行判断横,纵,3x3格子,算出当前还有哪些可以填
可以使用额外的空间存储 横,纵,3x3格子 占用情况, 可以使用bitmap存储占用情况

*/

// 位运算优化
func solveSudoku3(board [][]byte) {
	var rowUsed, colUsed [9]int // 行列占用
	var block [3][3]int         // 3x3块的占用
	var needFillIn [][2]int     // 需要填入数字的坐标 [0]row [1]col
	// 行 列 占用的数字 ,index=0~8,对应的位运算 ,使用异或 0->1或1->0
	fillFn := func(r, c, index int) {
		rowUsed[r] ^= 1 << uint(index)
		colUsed[c] ^= 1 << uint(index)
		block[r/3][c/3] ^= 1 << uint(index)
	}
	for r := range board {
		for c := range board[r] {
			if board[r][c] == '.' {
				needFillIn = append(needFillIn, [2]int{r, c})
			} else { // 设置占用情况
				digit := int(board[r][c] - '1')
				fillFn(r, c, digit)
			}
		}
	}
	mask := 1<<9 - 1
	// index 是值得 needFillIn得下标
	var dfsFn func(index int) bool
	dfsFn = func(index int) bool {
		if index == len(needFillIn) {
			return true
		}
		// 在 board[row][col] 尝试填入
		row, col := needFillIn[index][0], needFillIn[index][1]
		/*for i := 0; i < 9; i++ {
			bits := 1 << uint(i)
			if bits&rowUsed[row] == 0 && bits&colUsed[col] == 0 && bits&block[row/3][col/3] == 0 {
				fillFn(row, col, i)
				board[row][col]= byte(i+'1')
				if dfsFn(index + 1) {
					return true
				}
				fillFn(row, col, i)
			}
		}*/
		// 使用位操作减少for循环次数
		for bit := mask & (^(rowUsed[row] | colUsed[col] | block[row/3][col/3])); bit > 0; bit &= bit - 1 {
			pos := bit & -bit
			// pos 的数中只有1个, pos-1 表示 原来是1的那个位置变成0,后面则全变成1, 然后统计1的数量,就是pos那个1的位置
			oneIndex := bits.OnesCount(uint(pos - 1))
			fillFn(row, col, oneIndex)
			board[row][col] = byte(oneIndex + '1')
			if dfsFn(index + 1) {
				return true
			}
			fillFn(row, col, oneIndex)
		}
		return false
	}
	dfsFn(0)
}

// 优化内存
func solveSudoku2(board [][]byte) {
	n := len(board)
	// 判断填写的数字是否满足
	checkNumFn := func(row, column, num int) bool {
		char := byte(num + '1')
		rStart, cStart := (row/3)*3, (column/3)*3
		for i := 0; i < n; i++ {
			// 横向
			if board[row][i] == char {
				return false
			}
			// 纵向
			if board[i][column] == char {
				return false
			}
			// 3x3 小方块
			if board[rStart+(i/3)][cStart+(i%3)] == char {
				return false
			}
		}
		return true
	}

	var dfsFn func(index int) bool
	dfsFn = func(index int) bool {
		row, column := index/n, index%n
		if row == n-1 && column == n-1 { // 最后一步
			if board[row][column] != '.' {
				return true
			} else {
				for i := 0; i < 9; i++ {
					if checkNumFn(row, column, i) {
						board[row][column] = byte(i + '1')
						return true
					}
				}
				return false
			}
		}
		if board[row][column] == '.' {
			for i := 0; i < 9; i++ {
				if checkNumFn(row, column, i) {
					board[row][column] = byte(i + '1')
					if dfsFn(index + 1) {
						return true
					}
					board[row][column] = '.'
				}
			}

		} else {
			return dfsFn(index + 1)
		}
		return false
	}
	dfsFn(0)
}

func solveSudoku(board [][]byte) {
	n := len(board)
	//计算当前还有数字可以填
	calcNumFn := func(row, column int) []byte {
		used := make([]bool, 9)

		// 横向
		for c := 0; c < n; c++ {
			if c == column {
				continue
			}
			if board[row][c] != '.' {
				used[board[row][c]-'1'] = true
			}
		}
		// 纵向
		for r := 0; r < n; r++ {
			if r == row {
				continue
			}
			if board[r][column] != '.' {
				used[board[r][column]-'1'] = true
			}
		}
		// 3x3 小方块
		rStart, cStart := (row/3)*3, (column/3)*3
		for r := rStart; r < rStart+3; r++ {
			for c := cStart; c < cStart+3; c++ {
				if c == column && r == rStart {
					continue
				}
				if board[r][c] != '.' {
					used[board[r][c]-'1'] = true
				}
			}
		}
		var arr []byte
		for i, b := range used {
			if !b {
				arr = append(arr, byte('1'+i))
			}
		}
		return arr
	}

	var dfsFn func(index int) bool
	dfsFn = func(index int) bool {
		row, column := index/n, index%n
		if row == n-1 && column == n-1 { // 最后一步
			if board[row][column] != '.' {
				return true
			} else {
				arr := calcNumFn(row, column)
				if len(arr) > 0 {
					board[row][column] = arr[0]
					return true
				} else {
					return false
				}
			}
		}
		if board[row][column] == '.' {
			arr := calcNumFn(row, column)
			for _, char := range arr {
				board[row][column] = char
				if dfsFn(index + 1) {
					return true
				}
				board[row][column] = '.'
			}
		} else {
			return dfsFn(index + 1)
		}
		return false
	}
	dfsFn(0)
}

func Test_solveSudoku(t *testing.T) {
	//board := [][]byte{
	//	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	//	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	//	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	//	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	//	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	//	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	//	{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	board := [][]byte{
		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '2', '7', '5', '9', '.', '.'}}
	solveSudoku2(board)
	t.Logf("%q\n", board)

	//bit := 21
	//for ; bit > 0; bit &= bit - 1 {
	//	pos := bit & -bit
	//	oneIndex := bits.OnesCount(uint(pos - 1))
	//	t.Logf("%d bit=%b pos=%b oneIndex=%d", bit, bit, pos, oneIndex)
	//}
}
