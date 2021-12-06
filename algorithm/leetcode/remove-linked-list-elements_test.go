package leetcode

// 203. 移除链表元素 https://leetcode-cn.com/problems/remove-linked-list-elements/
/*
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]

输入：head = [], val = 1
输出：[]

输入：head = [7,7,7,7], val = 7
输出：[]

列表中的节点数目在范围 [0, 104] 内
1 <= Node.val <= 50
0 <= val <= 50

思路:
没什么难点,用个dummy节点简化nil判断
2. 递归的写法

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}

	prev := dummy
	cur := dummy.Next
	for cur != nil {
		next := cur.Next
		if cur.Val == val {
			prev.Next = next
			cur.Next = nil //gc
		} else {
			prev = cur
		}
		cur = next
	}

	return dummy.Next
}

// 递归的方式删除, 如自身是该值就返自己的下一个节点
func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements2(head.Next, val)
	if head.Val == val {
		return head.Next
	} else {
		return head
	}

}
