package leetcode

// 61. 旋转链表 https://leetcode-cn.com/problems/rotate-list/

/*
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]

输入：head = [0,1,2], k = 4
输出：[2,0,1]

链表中节点的数目在范围 [0, 500] 内
-100 <= Node.val <= 100
0 <= k <= 2 * 10^9

思路:
先求出链表的长度size,然后 使用 k%len=realStep获取真实的步数,然后使用快慢指针进行移动,将慢指针放在头部,快指针的尾部连上原来的头部
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	size := 0
	for cur := head; cur != nil; cur = cur.Next {
		size++
	}
	realStep := k % size
	if realStep == 0 { // 不需要移动直接返回
		return head
	}
	dummyHead := &ListNode{Next: head}
	slow, fast := dummyHead, dummyHead
	for fast != nil && fast.Next != nil {
		if realStep == 0 {
			slow = slow.Next
		} else {
			realStep--
		}
		fast = fast.Next
	}

	dummyHead.Next = slow.Next
	fast.Next = head
	slow.Next = nil
	return dummyHead.Next
}
