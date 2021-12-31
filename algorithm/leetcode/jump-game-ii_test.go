package leetcode

import (
	"github.com/bmizerany/assert"
	"testing"
)

// 45. 跳跃游戏 II https://leetcode-cn.com/problems/jump-game-ii/

func jump(nums []int) int {
	// 思路: 每步找到可走下一步的最大数,
	// 找出当前位置可到达的最大值
	size := len(nums)
	step := 0
	maxPos := 0
	end := 0
	for i := 0; i < size-1; i++ {
		if i+nums[i] > maxPos {
			maxPos = i + nums[i]
		}
		if i == end { // 当前能到达 到前一次最大值,说明此步可以走
			step++
			end = maxPos
		}
	}
	return step
}

func TestJump(t *testing.T) {
	nums := []int{2, 3, 0, 1, 4}
	assert.Equal(t, 2, jump(nums))

	nums = []int{2, 0, 0, 4}
	t.Log(jump(nums))
	//assert.Equal(t, 0, jump(nums))
	//
	//nums = []int{1}
	//assert.Equal(t, 1, jump(nums))
}
