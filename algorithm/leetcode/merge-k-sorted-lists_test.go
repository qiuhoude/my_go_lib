package leetcode

import "testing"

// 23. 合并K个升序链表 https://leetcode-cn.com/problems/merge-k-sorted-lists/

/*
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6

输入：lists = []
输出：[]

输入：lists = [[]]
输出：[]


k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4

思路:
将list放到队列中,每次取两个出来进行合并,让后在方回队列中 2变1, 直到队列中只有一个元素就返回

思路2:
分治思想,将链表数组进行一份为2,只有列表只有2个元素时才进行合并
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		l1, l2 := lists[0], lists[1]
		lists = lists[2:]
		lists = append(lists, merge2List(l1, l2))
	}
	return lists[0]
}

func mergeKLists(lists []*ListNode) *ListNode {
	size := len(lists)
	if size == 0 {
		return nil
	} else if size == 1 {
		return lists[0]
	} else if size == 2 {
		return merge2List(lists[0], lists[1])
	} else {
		middle := size / 2
		l1 := mergeKLists(lists[0:middle])
		l2 := mergeKLists(lists[middle:size])
		return merge2List(l1, l2)
	}
}

func merge2List(l1, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	cur1, cur2 := l1, l2
	for cur1 != nil && cur2 != nil {
		if cur1.Val > cur2.Val {
			cur.Next = cur2
			cur2 = cur2.Next
		} else {
			cur.Next = cur1
			cur1 = cur1.Next
		}
		cur = cur.Next
	}
	if cur1 != nil {
		cur.Next = cur1
	} else if cur2 != nil {
		cur.Next = cur2
	}
	return dummyHead.Next
}

func Test_mergeKLists(t *testing.T) {

}
