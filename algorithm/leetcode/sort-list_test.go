package leetcode

//148. 排序链表 https://leetcode-cn.com/problems/sort-list/

/*
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

进阶：
你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？


输入：head = [4,2,1,3]
输出：[1,2,3,4]


输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]

输入：head = []
输出：[]

链表中节点的数目在范围 [0, 5 * 104] 内
-105 <= Node.val <= 105

思路:
自顶向下:使用归并排序,将原链表切分成各个小链表,切分不能切分为止,然后进行合并

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 使用快慢指针对节点进行分割
	slow, fast := head, head
	var prev *ListNode // 慢指针的前一个节点
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil // 断开链表 ,此处list中必然有两个元素,所以prev必然会有值
	l1 := sortList(head)
	l2 := sortList(slow)
	return mergeNodeList(l1, l2) // 进行归并两个链表
}

func mergeNodeList(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	cur1, cur2 := l1, l2
	for cur1 != nil && cur2 != nil {
		if cur1.Val >= cur2.Val {
			cur.Next = cur2
			cur2 = cur2.Next
		} else {
			cur.Next = cur1
			cur1 = cur1.Next
		}
		cur = cur.Next
	}
	// 余下部分

	if cur1 != nil {
		cur.Next = cur1
	} else if cur2 != nil {
		cur.Next = cur2
	}
	return dummyHead.Next
}
