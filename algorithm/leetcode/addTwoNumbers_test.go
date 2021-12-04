package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

}

//https://leetcode-cn.com/problems/add-two-numbers
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
	carry := 0
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
