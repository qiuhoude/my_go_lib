package leetcode

import "testing"

// 147. 对链表进行插入排序 https://leetcode-cn.com/problems/insertion-sort-list/

/*
对链表进行插入排序。
插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。

插入排序算法：

插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
重复直到所有输入数据插入完为止。

输入: 4->2->1->3
输出: 1->2->3->4

输入: -1->5->3->4->0
输出: -1->0->3->4->5

思路:
将链表分区:
前部分是已经排好序的,后部分是未排序的 每次从后部分取出一个数放到前部分对应的位置
在插入时需要注意:插入值如果和链表中的值相等需要插入到最后一个相等数后面保证排序的稳定性

*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// beforeList是个双端链表
	dummyBeforeHead := &ListNode{Next: head}
	beforeTail := head // 前部分的的尾部,简化计算
	cur := head.Next
	beforeTail.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = nil                 // 将链表断开成前后部分
		if cur.Val >= beforeTail.Val { // 比前面最后一个还大,说明在最后就可以
			beforeTail.Next = cur
			beforeTail = beforeTail.Next
		} else { //插入到链表中间
			e := dummyBeforeHead.Next
			prev := dummyBeforeHead
			for e != nil {
				if e.Val > cur.Val { //添加到后面 e的前面
					prev.Next = cur
					cur.Next = e
					break
				} else { // e.Val >= cur.Val
					prev = e
					e = e.Next
				}
			}

		}
		cur = next
	}
	return dummyBeforeHead.Next
}

func Test_insertionSortList(t *testing.T) {
	res := insertionSortList(BuildListNode([]int{4, 2, 1, 3}))

	t.Log(ListNodeToString(res))
}
