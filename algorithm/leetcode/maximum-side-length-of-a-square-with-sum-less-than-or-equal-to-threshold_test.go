package leetcode

import (
	"github.com/bmizerany/assert"
	"math"
	"testing"
)

// 1292. 元素和小于等于阈值的正方形的最大边长 https://leetcode.cn/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/

/*
给你一个大小为 m x n 的矩阵 mat 和一个整数阈值 threshold。

请你返回元素总和小于或等于阈值的正方形区域的最大边长；如果没有这样的正方形区域，则返回 0 。


示例 1：

输入：mat = [[1,1,3,2,4,3,2],[1,1,3,2,4,3,2],[1,1,3,2,4,3,2]], threshold = 4
输出：2
解释：总和小于或等于 4 的正方形的最大边长为 2，如图所示。

示例 2：

输入：mat = [[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2]], threshold = 1
输出：0

提示：

m == mat.length
n == mat[i].length
1 <= m, n <= 300
0 <= mat[i][j] <= 104
0 <= threshold <= 105

思路:

1. 构建二维数组前缀和, 公式: P[i][j] = P[i - 1][j] + P[i][j - 1] - P[i - 1][j - 1] + mat[i - 1][j - 1]; (小技巧初始化的时多创建1列和1行)
2. 正方形的区域边长公式  左上点(x1,y1), 右下点(x2,y2)  P[x2][y2] - P[x1-1][y2] - P[x2][y1-1] + P[x1-1][y1-1]
3. 答案在 1 ~ min(w,h) 之间, 所以可以进行二分查找, 检查是否满足阈值条件, 满足 -> 右边找，反之左边
*/

func maxSideLength(mat [][]int, threshold int) int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}
	h, w := len(mat), len(mat[0])
	// 算出前缀和
	P := make([][]int, h+1)
	P[0] = make([]int, w+1)
	for i := 1; i <= h; i++ {
		P[i] = make([]int, w+1)
		for j := 1; j <= w; j++ {
			P[i][j] = P[i-1][j] + P[i][j-1] - P[i-1][j-1] + mat[i-1][j-1]
		}
	}

	l, r, ans := 1, int(math.Min(float64(h), float64(w))), 0

	// 正方形的范围
	getRect := func(x1, y1, x2, y2 int) int {
		return P[x2][y2] - P[x1-1][y2] - P[x2][y1-1] + P[x1-1][y1-1]
	}

	findRect := func(mid int) bool {
		for i := 1; i <= h-mid+1; i++ {
			for j := 1; j <= w-mid+1; j++ {
				if getRect(i, j, i+mid-1, j+mid-1) <= threshold {
					return true
				}
			}
		}
		return false
	}

	for l <= r {
		mid := l + ((r - l) >> 1)
		if findRect(mid) {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return ans
}

func Test_maxSideLength(t *testing.T) {
	// mat = [[1,1,3,2,4,3,2],[1,1,3,2,4,3,2],[1,1,3,2,4,3,2]], threshold = 4 , expected =2
	mat := [][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}}
	threshold := 4
	expected := 2
	got := maxSideLength(mat, threshold)
	assert.Equal(t, expected, got)
}
