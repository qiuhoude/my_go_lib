package leetcode

// 447. 回旋镖的数量 https://leetcode-cn.com/problems/number-of-boomerangs/

/*

给定平面上 n 对 互不相同 的点 points ，其中 points[i] = [xi, yi] 。
回旋镖 是由点 (i, j, k) 表示的元组 ，其中 i 和 j 之间的距离和 i 和 k 之间的欧式距离相等（需要考虑元组的顺序）。
返回平面上所有回旋镖的数量。

输入：points = [[0,0],[1,0],[2,0]]
输出：2
解释：两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]

输入：points = [[1,1],[2,2],[3,3]]
输出：2

输入：points = [[1,1]]
输出：0

n == points.length
1 <= n <= 500
points[i].length == 2
-104 <= xi, yi <= 104
所有点都 互不相同

思路:
查表法
遍历所有点,将所有两点之间的距离存储到map中,
把距离当作key,某个点出现频率当作value, value>=2说明有回旋镖产生
统计在相等距离的情况下,某个点i出现的次数,i为中间点产生回旋镖的个数=cnt*(cnt-1) cnt>=2
*/

// 优化版本
func numberOfBoomerangs(points [][]int) int {
	res := 0
	if len(points) < 3 {
		return res
	}
	// 优化: 不用map嵌套map把所有距离情况记录下来再进行计算
	distinctMap := make(map[int]int) // key:距离, 出现次数

	for i := 0; i < len(points); i++ {
		// 将i作为中间点
		//distinctMap = make(map[int]int) // 清空map
		//clear map
		for k := range distinctMap {
			delete(distinctMap, k)
		}

		for j := 0; j < len(points); j++ {
			if i != j {
				// 距离:Sqrt((x1-x2)^2 + (y1-y2)^2)
				xd := points[i][0] - points[j][0]
				yd := points[i][1] - points[j][1]
				dis := xd*xd + yd*yd
				distinctMap[dis]++
			}
		}
		for _, cnt := range distinctMap {
			if cnt >= 2 {
				res += cnt * (cnt - 1)
			}
		}
	}
	return res
}

// 最初版本
func numberOfBoomerangsOld(points [][]int) int {
	res := 0
	if len(points) < 3 {
		return res
	}
	// 两点之间距离当key1, 把回旋的公共点位置当作key2,出现的次数为val,同一距离出现次数>=2才能形成回旋
	distinctMap := make(map[int]map[int]int)

	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			// 距离:Sqrt((x1-x2)^2 + (y1-y2)^2)
			xd := points[i][0] - points[j][0]
			yd := points[i][1] - points[j][1]
			dis := xd*xd + yd*yd
			if _, ok := distinctMap[dis]; !ok { // 不存在创建
				distinctMap[dis] = make(map[int]int)
			}
			distinctMap[dis][i]++
			distinctMap[dis][j]++

		}
	}
	for _, m := range distinctMap {
		for _, cnt := range m {
			// cnt 出现, 对应的可能情况 2次->2*1, 3次->3*2, 4次->4*3 公式: cnt*(cnt-1)
			if cnt >= 2 {
				res += cnt * (cnt - 1)
			}
		}
	}
	return res
}
