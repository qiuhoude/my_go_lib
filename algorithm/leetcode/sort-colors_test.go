package leetcode

// 75. 颜色分类 https://leetcode-cn.com/problems/sort-colors/

/*


给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]

输入：nums = [2,0,1]
输出：[0,1,2]

输入：nums = [0]
输出：[0]

输入：nums = [1]
输出：[1]

思路1:
1. 计数排序方式
统计 0,1,2 的个数,然后再填充数组.

思路2:
分区思想,第一波吧0放在最前面,第二波把1放在1后面

思路3:
三路快排,分区3个区域(0,1,2),
p0指向0 ,p2指向最后2,
区域划分
0-> [0,p0)
1-> [p0,i)
2-> [p2,len)

*/

func sortColors(nums []int) {
	if len(nums) < 1 {
		return
	}
	rwb := [3]int{0, 0, 0}
	for i := 0; i < len(nums); i++ {
		rwb[nums[i]]++
	}
	gi := 0
	for n, cnt := range rwb {
		for j := 0; j < cnt; j++ {
			nums[gi] = n
			gi++
		}
	}
}

func sortColors2(nums []int) {
	if len(nums) < 1 {
		return
	}
	gi := 0 // 全局索引
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i], nums[gi] = nums[gi], nums[i]
			gi++
		}
	}
	for i := gi; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[i], nums[gi] = nums[gi], nums[i]
			gi++
		}
	}
}

func sortColors3(nums []int) {
	if len(nums) < 1 {
		return
	}
	p0, p2 := 0, len(nums)-1 // 区域:0[0,p0) 1[p0,i) 2[p2,len)
	for i := 0; i <= p2; {
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			p0++
			i++
		} else if nums[i] == 2 {
			nums[p2], nums[i] = nums[i], nums[p2]
			p2--
		} else {
			i++
		}
	}
}
