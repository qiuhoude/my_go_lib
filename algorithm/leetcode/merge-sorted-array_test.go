package leetcode

// 88. 合并两个有序数组 https://leetcode-cn.com/problems/merge-sorted-array/

/*
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，
其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。

输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]
解释：需要合并 [1] 和 [] 。
合并结果是 [1] 。

输入：nums1 = [0], m = 0, nums2 = [1], n = 1
输出：[1]
解释：需要合并的数组是 [] 和 [1] 。
合并结果是 [1] 。
注意，因为 m = 0 ，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中。

思路:
归并排序的思路,创建temp数组, 将归并的数据放到temp,然后拷贝到num1中

思路2:
逆指针, 因为num2是有全部空间大小并且后半部部分是空的,优化可以不用进行开辟新空间进行处理
*/

func merge2(nums1 []int, m int, nums2 []int, n int) {
	totalLen := m + n
	if totalLen <= 0 {
		return
	}
	p1, p2, pg := m-1, n-1, totalLen-1 // nums1 和 nums2两个数组的指针, pg全局指针
	for ; p1 >= 0 || p2 >= 0; pg-- {
		if p1 == -1 {
			nums1[pg] = nums2[p2]
			p2--
		} else if p2 == -1 {
			nums1[pg] = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			nums1[pg] = nums1[p1]
			p1--
		} else {
			nums1[pg] = nums2[p2]
			p2--
		}
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	totalLen := m + n
	if totalLen <= 0 {
		return
	}
	temp := make([]int, totalLen)

	p1, p2, pg := 0, 0, 0 // nums1 和 nums2两个数组的指针, pg全局指针
	for ; p1 < m && p2 < n; pg++ {
		if nums1[p1] <= nums2[p2] {
			temp[pg] = nums1[p1]
			p1++
		} else {
			temp[pg] = nums2[p2]
			p2++
		}
	}
	// 余下部分
	for ; p1 < m; pg++ {
		temp[pg] = nums1[p1]
		p1++
	}
	for ; p2 < n; pg++ {
		temp[pg] = nums2[p2]
		p2++
	}
	copy(nums1, temp)
}
