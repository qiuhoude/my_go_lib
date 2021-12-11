package leetcode

// 341. 扁平化嵌套列表迭代器 https://leetcode-cn.com/problems/flatten-nested-list-iterator/

/*
给你一个嵌套的整数列表 nestedList 。每个元素要么是一个整数，要么是一个列表；该列表的元素也可能是整数或者是其他列表。请你实现一个迭代器将其扁平化，使之能够遍历这个列表中的所有整数。

实现扁平迭代器类 NestedIterator ：

NestedIterator(List<NestedInteger> nestedList) 用嵌套列表 nestedList 初始化迭代器。
int next() 返回嵌套列表的下一个整数。
boolean hasNext() 如果仍然存在待迭代的整数，返回 true ；否则，返回 false 。
你的代码将会用下述伪代码检测：

initialize iterator with nestedList
res = []
while iterator.hasNext()
    append iterator.next() to the end of res
return res
如果 res 与预期的扁平化列表匹配，那么你的代码将会被判为正确。

输入：nestedList = [[1,1],2,[1,1]]
输出：[1,1,2,1,1]
解释：通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,1,2,1,1]。

输入：nestedList = [1,[4,[6]]]
输出：[1,4,6]
解释：通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,4,6]。

1 <= nestedList.length <= 500
嵌套列表中的整数值在范围 [-106, 106] 内

思路:
1. 使用深度优先将数据放到切片中
2. 使用栈模拟dfs,将数据放到栈中
*/

type NestedInteger struct {
	val  int
	list []*NestedInteger
}

func (this NestedInteger) IsInteger() bool           { return len(this.list) == 0 }
func (this NestedInteger) GetInteger() int           { return this.val }
func (this *NestedInteger) SetInteger(value int)     { this.val = value }
func (this *NestedInteger) Add(elem NestedInteger)   { this.list = append(this.list, &elem) }
func (this NestedInteger) GetList() []*NestedInteger { return this.list }

/*
type NestedIterator struct {
	arr []int
}

func ConstructorNestedIterator(nestedList []*NestedInteger) *NestedIterator {
	res := &NestedIterator{}
	res.dfs(nestedList)
	return res
}

func (this *NestedIterator) dfs(nestedList []*NestedInteger) {
	for _, v := range nestedList {
		if v.IsInteger() {
			this.arr = append(this.arr, v.GetInteger())
		} else {
			this.dfs(v.GetList())
		}
	}
}

func (this *NestedIterator) Next() int {
	res := this.arr[0]
	this.arr = this.arr[1:]
	return res
}

func (this *NestedIterator) HasNext() bool {
	return len(this.arr) > 0
}*/

type NestedIterator struct {
	stack [][]*NestedInteger // 第一层是个栈,第二层是个队列
}

func ConstructorNestedIterator(nestedList []*NestedInteger) *NestedIterator {
	res := &NestedIterator{stack: [][]*NestedInteger{nestedList}}
	res.next()
	return res
}

func (this *NestedIterator) next() {
	// 每次确保stack栈顶的que第一个值是integer值
	for len(this.stack) > 0 {
		que := this.stack[len(this.stack)-1]
		if len(que) == 0 { // 空的队列直pop
			this.stack = this.stack[:len(this.stack)-1]
			continue
		}
		nested := que[0]
		if nested.IsInteger() { // que第一个值是integer值,推出循环
			break
		}
		//不是integer就加到stack中,注意需要 nested 从队列中取出,避免重复添加
		this.stack[len(this.stack)-1] = que[1:]
		this.stack = append(this.stack, nested.GetList())
	}
}

func (this *NestedIterator) Next() int {
	defer this.next()
	val := this.stack[len(this.stack)-1][0].GetInteger()
	this.stack[len(this.stack)-1] = this.stack[len(this.stack)-1][1:]
	return val
}

func (this *NestedIterator) HasNext() bool {
	return len(this.stack) > 0
}
