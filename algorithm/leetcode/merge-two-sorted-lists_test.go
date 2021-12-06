package leetcode

// 21. 合并两个有序链表 https://leetcode-cn.com/problems/merge-two-sorted-lists/

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

输入：l1 = [], l2 = []
输出：[]

输入：l1 = [], l2 = [0]
输出：[0]

两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按 非递减顺序 排列

思路:
创建一个新的链表,依次遍历两个链表进行添加到链表中,使用dummy节点简化操作
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyResList := &ListNode{}
	endResList := dummyResList

	cur1, cur2 := list1, list2
	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			endResList.Next = cur1
			cur1 = cur1.Next
		} else {
			endResList.Next = cur2
			cur2 = cur2.Next
		}
		endResList = endResList.Next
	}
	if cur1 != nil {
		endResList.Next = cur1
	}
	if cur2 != nil {
		endResList.Next = cur2
	}
	return dummyResList.Next
}
