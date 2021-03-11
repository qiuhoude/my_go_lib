package binary_search

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestBSearch(t *testing.T) {
	arr := []int{1, 3, 4, 5, 6, 7}
	r1 := BSearch(arr, 5)
	assert.Equal(t, 3, r1)
	r2 := BSearchRecusion(arr, 5)
	assert.Equal(t, 3, r2)
}

func TestBSearch2(t *testing.T) {
	arr := []int{1, 3, 4, 5, 6, 6, 8, 8, 8, 11, 18} //重复元素的切片

	first := BSearchFirst(arr, 8) //查找第一个值等于给定值的元素
	assert.Equal(t, 6, first)
	first2 := BSearchFirst2(arr, 8)
	assert.Equal(t, 6, first2)

	last := BSearchLast(arr, 8) //查找最后一个值等于给定值的元素
	assert.Equal(t, 8, last)

	fge := BSearchFirstGeVal(arr, 8) // 查找第一个大于等于给定值的元素
	assert.Equal(t, 6, fge)

	fg := BSearchFirstGVal(arr, 6) // 查找第一个大于给定值的元素
	assert.Equal(t, 6, fg)

	fg2 := BSearchFirstGVal(arr, 4) // 查找第一个大于给定值的元素
	assert.Equal(t, 3, fg2)

	lle := BSearchLastLeVal(arr, 8) // 查找最后一个小于等于给定值的元素
	assert.Equal(t, 8, lle)

}

func TestBSqrt(t *testing.T) {
	ret := BSqrt(3, 1e-6)
	t.Log(ret)
	ret2 := BSqrt(0.04, 1e-6)
	t.Log(ret2)
}
