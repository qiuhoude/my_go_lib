package leetcode

import (
	"github.com/bmizerany/assert"
	"testing"
)

// 55. 跳跃游戏 https://leetcode-cn.com/problems/jump-game/

func canJump(nums []int) bool {
	size := len(nums)
	if size == 0 {
		return false
	}
	// 思路: 遍历异步数组,计算一下当前位置可以跳到的最大位置
	reach := 0
	for i := 0; i < size; i++ {
		if i > reach { // 中途都跨越不了就直接
			return false
		}
		if i+nums[i] > reach { // 更新当前可到达的最大位置
			reach = i + nums[i]
		}
		if reach > size-1 {
			return true
		}
	}
	return true
}

func TestCaCanJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	assert.Equal(t, true, canJump(nums))
	nums = []int{3, 2, 1, 0, 4}
	assert.Equal(t, false, canJump(nums))
	nums = []int{1}
	assert.Equal(t, true, canJump(nums))
	nums = []int{}
	assert.Equal(t, false, canJump(nums))

}
