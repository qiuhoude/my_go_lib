package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

}

//2. 两数相加 https://leetcode-cn.com/problems/add-two-numbers/
/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。


输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

输入：l1 = [0], l2 = [0]
输出：[0]

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]

每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零

思路: 同时遍历两个链表,获取数字并相加,记录进位的数字carry, 下次两数字相加需要加上carry
*/
// 两个数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// 正常情况
	// 2 -> 4 -> 3
	// 5 -> 6 -> 4
	// 7 -> 0 -> 8

	// 长度不相等
	// 4 -> 4
	// 5 -> 6 -> 4
	// 9 -> 0 -> 5

	// 有个链表没有
	//
	// 5 -> 6 -> 4
	// 5 -> 6 -> 4
	dummyHead := &ListNode{}
	cur := dummyHead
	carry := 0 // 进位记录
	var n1, n2, sum int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum = n1 + n2 + carry
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		carry = sum / 10
		n1 = 0
		n2 = 0
		sum = 0
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummyHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 通过切片构建链表
func BuildListNode(arr []int) *ListNode {
	var head, prev *ListNode
	for _, v := range arr {
		cur := &ListNode{Val: v}
		if head == nil {
			head = cur
		}
		if prev != nil {
			prev.Next = cur
		}
		prev = cur
	}
	return head
}
func ListNodeToString(head *ListNode) string {
	var sb strings.Builder
	for cur := head; cur != nil; cur = cur.Next {
		sb.WriteString(fmt.Sprintf("%v->", cur.Val))
	}
	if sb.Len() > 0 {
		sb.WriteString("NULL")
	}
	return sb.String()
}

func TestBuildListNode(t *testing.T) {
	l := BuildListNode([]int{1, 2, 3, 4, 5})
	s := ListNodeToString(l)
	t.Logf("%s", s)
}
