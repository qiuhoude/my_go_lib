package leetcode

//1. 两数之和 https://leetcode-cn.com/problems/two-sum/

/*
思路1:
使用查表法,将目标减去cur的值,看cur之前的数是否再hash表中存在

思路2:
排序后使用对撞指针,注意排序是下标进行排序
*/
func twoSum(nums []int, target int) []int {
	length := len(nums)
	res := make([]int, 2)
	tmpMap := make(map[int]int)

	var t int
	for i := 0; i < length; i++ {
		t = target - nums[i]
		if v, ok := tmpMap[t]; ok {
			res[0] = i
			res[1] = v
			return res
		}
		tmpMap[nums[i]] = i
	}
	panic("没有找到")
}
