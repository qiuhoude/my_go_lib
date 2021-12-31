package leetcode

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// 297 二叉树的序列化与反序列化 https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func ConstructorCodec() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	var que []*TreeNode
	que = append(que, root)
	var sb strings.Builder
	sb.WriteString("[")
	sb.WriteString(strconv.Itoa(root.Val) + ",")
	for len(que) != 0 {
		n := que[0]
		que = que[1:]
		if n.Left != nil {
			que = append(que, n.Left)
			sb.WriteString(fmt.Sprintf("%d,", n.Left.Val))
		} else {
			sb.WriteString("nil,")
		}
		if n.Right != nil {
			que = append(que, n.Right)
			sb.WriteString(fmt.Sprintf("%d,", n.Right.Val))
		} else {
			sb.WriteString("nil,")
		}
	}
	s := sb.String()
	s = s[0 : len(s)-1]
	return s + "]"
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "[]" {
		return nil
	}
	data = data[1 : len(data)-1]
	arr := strings.Split(data, ",")
	var que []*TreeNode
	rootVal, _ := strconv.Atoi(arr[0])
	root := &TreeNode{Val: rootVal}
	que = append(que, root)

	idx := 0
	for len(que) != 0 && idx < len(arr)-2 {
		n := que[0]
		que = que[1:]
		idx++
		left := arr[idx]
		if left != "nil" {
			val, _ := strconv.Atoi(left)
			n.Left = &TreeNode{Val: val}
			que = append(que, n.Left)
		}
		idx++
		right := arr[idx]
		if right != "nil" {
			val, _ := strconv.Atoi(right)
			n.Right = &TreeNode{Val: val}
			que = append(que, n.Right)
		}
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

func TestName(t *testing.T) {
	//str := "1,2,3,4,"
	//split := strings.Split(str, ",")
	//for i := range split {
	//	t.Logf("%v", split[i])
	//}
	codec := ConstructorCodec()
	nStr := "[1,2,3,nil,nil,4,5,nil,nil,nil,nil]"
	root := codec.deserialize(nStr)
	outStr := codec.serialize(root)
	t.Log(outStr)
}
