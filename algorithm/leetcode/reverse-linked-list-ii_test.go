package leetcode

// 92. 反转链表 II https://leetcode-cn.com/problems/reverse-linked-list-ii/

/*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

输入：head = [5], left = 1, right = 1
输出：[5]

链表中节点数目为 n
1 <= n <= 500
-500 <= Node.val <= 500
1 <= left <= right <= n

进阶： 你可以使用一趟扫描完成反转吗？

思路: 和206反转链表基本一致,需要新增leftPrev和leftNode节点的指针,用与结束后进行链接, 创建dummy哑节点可以简化操作和边界判断

  2~4
  p c n
  d,1,2,3,4,5
  d,1,4,3,2,5

  p c n
  d,1,3
  d,3,1
*/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	dummy := &ListNode{Next: head}   // 创建dummy节点操作简化边界条件判断
	var leftPrev, leftNode *ListNode // 记录开始的节点和开始的前一个节点
	var prev *ListNode = nil
	cur := dummy
	for i := 0; cur != nil; i++ { // left right下标是从1开始,创建哑节点后可以从0开始
		next := cur.Next
		if i == left {
			leftPrev = prev
			leftNode = cur
		}
		if left < i && i <= right {
			cur.Next = prev
		}
		prev = cur
		cur = next

		if i == right { // 到结束
			leftPrev.Next = prev
			leftNode.Next = next
			break
		}
	}
	return dummy.Next
}
