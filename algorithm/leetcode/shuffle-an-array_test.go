package leetcode

import (
	"math/rand"
	"testing"
)

// 384. 打乱数组 https://leetcode-cn.com/problems/shuffle-an-array/

type Solution struct {
	origin []int
}

func ConstructorSolution(nums []int) Solution {
	o := make([]int, len(nums))
	copy(o, nums)
	return Solution{origin: o}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.origin
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	size := len(this.origin)
	res := make([]int, size)
	copy(res, this.origin)
	for i := size; i > 1; i-- {
		r := rand.Intn(i)
		res[i-1], res[r] = res[r], res[i-1]
	}
	return res
}

func TestShuffle_an_array(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	obj := ConstructorSolution(nums)
	param_1 := obj.Reset()
	param_2 := obj.Shuffle()
	t.Logf("param_1:%v, param_2:%v", param_1, param_2)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
