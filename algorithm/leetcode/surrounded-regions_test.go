package leetcode

import "testing"

// 130. 被围绕的区域 https://leetcode-cn.com/problems/surrounded-regions/

/*

给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，
或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。

输入：board = [["X"]]
输出：[["X"]]


m == board.length
n == board[i].length
1 <= m, n <= 200
board[i][j] 为 'X' 或 'O'

思路:
和200题找岛屿数量很类似, 只需要找与边界相连接的 "O"开始DFS找到相连并且给予标记, 剩余的"O"就是都要换成"X"
*/

func solve(board [][]byte) {
	direction := [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} // 上，右，下，左 的顺序
	h, w := len(board), len(board[0])                       // 边界
	mark := make([][]bool, h)                               // 标记
	for i := range mark {
		mark[i] = make([]bool, w)
	}
	var dfsFn func(curX, curY int)
	dfsFn = func(curX, curY int) {
		if board[curY][curX] != 'O' || mark[curY][curX] {
			return
		}
		mark[curY][curX] = true // 标记
		for i := range direction {
			newX, newY := curX+direction[i][0], curY+direction[i][1]
			if 0 <= newX && newX < w && 0 <= newY && newY < h && // 边界
				board[newY][newX] == 'O' &&
				!mark[newY][newX] { // 未被标记
				dfsFn(newX, newY)
			}
		}
	}
	// 从边界开始找 'O',使用dfs标记与边界相连接的 'O'
	for x := 0; x < w; x++ {
		dfsFn(x, 0)
		dfsFn(x, h-1)
	}
	for y := 0; y < h; y++ {
		dfsFn(0, y)
		dfsFn(w-1, y)
	}
	//  剩余的"O"就是都要换成"X"
	for y := 1; y < h-1; y++ {
		for x := 1; x < w-1; x++ {
			if board[y][x] == 'O' && !mark[y][x] {
				board[y][x] = 'X'
			}
		}
	}
}

func Test_solve(t *testing.T) {
	//board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	//board := [][]byte{{'X'}}
	board := [][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}}
	solve(board)
	t.Logf("%q", board)
}
