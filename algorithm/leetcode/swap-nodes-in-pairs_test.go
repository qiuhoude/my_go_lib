package leetcode

//24. 两两交换链表中的节点 https://leetcode-cn.com/problems/swap-nodes-in-pairs/

/*
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

输入：head = [1,2,3,4]
输出：[2,1,4,3]

输入：head = []
输出：[]

输入：head = [1]
输出：[1]

链表中节点的数目在范围 [0, 100] 内
0 <= Node.val <= 100

思路:
p c
  1,2,3,4
理解为创建一个新的链表,把原链表的node添加到新链表中
用i计数,i是偶数添加cur和prev,i是奇数并赋值prev就下一步, 遍历完成后看i是偶数就添加prev到新链表中

思路2:
递归思路, 只要 head和head.next都不为nil,就替换head和head.next的位置,并返回新头节点
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tail := dummyHead // 新链表的尾部
	prev, cur := dummyHead, head
	i := 1 // 计数偶数就添加
	for cur != nil {
		next := cur.Next
		cur.Next = nil
		if i&1 == 0 { //偶数 ,添加 cur 和 prev
			tail.Next = cur
			tail = tail.Next
			tail.Next = prev
			tail = tail.Next
		} else {
			prev = cur
		}
		i++
		cur = next
	}
	if i&1 == 0 { // 偶数
		tail.Next = prev
		tail = tail.Next
	}
	tail.Next = nil
	return dummyHead.Next
}

// 递归的思路
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next //新的头节点
	head.Next = swapPairs2(newHead.Next)
	newHead.Next = head
	return newHead
}

//func Test_swapPairs(t *testing.T) {
//	res := swapPairs2(BuildListNode([]int{1, 2, 3, 4}))
//	t.Log(ListNodeToString(res))
//}
