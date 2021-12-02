package leetcode

import (
	"container/list"
	"sort"
	"testing"
)

// 220. 存在重复元素 III https://leetcode-cn.com/problems/contains-duplicate-iii/

/*
给你一个整数数组 nums 和两个整数 k 和 t 。请你判断是否存在 两个不同下标 i 和 j，
使得 abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k 。
如果存在则返回 true，不存在返回 false。

输入：nums = [1,2,3,1], k = 3, t = 0
输出：true

输入：nums = [1,0,1,1], k = 1, t = 2
输出：true

输入：nums = [1,5,9,1,5,9], k = 2, t = 3
输出：false


0 <= nums.length <= 2 * 104
-2^31 <= nums[i] <= 2^31 - 1
0 <= k <= 1^4
0 <= t <= 2^31 - 1

思路1:
滑动窗口+有序的链表
有序列表存储的是滑动中的所有值,并且是顺序存储,当新值插入进来后只需要计算插入值位置前后两个位置差值的绝对值就可以了

思路2:
滑动窗口 + 对数据分桶
换个问题进行思考,要在一堆数中找到是否有 abs(a,b)<=t的两个数, 可以将这堆数进行分桶装起来,如发现其中一个桶有2个数以上就说明就有,
还需要比较相邻桶的差如有<=t的值
如何分桶 例如 t=3, 表示差3 0号桶 [0,1,2,3] 1号桶 [4,5,6,7] 公式: n/(t+1) t=0公式也适用相当于每个值就是个桶,
如果数字是负值 -1号桶 [-4,-3,-2,-1] -2号桶 [-8,-7,-6,-5]  公式: (n+1)/(t+1)-1
*/
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	// 滑动窗口 + 分桶
	if k < 1 || len(nums) < 2 {
		return false
	}
	bucketIdFunc := func(n int) int { // 计算桶id函数
		if n >= 0 {
			return n / (t + 1)
		} else {
			return (n+1)/(t+1) - 1
		}
	}

	bucket := make(map[int]int) // <桶id,值>

	for i := range nums {
		id := bucketIdFunc(nums[i])
		//fmt.Printf("%v, v=%v t=%v id=%v  bucket=%v\n", i, nums[i], t, id, bucket)
		if _, ok := bucket[id]; ok { // 直接找到了
			return true
		}
		bucket[id] = nums[i]
		// 没找到到相领的桶取查找
		preId, nextId := id-1, id+1
		if v, ok := bucket[preId]; ok && nums[i]-v <= t { // pre 数据肯定比 nums[i]小
			return true
		}
		if v, ok := bucket[nextId]; ok && v-nums[i] <= t { // next 数据肯定比 nums[i]大
			return true
		}
		if i-k >= 0 {
			delete(bucket, bucketIdFunc(nums[i-k]))
		}
	}
	return false

}

func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	// 滑动窗口+有序的链表
	if k < 1 || len(nums) < 2 {
		return false
	}
	n := k + 1
	if len(nums) < n { // 防止k比整个数组len还大
		n = len(nums)
	}
	wArr := make([]int, n)
	copy(wArr, nums[0:n])
	sort.Ints(wArr) // 排序

	// 初始化有序链表
	seqList := list.New()
	for i, v := range wArr {
		seqList.PushBack(wArr[i])
		if i > 0 {
			if int(absI64(int64(wArr[i-1])-int64(v))) <= t { // 提前找到
				return true
			}
		}
	}
	for r := n; r < len(nums); r++ {
		// 移除上一个数的值
		for e := seqList.Front(); e != nil; e = e.Next() {
			if e.Value.(int) == nums[r-n] {
				seqList.Remove(e)
				break
			}
		}

		curV := nums[r] // 当前的数
		seqHeadV := seqList.Front().Value.(int)
		seqTailV := seqList.Back().Value.(int)
		if seqHeadV > curV { // 小于链表头
			if int(absI64(int64(seqHeadV)-int64(curV))) <= t {
				return true
			}
			seqList.PushFront(curV)
		} else if seqTailV < curV { // 大于链表尾
			if int(absI64(int64(seqTailV)-int64(curV))) <= t {
				return true
			}
			seqList.PushBack(curV)
		} else { // 在链表中间
			// 找到第一个大于curV的值,插入到该值的前面
			for e := seqList.Front(); e != nil; e = e.Next() {
				val := e.Value.(int)
				if val >= curV {
					if val == curV && e.Next() != nil && e.Next().Value.(int) == val { // 跳过重复的
						continue
					}
					if int(absI64(int64(val)-int64(curV))) <= t {
						return true
					}
					if e.Prev() != nil {
						valPrev := e.Prev().Value.(int)
						if int(absI64(int64(valPrev)-int64(curV))) <= t {
							return true
						}
					}
					seqList.InsertBefore(curV, e)
					break
				}
			}
		}
	}
	return false
}

func absI64(a int64) int64 {
	if a > 0 {
		return a
	}
	return -a
}

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	tests := []struct {
		arg1 []int
		arg2 int
		arg3 int
		want bool
	}{
		//{[]int{1, 2, 3, 1}, 3, 0, true},
		//{[]int{1, 0, 1, 1}, 1, 2, true},
		//{[]int{1, 5, 9, 1, 5, 9}, 2, 3, false},
		//{[]int{1, 2, 1, 1}, 1, 0, true},
		{[]int{2, 0, -2, 2}, 2, 1, false},
	}
	for _, tt := range tests {
		if got := containsNearbyAlmostDuplicate(tt.arg1, tt.arg2, tt.arg3); got != tt.want {
			t.Errorf("arg1(%v) got=%v  want=%v", tt.arg1, got, tt.want)
		}
	}
}
