package leetcode

import (
	"fmt"
	"math"
	"testing"
)

// 149. 直线上最多的点数 https://leetcode-cn.com/problems/max-points-on-a-line/

/*
给你一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点。求最多有多少个点在同一条直线上。

输入：points = [[1,1],[2,2],[3,3]]
输出：3

输入：points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
输出：4

1 <= points.length <= 300
points[i].length == 2
-10^4 <= xi, yi <= 10^4
points 中的所有点 互不相同

思路:
查表法, 直线方程的 斜截式 y=kx+b, k=(y1-y2)/(x1-x2)=dy/dx
遍历所有点求出两两点之间的斜率, 例如取i点, 求出i点到其他点的斜率当作key,斜率出现频率当val存入map中,算出频率最高的值
两直线斜率相等就在一条直线上？ 答: 如果两条直线的斜率相等并且都经过同一点i，那么它们肯定是同一条直线
斜率的表示方式,如果直接使用 dy/dx 定会有浮点数存在, 可以将其转换成两个值最简式进行存储,例如 3/6 -> 1/3,dy和dx都除他们两个最大公约数可得
k=dy/dx=1/3,可以使用[2]int16{dx,dy}当key进行存储
分母为0情况[0][1]表示,分子为0情况[1][0]表示
负数的情况,统一让dy为正数

优化方向，减少循环的次数
当前已经找到线上最多点 大于 len(points)-i（剩余没有遍历的点）说明已经找到最多了不用继续下去
*/
func maxPoints(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}
	res := 0
	tabK := make(map[[2]int16]int) // <斜率,次数>
	for i := 0; i < len(points); i++ {
		if res > len(points)-i || res > len(points)/2 {
			break
		}
		// 清空map
		for k := range tabK {
			delete(tabK, k)
		}
		for j := i + 1; j < len(points); j++ {
			dy := points[i][1] - points[j][1]
			dx := points[i][0] - points[j][0]
			var key [2]int16
			if dy == 0 {
				key = [2]int16{1, 0}
			} else if dx == 0 {
				key = [2]int16{0, 1}
			} else {
				// 求最大公约数
				bcd := bcd(int(math.Abs(float64(dy))), int(math.Abs(float64(dx))))
				if dy < 0 { //让dy必须为正数
					dx, dy = -dx, -dy
				}
				key = [2]int16{int16(dx / bcd), int16(dy / bcd)}
			}
			tabK[key]++
			cnt := tabK[key]
			if cnt+1 > res {
				res = cnt + 1
			}
		}
	}
	return res
}

// 最大公约数辗转相除法
func bcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func Test_maxPoints(t *testing.T) {
	tests := []struct {
		expected int
		arg1     [][]int
	}{
		{3, [][]int{{9, -25}, {-4, 1}, {-1, 5}, {-7, 7}}},
	}
	for _, tt := range tests {
		res := maxPoints(tt.arg1)
		t.Run(fmt.Sprintf("%v", tt.arg1), func(t *testing.T) {
			if res != tt.expected {
				t.Errorf("expected:%v but got:%v", tt.expected, res)
			}
		})

	}
}
