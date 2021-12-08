package leetcode

// 143. 重排链表 https://leetcode-cn.com/problems/reorder-list/

/*
给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

输入：head = [1,2,3,4]
输出：[1,4,2,3]

输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]

链表的长度范围为 [1, 5 * 10^4]
1 <= node.val <= 1000

思路:
快慢指针找出链表的中间位置, 将原链表分成前后两部分,将后半部分进行反转,然后与前半部分进行合并
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// 将链表一分成两部分
	slow, fast := head, head
	var prev1 *ListNode
	for fast != nil && fast.Next != nil {
		prev1 = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev1.Next = nil
	// 反转后半部分
	var prev2 *ListNode
	for cur := slow; cur != nil; {
		next := cur.Next
		cur.Next = prev2
		prev2 = cur
		cur = next
	}
	head2 := prev2
	// 将 head2 合并到 head 中
	cur1, cur2 := head, head2
	for cur1 != nil && cur1.Next != nil && cur2 != nil {
		next1 := cur1.Next
		next2 := cur2.Next
		cur2.Next = cur1.Next
		cur1.Next = cur2
		cur1 = next1
		cur2 = next2
	}
	if cur2 != nil {
		cur1.Next = cur2
	}
}

//func Test_reorderList(t *testing.T) {
//	head := BuildListNode([]int{1, 2, 3, 4,5})
//	reorderList(head)
//	t.Log(ListNodeToString(head))
//}
