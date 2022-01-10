package leetcode

// 213. 打家劫舍 II https://leetcode-cn.com/problems/house-robber-ii/

/*
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。

输入：nums = [1,2,3,1]
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
     偷窃到的最高金额 = 1 + 3 = 4 。

输入：nums = [0]
输出：0


1 <= nums.length <= 100
0 <= nums[i] <= 1000

思路:
和 198题打家劫舍1 思路与基本一样，因为头尾是相邻的,选择了头就不能选尾部,选择尾部就不能用头部，
可以使用两个dp记录，一个是 dp1(0~n-1), dp2(1~n), 最后比较max(dp1[n-1],dp2[n])
f(0) = nums[0]
f(n) = max(f(n-2)+nums[n],f(n-1))

*/

func robii(nums []int) int {
	n := len(nums)
	maxFn := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	if n < 3 {
		if n == 1 {
			return nums[0]
		}
		if n == 2 {
			return maxFn(nums[0], nums[1])
		}
	}

	dp1 := make([]int, n-1)
	dp2 := make([]int, n-1)
	dp1[0], dp2[0] = nums[0], nums[1]
	dp1[1], dp2[1] = maxFn(nums[0], nums[1]), maxFn(nums[1], nums[2])

	for i := 2; i < n; i++ {
		dp1Index, dp2Index := i, i-1
		if 2 <= dp1Index && dp1Index <= n-2 {
			dp1[dp1Index] = maxFn(dp1[dp1Index-2]+nums[i], dp1[dp1Index-1])
		}
		if 2 <= dp2Index && dp2Index <= n-2 {
			dp2[dp2Index] = maxFn(dp2[dp2Index-2]+nums[i], dp2[dp2Index-1])
		}
	}
	res := maxFn(dp1[n-2], dp2[n-2])
	return res
}

//func Test_robii(t *testing.T) {
//	res := robii([]int{1, 2, 1, 1})
//	t.Logf("%v\n", res)
//}
