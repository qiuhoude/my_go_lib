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
权重随机优化1, 先使用计数排序,然后使用二分查找,时间复杂度降低到 O(logn)

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

/*
权重随机优化2 别名法 Alias method O(1)
解释: https://www.cnblogs.com/Lee-yl/p/12749070.html
步骤:
1. 求出每一项的出现的概率 float64, 将 概率 * N(项的总数) 得到 probes
2. 将 probes 中 大于1的下标添加到large列表中, 小于1的下标添加到small
3. 每次从small,large中各取一个，将大的补充到小的之中，小的出队列，再看大的减去补给之后，如果大于1，
继续放入large中，如果等于1，则也出去，如果小于1则放入small
4. 计算 accept, alias 列表
参数 [0]val,[1]weighted
*/
func WeightedRandom3(vw [][2]int) (CalcWeight, error) {
	n := len(vw)
	if n == 0 {
		return nil, errors.New("vw is nil")
	}
	// 计算总权总
	sum := 0
	for i := range vw {
		sum += vw[i][1]
	}

	ratios := make([]float64, n, n) // 自己在该位置占比
	accept := make([]float64, n, n) // accept存放第i列对应的事件i矩形的面积百分比;
	alias := make([]int, n, n)      // alias存放第i列不是事件i的另外一个事件的标号;

	var small, large []int // 存储 <1, >1的下标
	for i := 0; i < n; i++ {
		accept[i] = 0.0 // 初始化
		alias[i] = -1   // 默认值 -1

		// 分成 small 和 large列表
		ratios[i] = float64(vw[i][1]) / float64(sum) * float64(n)
		if ratios[i] < 1.0 {
			small = append(small, i)
		} else {
			large = append(large, i)
		}
	}

	for len(small) > 0 && len(large) > 0 {
		smallIdx, largeIdx := small[0], large[0]
		small, large = small[1:], large[1:]
		accept[smallIdx] = ratios[smallIdx]
		alias[smallIdx] = largeIdx
		ratios[largeIdx] = ratios[largeIdx] - (1.0 - ratios[smallIdx]) // 将大的补充给小的
		if ratios[largeIdx] < 1.0 {
			small = append(small, largeIdx)
		} else {
			large = append(large, largeIdx)
		}
	}
	// 剩余部分
	for len(small) > 0 {
		idx := small[0]
		small = small[1:]
		accept[idx] = 1
	}
	for len(large) > 0 {
		idx := large[0]
		large = large[1:]
		accept[idx] = 1
	}

	random := rand.New(rand.NewSource(time.Now().Unix()))
	return func() [2]int {
		/*
			随机采样1~cnt 之间的整数i，决定落在哪一列。
			随机采样0~1之间的一个概率值，
			如果小于accept[i]，则采样i，
			如果大于accept[i]，则采样alias[i]；
		*/
		cnt := len(accept)
		index := int(random.Float64() * float64(cnt))
		r := random.Float64()
		if r < accept[index] {
			return vw[index]
		} else {
			return vw[alias[index]]
		}
	}, nil
}
