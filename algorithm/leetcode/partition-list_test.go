package leetcode

import "testing"

// 86. 分隔链表 https://leetcode-cn.com/problems/partition-list/

/*
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
你应当 保留 两个分区中每个节点的初始相对位置。

输入：head = [1,4,3,2,5,2], x = 3
输出：[1,2,2,4,3,5]

输入：head = [2,1], x = 2
输出：[1,2]


链表中节点的数目在范围 [0, 200] 内
-100 <= Node.val <= 100
-200 <= x <= 200

思路:
将原来的链表分成 smallLinked, bigLinked 两个个链表,链表头都是用dummy节点, 最后smallLinked.end.next -> bigLinked.dummy.next
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partitionLinkedList(head *ListNode, x int) *ListNode {
	// 将原来的链表分成 smallLinked, bigLinked 两个个链表,链表头都是用dummy节点, 最后smallLinked.end.next -> bigLinked.dummy.next
	smallDummy, bigDummy := &ListNode{}, &ListNode{}
	var smallEnd, bigEnd = smallDummy, bigDummy

	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = nil
		if cur.Val >= x {
			// big
			bigEnd.Next = cur
			bigEnd = cur

		} else { // < x
			// small
			smallEnd.Next = cur
			smallEnd = cur
		}
		cur = next
	}
	smallEnd.Next = bigDummy.Next
	return smallDummy.Next
}

func Test_partitionLinkedList(t *testing.T) {
	/*
		输入：head = [1,4,3,2,5,2], x = 3
		输出：[1,2,2,4,3,5]
	*/
	l := BuildListNode([]int{1, 4, 3, 2, 5, 2})
	list := partitionLinkedList(l, 3)
	t.Log(ListNodeToString(list))
}
