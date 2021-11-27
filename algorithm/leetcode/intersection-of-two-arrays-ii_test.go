package leetcode

import (
	"sort"
	"testing"
)

// 350. 两个数组的交集 II https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/

/*
给定两个数组，编写一个函数来计算它们的交集。

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]
说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
我们可以不考虑输出结果的顺序。
进阶：

如果给定的数组已经排好序呢？你将如何优化你的算法？(使用指针技术)
如果 nums1 的大小比 nums2 小很多，哪种方法更优？(hashtable的方式,将nums1放到hashtable中)
如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

思路:
1. 使用hashtable方式,统计nums1元素出现的次数m,
再迭代nums2每个值再hashMap查找如果有值并且出现次数>0就添加到返回列表中并且出现次数减1
2. 和349解题思路一致. 先排序, 使用双指针
*/

func intersectii(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var res []int
	idx1, idx2 := 0, 0
	for idx1 < len(nums1) && idx2 < len(nums2) {
		if nums1[idx1] > nums2[idx2] {
			idx2++
		} else if nums1[idx1] < nums2[idx2] {
			idx1++
		} else {
			tmp := nums1[idx1]
			res = append(res, tmp)
			idx1++
			idx2++
		}
	}
	return res

}

func Test_intersectii(t *testing.T) {
	tests := []struct {
		expected []int
		arg1     []int
		arg2     []int
	}{
		{[]int{2, 2}, []int{1, 2, 2, 1}, []int{2, 2}},
		{[]int{4, 9}, []int{4, 9, 5}, []int{9, 4, 9, 8, 4}},
	}
	for _, tt := range tests {
		res := intersectii(tt.arg1, tt.arg2)
		if !eqSliceInt(res, tt.expected) {
			t.Logf("expected:%v but got:%v", tt.expected, res)
		}
	}
}
