package randweighted

import (
	"errors"
	"math/rand"
	"time"
)

/*
游戏中计算权重
*/

// 计算函数
type CalcWeight func() [2]int

/*
权重随机普通算法 O(n) 级别, 只要随机的数量不多可以进行

参数 [0]val,[1]weighted
*/
func WeightedRandom1(vw [][2]int) (CalcWeight, error) {
	if len(vw) == 0 {
		return nil, errors.New("vw is nil")
	}
	// 计算总权总
	sum := 0
	for i := range vw {
		sum += vw[i][1]
	}
	random := rand.New(rand.NewSource(time.Now().Unix()))
	return func() [2]int {
		n := random.Intn(sum)
		for i := range vw {
			if n <= vw[i][1] {
				return vw[i]
			}
			n -= vw[i][1]
		}
		return vw[0]
	}, nil
}

/*
权重随机优化, 先使用计数排序,然后使用二分查找,时间复杂度降低到 O(logn)

参数 [0]val,[1]weighted
*/
func WeightedRandom2(vw [][2]int) (CalcWeight, error) {
	if len(vw) == 0 {
		return nil, errors.New("vw is nil")
	}
	// 计算总权总
	sum := 0
	cntSlice := make([]int, len(vw))
	for i := range vw {
		sum += vw[i][1]
		cntSlice[i] = sum // 计数排序后变成有序的切片
	}
	random := rand.New(rand.NewSource(time.Now().Unix()))
	return func() [2]int {
		n := random.Intn(sum)
		// 二分查找,在cntSlice中查找第一个>=n值大的
		low, high := 0, len(cntSlice)-1
		for low <= high {
			mid := low + ((high - low) >> 1)
			if cntSlice[mid] >= n {
				if mid == 0 || cntSlice[mid-1] < n {
					return vw[mid]
				} else {
					high = mid - 1
				}
			} else if cntSlice[mid] < n {
				low = mid + 1
			}
		}
		return vw[0]

	}, nil

}
