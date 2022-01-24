package leetcode

// 1049. 最后一块石头的重量 II https://leetcode-cn.com/problems/last-stone-weight-ii/

/*
有一堆石头，用整数数组 stones 表示。其中 stones[i] 表示第 i 块石头的重量。

每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：

如果 x == y，那么两块石头都会被完全粉碎；
如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。

输入：stones = [2,7,4,1,8,1]
输出：1
解释：
组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]，
组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]，
组合 2 和 1，得到 1，所以数组转化为 [1,1,1]，
组合 1 和 1，得到 0，所以数组转化为 [1]，这就是最优值。

输入：stones = [31,26,33,21,40]
输出：5

输入：stones = [1,2]
输出：1

1 <= stones.length <= 30
1 <= stones[i] <= 100

思路:
 和 416. 分割等和子集 思路很相近, 难度在将问题转换成 0-1背包问题上

转换成1 => 将一堆石头分成两部分,求解两堆时候的差值最小
转换成2 => 差值最小就是两堆石头谁更接近接近于 sum(store)/2, A,B两堆石头
转换成3 => 将石头放进 小为sum(store)/2的背包, 使其重量最大
那么剩余最小必然是 sum(store) - 2*maxWeight

*/

func lastStoneWeightII(stones []int) int {
	maxFn := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(stones)
	sum := 0
	for _, v := range stones {
		sum += v
	}
	capacity := sum / 2 //背包容量

	//dp := make([][]int, n)
	//for i := range dp {
	//	dp[i] = make([]int, capacity+1)
	//}
	//第一排
	//for j := stones[0]; j <= capacity; j++ {
	//	dp[0][j] = stones[0]
	//}
	//
	//for i := 1; i < n; i++ {
	//	for j := 0; j <= capacity; j++ {
	//		if j-stones[i] < 0 {
	//			dp[i][j] = dp[i-1][j] // 当前石头重量超过最大容量,肯定不能选择石头i, 照抄上一行
	//		} else {
	//			// 选择i 不选择i 进行选择,看谁的重量大
	//			dp[i][j] = maxFn(dp[i-1][j], dp[i-1][j-stones[i]]+stones[i])
	//		}
	//	}
	//}
	//maxWeight := dp[n-1][capacity]

	// 优化成单数组
	dp := make([]int, capacity+1)
	dp[0] = 0
	for i := 0; i < n; i++ {
		for j := capacity; j >= stones[i]; j-- {
			// 选择i 不选择i 进行选择,看谁的重量大
			dp[j] = maxFn(dp[j-stones[i]]+stones[i], dp[j])
		}
	}
	maxWeight := dp[capacity]

	return sum - 2*maxWeight
}

//func Test_lastStoneWeightII(t *testing.T) {
//	res := lastStoneWeightII([]int{21, 60, 61, 20, 31})
//	//res := lastStoneWeightII([]int{1, 2})
//	t.Logf("%v\n", res)
//}
