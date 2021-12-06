package leetcode

//445. 两数相加 II https://leetcode-cn.com/problems/add-two-numbers-ii/

/*
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
你可以假设除了数字 0 之外，这两个数字都不会以零开头。

输入：l1 = [7,2,4,3], l2 = [5,6,4]
输出：[7,8,0,7]

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[8,0,7]


输入：l1 = [0], l2 = [0]
输出：[0]

链表的长度范围为 [1, 100]
0 <= node.val <= 9
输入数据保证链表代表的数字无前导 0

进阶：如果输入链表不能翻转该如何解决？

思路: 不能反转,就只能借助栈来实现, 将链表的数据丢到stack中,然后一个个弹出来进行计算结果在丢到栈中,最后将结果栈中的数据弹出来形成链表
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbersii(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack1, stack2, resStack []int
	l1c, l2c := l1, l2
	for l1c != nil || l2c != nil {
		if l1c != nil {
			stack1 = append(stack1, l1c.Val)
			l1c = l1c.Next
		}
		if l2c != nil {
			stack2 = append(stack2, l2c.Val)
			l2c = l2c.Next
		}
	}
	i1, i2 := len(stack1)-1, len(stack2)-1
	// 进位记录
	n1, n2, carry := 0, 0, 0
	for i1 >= 0 || i2 >= 0 {

		if i1 >= 0 {
			n1 = stack1[i1]
			i1--
		}
		if i2 >= 0 {
			n2 = stack2[i2]
			i2--
		}
		sum := n1 + n2 + carry
		resStack = append(resStack, sum%10)
		carry = sum / 10
		n1 = 0
		n2 = 0
	}
	if carry > 0 {
		resStack = append(resStack, carry)
	}
	headdummy := &ListNode{}
	// build linkedList
	cur := headdummy
	for i := len(resStack) - 1; i >= 0; i-- {
		tmp := &ListNode{Val: resStack[i]}
		cur.Next = tmp
		cur = tmp
	}
	return headdummy.Next
}

func addTwoNumbersii2(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack1, stack2, resStack []int
	l1c, l2c := l1, l2
	for l1c != nil || l2c != nil {
		if l1c != nil {
			stack1 = append(stack1, l1c.Val)
			l1c = l1c.Next
		}
		if l2c != nil {
			stack2 = append(stack2, l2c.Val)
			l2c = l2c.Next
		}
	}

	// 进位记录
	n1, n2, carry := 0, 0, 0
	for len(stack1) > 0 || len(stack2) > 0 {
		if len(stack1) > 0 {
			n1 = stack1[len(stack1)-1]
			stack1 = stack1[0 : len(stack1)-1]
		}
		if len(stack2) > 0 {
			n2 = stack2[len(stack2)-1]
			stack2 = stack2[0 : len(stack2)-1]
		}
		sum := n1 + n2 + carry
		resStack = append(resStack, sum%10)
		carry = sum / 10
		n1 = 0
		n2 = 0
	}
	if carry > 0 {
		resStack = append(resStack, carry)
	}
	headDummy := &ListNode{}
	// build linkedList
	cur := headDummy
	for i := len(resStack) - 1; i >= 0; i-- {
		tmp := &ListNode{Val: resStack[i]}
		cur.Next = tmp
		cur = tmp
	}
	return headDummy.Next
}

//func Test_addTwoNumbersii2(t *testing.T) {
//	l1 := BuildListNode([]int{2, 4, 3})
//	l2 := BuildListNode([]int{5, 6, 4})
//	listNode := addTwoNumbersii2(l1, l2)
//	t.Log(ListNodeToString(listNode))
//}
