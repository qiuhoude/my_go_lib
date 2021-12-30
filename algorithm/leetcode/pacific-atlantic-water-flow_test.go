package leetcode

// 417. 太平洋大西洋水流问题 https://leetcode-cn.com/problems/pacific-atlantic-water-flow/

/*
给定一个 m x n 的非负整数矩阵来表示一片大陆上各个单元格的高度。“太平洋”处于大陆的左边界和上边界，而“大西洋”处于大陆的右边界和下边界。
规定水流只能按照上、下、左、右四个方向流动，且只能从高到低或者在同等高度上流动。
请找出那些水流既可以流动到“太平洋”，又能流动到“大西洋”的陆地单元的坐标。

输出坐标的顺序不重要
m 和 n 都小于150

给定下面的 5x5 矩阵:

  太平洋 ~   ~   ~   ~   ~
       ~  1   2   2   3  (5) *
       ~  3   2   3  (4) (4) *
       ~  2   4  (5)  3   1  *
       ~ (6) (7)  1   4   5  *
       ~ (5)  1   1   2   4  *
          *   *   *   *   * 大西洋

返回:
[[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (上图中带括号的单元).

思路1:
对每个坐标进行DFS, 遇到大西洋标记已到, 遇到太平洋记已到 , 然后就停止遍历 (效率极低)

思路2:
找到能达到太平的的所有点放到setLU，找到能到达大西洋的所有点放到setRD中,找出set1与set2的交集
*/

func pacificAtlantic2(heights [][]int) [][]int {
	direction := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} // 左,上，右，下 的顺序
	h, w := len(heights), len(heights[0])                   // 边界
	setLU := make([][]bool, h)                              // 标记集合
	setRD := make([][]bool, h)
	for i := 0; i < h; i++ {
		setLU[i] = make([]bool, w)
		setRD[i] = make([]bool, w)
	}
	var dfsFn func(curX, curY int, set [][]bool)
	dfsFn = func(curX, curY int, set [][]bool) {
		if set[curY][curX] {
			return
		}
		set[curY][curX] = true
		for i := range direction {
			newX, newY := curX+direction[i][0], curY+direction[i][1]
			if 0 <= newX && newX < w && 0 <= newY && newY < h && // 边界
				heights[newY][newX] >= heights[curY][curX] && // 逆流
				!set[newY][newX] {
				dfsFn(newX, newY, set)
			}
		}
	}
	for x := 0; x < w; x++ {
		dfsFn(x, 0, setLU)
		dfsFn(x, h-1, setRD)
	}
	for y := 0; y < h; y++ {
		dfsFn(0, y, setLU)
		dfsFn(w-1, y, setRD)
	}
	// 找交集
	var res [][]int
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if setLU[y][x] && setRD[y][x] {
				res = append(res, []int{y, x})
			}
		}
	}
	return res
}

func pacificAtlantic(heights [][]int) [][]int {
	direction := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} // 左,上，右，下 的顺序
	h, w := len(heights), len(heights[0])                   // 边界
	used := make([][]bool, h)                               // 标记
	for i := range used {
		used[i] = make([]bool, w)
	}

	lu, rd := false, false // leftUp,rightDown ,标记太平洋和大西洋
	var dfsFn func(curX, curY int)
	dfsFn = func(curX, curY int) {
		if curX == 0 || curY == 0 {
			lu = true
		}
		if curX == w-1 || curY == h-1 {
			rd = true
		}
		if lu && rd {
			return
		}
		used[curY][curX] = true
		for i := range direction {
			newX, newY := curX+direction[i][0], curY+direction[i][1]
			if 0 <= newX && newX < w && 0 <= newY && newY < h && // 边界
				heights[newY][newX] < heights[curY][curX] &&
				!used[newY][newX] { // 未被使用
				dfsFn(newX, newY)
			}
		}
		used[curY][curX] = false
	}
	var res [][]int
	for y := range heights {
		for x := range heights[y] {
			lu, rd = false, false
			dfsFn(x, y)
			if lu && rd {
				res = append(res, []int{y, x})
			}
		}
	}
	return res
}
