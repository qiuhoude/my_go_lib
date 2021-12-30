package leetcode

import "testing"

// 200. 岛屿数量 https://leetcode-cn.com/problems/number-of-islands/

/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1

输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'

思路:
与79题思路基本类似, 使用深度优先遍历。
从"1"点开始进行深度优先遍历，方向为上右下左的顺序进行, 1可以走0不能走加上边界判断，每走一个就标记以走后面就不能重复走,
直到不能走为止这就是一个岛屿.

*/

func numIslands(grid [][]byte) int {
	direction := [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} //  上，右，下，左 的顺序
	h, w := len(grid), len(grid[0])
	used := make([][]bool, h) // 坐标是否已经使用
	for i := range used {
		used[i] = make([]bool, w)
	}
	var dfsFn func(curX, curY int)
	dfsFn = func(curX, curY int) {
		used[curY][curX] = true // 占用
		for i := range direction {
			newX, newY := curX+direction[i][0], curY+direction[i][1]
			if 0 <= newX && newX < w && 0 <= newY && newY < h && // 边界
				grid[newY][newX] == '1' && // 陆地才能走
				!used[newY][newX] { // 未被占用
				dfsFn(newX, newY)
			}
		}
	}
	ans := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '1' && !used[y][x] {
				ans++
				dfsFn(x, y)
			}
		}
	}
	return ans
}

func Test_numIslands(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]byte
		expected int
	}{
		{"1", [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'}}, 1,
		},
		{"2", [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'}}, 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.grid); got != tt.expected {
				t.Errorf("numIslands() = %v, expected %v", got, tt.expected)
			}
		})
	}

}
