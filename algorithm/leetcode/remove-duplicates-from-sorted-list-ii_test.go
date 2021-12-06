package leetcode

// 82. 删除排序链表中的重复元素 II https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
/*
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。
返回同样按升序排列的结果链表。

输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]

输入：head = [1,1,1,2,3]
输出：[2,3]

链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序排列

思路:
和83题思路基本一致, 需要加上dummy节点减少头节点的改变, 83只删除重复的部分,所有head可以永远不改变

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicatesii(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	prev, cur := dummyHead, dummyHead.Next
	for cur != nil {
		next := cur.Next
		if next != nil && cur.Val == next.Val { // 发现重复
			for next != nil && cur.Val == next.Val {
				cur = next
				next = next.Next
			}
			prev.Next = cur.Next
			cur = next
		} else {
			prev = cur
			cur = next
		}
	}
	return dummyHead.Next
}
