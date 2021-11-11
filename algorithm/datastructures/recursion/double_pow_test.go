package recursion

import (
	"github.com/bmizerany/assert"
	"testing"
)

// 实现 double 类型的n次方

// 循环的方式实现 O(n) 级别的时间复杂度
func double_pow_loop(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	var c int
	if n < 0 {
		c = -n
	} else {
		c = n
	}
	var rt = x
	for i := 0; i < c-1; i++ {
		rt = rt * x
	}

	if n < 0 {
		return 1.0 / rt
	} else {
		return rt
	}
}

// 递归的方式实现 O(logn) 级别
// 思路: 要求x 的n次方, 只需要求 n/2 次方, 如果 n是偶数 n/2的结果相乘就ok
// 如果是 n是奇数,只需要在乘以x就可以
func double_pow_recursion(x float64, n int) float64 {

	if n == 0 {
		return 1.0
	}
	var c int
	if n < 0 {
		c = -n
	} else {
		c = n
	}
	t := double_pow_recursion(x, c/2)

	var rt float64
	if c%2 == 0 { //n 是偶数
		rt = t * t
	} else {
		rt = x * t * t
	}

	if n < 0 {
		return 1 / rt
	} else {
		return rt
	}
}

func Test_double_pow_loop(t *testing.T) {
	assert.Equal(t, 16.0, double_pow_loop(2.0, 4))
	assert.Equal(t, 0.04, double_pow_loop(5.0, -2))

}

func Test_double_pow_recursion(t *testing.T) {
	assert.Equal(t, 16.0, double_pow_recursion(2.0, 4))
	assert.Equal(t, 0.04, double_pow_recursion(5.0, -2))

}
