package leetcode

// 25. K 个一组翻转链表 https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

进阶：
你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。


输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]

输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]

输入：head = [1,2,3,4,5], k = 1
输出：[1,2,3,4,5]

输入：head = [1], k = 1
输出：[1]

列表中节点的数量在范围 sz 内
1 <= sz <= 5000
0 <= Node.val <= 1000
1 <= k <= sz

思路:
k==1 或 head==nil 或 head.Next==nil 都直接返回
创建几个指针
prevGroupTail 上一组的尾节点,初始化是dummy节点, prevGroupTail.Next默认都是nil
curGroupHead 当前组的头节点,初始化是head
cur 当前节点
i 进行计数 i从1开始
当i%k==0 (k>1) 对 curGroupHead到cur之间进行反转反转
最后查看(i-1)%k!=0直接将 prevGroupTail.Next = curGroupHead
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 || head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	prevGroupTail := dummy
	curGroupHead := head
	cur := head
	i := 1
	for cur != nil {
		next := cur.Next
		if i%k == 0 {
			// 反转 curGroupHead~cur的链表
			cur.Next = nil                             //断开链表,方便反转
			newHead := reverseLinkedList(curGroupHead) //反转
			// 此时 curGroupHead 成为尾部
			prevGroupTail.Next = newHead
			prevGroupTail = curGroupHead
			prevGroupTail.Next = nil
			curGroupHead = next
		}
		cur = next
		i++
	}
	if (i-1)%k != 0 {
		prevGroupTail.Next = curGroupHead // 剩余的部分连接起来
	}

	return dummy.Next
}

func reverseLinkedList(head *ListNode) *ListNode {
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

//func Test_reverseKGroup(t *testing.T) {
//	/*
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
//
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]
//
//输入：head = [1,2,3,4,5], k = 1
//输出：[1,2,3,4,5]
//
//输入：head = [1], k = 1
//输出：[1]
//
//	 */
//	arg1 := BuildListNode([]int{1, 2})
//	k := 2
//	res := reverseKGroup(arg1, k)
//	t.Log(ListNodeToString(res))
//
//}
