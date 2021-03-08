package lru_cache

const (
	// 判断主机是否是64位机器
	// ^uint(0) 32位机器 0XFFFFFFFF , 64位机器 0xFFFFFFFFFFFFFFFF
	// uint64(^uint(0)) 强转成 uint64, 看是否与 ^uint64(0)相等
	hostbit = uint64(^uint(0)) == ^uint64(0)
	LENGTH  = 128
)

/*
和 java 中的 LinkedHashMap 思路基本一致, 都有hash表 和 链表 两种结构,链表保证顺序,hash表保证效率O(1)级别;
添加或获取节点时,将操作的节点移动到链表的尾部(也可以是头部) (操作的节点就时最近使用的节点), 超过容量就移除头部节点(最早使用的那个节点)
*/
type lruNode struct {
	prev *lruNode
	next *lruNode

	key   int // lru key
	value int // lru value

	hnext *lruNode // 拉链
}

type LRUCache struct {
	node []lruNode // hash list

	head *lruNode // lru head node
	tail *lruNode // lru tail node

	capacity int //
	used     int //
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		node:     make([]lruNode, LENGTH),
		head:     nil,
		tail:     nil,
		capacity: capacity,
		used:     0,
	}
}

func (this *LRUCache) Get(key int) int {
	if this.tail == nil {
		return -1
	}

	if tmp := this.searchNode(key); tmp != nil {
		this.moveToTail(tmp)
		return tmp.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 1. 首次插入数据
	// 2. 插入数据不在 LRU 中
	// 3. 插入数据在 LRU 中
	// 4. 插入数据不在 LRU 中, 并且 LRU 已满

	if tmp := this.searchNode(key); tmp != nil {
		tmp.value = value
		this.moveToTail(tmp)
		return
	}
	this.addNode(key, value)

	if this.used > this.capacity {
		this.delNode()
	}
}

func (this *LRUCache) addNode(key int, value int) {
	newNode := &lruNode{
		key:   key,
		value: value,
	}

	tmp := &this.node[hash(key)]
	newNode.hnext = tmp.hnext
	tmp.hnext = newNode
	this.used++

	if this.tail == nil {
		this.tail, this.head = newNode, newNode
		return
	}
	this.tail.next = newNode
	newNode.prev = this.tail
	this.tail = newNode
}

func (this *LRUCache) delNode() {
	if this.head == nil {
		return
	}
	prev := &this.node[hash(this.head.key)]
	tmp := prev.hnext

	for tmp != nil && tmp.key != this.head.key {
		prev = tmp
		tmp = tmp.hnext
	}
	if tmp == nil {
		return
	}
	prev.hnext = tmp.hnext
	this.head = this.head.next
	this.head.prev = nil
	this.used--
}

func (this *LRUCache) searchNode(key int) *lruNode {
	if this.tail == nil {
		return nil
	}

	// 查找
	tmp := this.node[hash(key)].hnext
	for tmp != nil {
		if tmp.key == key {
			return tmp
		}
		tmp = tmp.hnext
	}
	return nil
}

func (this *LRUCache) moveToTail(node *lruNode) {
	if this.tail == node {
		return
	}
	if this.head == node {
		this.head = node.next
		this.head.prev = nil
	} else {
		node.next.prev = node.prev
		node.prev.next = node.next
	}

	node.next = nil
	this.tail.next = node
	node.prev = this.tail

	this.tail = node
}

func hash(key int) int {
	if hostbit { // 64位机器
		return (key ^ (key >> 32)) & (LENGTH - 1)
	}
	return (key ^ (key >> 16)) & (LENGTH - 1) // 与java中HashMap#hash() 基本一致
}
