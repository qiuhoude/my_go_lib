package leetcode

// 83. 删除排序链表中的重复元素 https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/

/*
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。
返回同样按升序排列的结果链表。

输入：head = [1,1,2]
输出：[1,2]
  p   c
  1,1,2

输入：head = [1,1,2,3,3]
输出：[1,2,3]

链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序排列

思路:
使用指针记录前一个值的指针,prev.Val==cur.Val 对cur进行删除
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		if prev != nil && cur.Val == prev.Val {
			prev.Next = next
			cur.Next = nil //释放
		} else {
			prev = cur
		}
		cur = next
	}
	return head
}
