package minheap

/*
索引堆:
在堆的基础上加上了索引和反向索引,改变只是索引和反向索引进行修改

反向索引的特点:
indexes[i] = j
reverse[j] = i
indexes[reverse[i]] = i
reverse[indexes[i]] = i
*/

type IndexHeap struct {
	data    []interface{} //存储的数据
	cmpFunc CompareFunc   // 比较函数
	indexes []int         // 正向索引
	reverse []int         // 反向索引
}

func NewIndexHeap(f CompareFunc) *IndexHeap {
	return &IndexHeap{cmpFunc: f}
}

func (h *IndexHeap) Heapify(d ...interface{}) {
	// 思路: 跳过叶子节点,对最小的父节点进行下沉操作,一直到根部
	// 最小的叶子节点的服节点就 parent(len()-1)
	if d == nil || len(d) == 0 {
		return
	}
	n := len(d)
	h.data = make([]interface{}, n)
	h.indexes = make([]int, n)
	h.reverse = make([]int, n)
	copy(h.data, d)
	// 初始化
	for i := 0; i < n; i++ {
		h.indexes[i] = i
		h.reverse[i] = i
	}

	for i := parent(h.Len() - 1); i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *IndexHeap) Len() int {
	if h.data == nil {
		return 0
	}
	return len(h.data)
}

func (h *IndexHeap) Poll() interface{} {
	//1. 取出队头元素
	//2. 将堆尾元素,移到顶部
	//3. 移除尾部
	//4. 对头部下沉
	if h.Len() == 0 {
		return nil
	}
	maxIndex := h.Len() - 1
	ret := h.Peek()

	h.swapData(h.indexes[0], h.indexes[maxIndex])
	h.swapIndex(0, maxIndex)
	h.siftDown(0)
	return ret
}

func (h *IndexHeap) Add(e interface{}) {
	h.data = append(h.data, e)
	h.indexes = append(h.indexes, h.Len()-1)
	h.reverse = append(h.reverse, h.Len()-1)
	h.siftUp(h.Len() - 1)
}

func (h *IndexHeap) Peek() interface{} {
	if h.Len() == 0 {
		return nil
	}
	return h.data[h.indexes[0]]
}

// 上浮
func (h *IndexHeap) siftUp(i int) {
	ci := i
	pi := parent(ci)
	for ci > 0 && h.cmpFunc(h.data[h.indexes[ci]], h.data[h.indexes[pi]]) < 0 {
		h.swapIndex(pi, ci)
		ci = pi
		pi = parent(ci)
	}
}

// 下沉
func (h *IndexHeap) siftDown(i int) {
	ci := i
	dataLen := h.Len() // 数据大小
	for leftChild(ci) < dataLen {
		mi := leftChild(ci)                                                                  //  较小值的孩子的下标
		if mi+1 < dataLen && h.cmpFunc(h.data[h.indexes[mi]], h.data[h.indexes[mi+1]]) > 0 { // mi + 1 表示右边下标
			// 右孩子的值小些
			mi += 1
		}
		if h.cmpFunc(h.data[h.indexes[ci]], h.data[h.indexes[mi]]) <= 0 {
			// 已经比孩子小了不用下沉
			break
		}
		h.swapIndex(mi, ci)
		ci = mi
	}
}

func (h *IndexHeap) swapIndex(i, j int) {
	/*
		indexes[i] = j
		reverse[j] = i
		indexes[reverse[i]] = i
		reverse[indexes[i]] = i
	*/
	//h.indexes[i], h.indexes[j] = h.indexes[j], h.indexes[i]
	tmp := h.indexes[i]
	h.indexes[i] = h.indexes[j]
	h.reverse[h.indexes[j]] = i
	h.indexes[j] = tmp
	h.reverse[tmp] = j
}

func (h *IndexHeap) swapData(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
