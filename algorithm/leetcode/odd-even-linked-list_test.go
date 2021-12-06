package leetcode

//328. 奇偶链表 https://leetcode-cn.com/problems/odd-even-linked-list/

/*
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL

输入: 2->1->3->5->6->4->7->NULL
输出: 2->3->6->7->1->5->4->NULL

应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。

思路: 与86题思路基本一致, 就是将原链表分成2个链表,然后将 oddEnd.Next->evendummy.Next,创建dummy有助于减少nil判断的代码
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	oddDummy, evenDummy := &ListNode{}, &ListNode{}
	oddEnd, evenEnd := oddDummy, evenDummy

	cur := head
	for i := 1; cur != nil; i++ {
		next := cur.Next
		cur.Next = nil
		if i&1 == 0 { //even
			evenEnd.Next = cur
			evenEnd = cur
		} else {
			oddEnd.Next = cur
			oddEnd = cur
		}
		cur = next
	}
	oddEnd.Next = evenDummy.Next
	return oddDummy.Next
}
