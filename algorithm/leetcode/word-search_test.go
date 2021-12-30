package leetcode

import (
	"testing"
)

// 79. 单词搜索 https://leetcode-cn.com/problems/word-search/

/*
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。


输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
输出：true

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
输出：false

m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board 和 word 仅由大小写英文字母组成
进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？

思路:
使用递归回溯的方式, 依次比较word与board每个字符，board下一个字符是上右下左的顺序(经历过的不能走，边界不能走)，word只需要++index就可以
如果不相等board就退回上一步，直到word的 index==len(word) 就找到。 如果board每个字符当作首字符都试过并且没有找到，说明就是没有


*/

func exist(board [][]byte, word string) bool {
	direction := [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} //  上，右，下，左 的顺序
	h, w := len(board), len(board[0])                       // 高(y)，宽(x)
	used := make([][]bool, h)                               // 坐标是否已经使用
	for i := range used {
		used[i] = make([]bool, w)
	}
	var dfsFn func(curX, curY, curIndex int) bool
	dfsFn = func(curX, curY, curIndex int) bool {
		if board[curY][curX] != word[curIndex] { // 不相等直接返回
			return false
		}
		if curIndex >= len(word)-1 {
			return true
		}
		ret := false
		used[curY][curX] = true // 占用
		for i := range direction {
			dx, dy := direction[i][0], direction[i][1]
			newX, newY := dx+curX, dy+curY
			if 0 <= newX && newX < w && 0 <= newY && newY < h && !used[newY][newX] { // 在边界中，并且没有被使用
				if dfsFn(newX, newY, curIndex+1) {
					ret = true
					break
				}
			}
		}
		used[curY][curX] = false // 取消占用
		return ret
	}

	for i := range board {
		for j := range board[i] {
			if dfsFn(j, i, 0) {
				return true // 找到
			}
		}
	}
	return false
}

func Test_exist(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		board    [][]byte
		expected bool
	}{
		{"ABCCED", "ABCCED", [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, true},
		{"SEE", "SEE", [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, true},
		{"ABCB", "ABCB", [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, false},
		{"a", "a", [][]byte{{'a'}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.board, tt.word); got != tt.expected {
				t.Errorf("exist() = %v, expected %v", got, tt.expected)
			}
		})
	}

}
