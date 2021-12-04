package leetcode

// 206. 反转链表 https://leetcode-cn.com/problems/reverse-linked-list/

/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

输入：head = [1,2]
输出：[2,1]

输入：head = []
输出：[]


链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000

思路: 遍历链表时使用一个指针记录前节点,每次在遍历都都将当前next指向上一个node
 p c
[1,2,3,4,5]
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	/*if head == nil || head.Next == nil {
		return head
	}
	prev := head
	cur := prev.Next
	prev.Next = nil
	for cur != nil {
		tmp := cur
		cur = cur.Next
		tmp.Next = prev
		prev = tmp
	}
	return prev*/

	var prev *ListNode = nil
	var cur = head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev

}
