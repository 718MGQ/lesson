package array

import (
	"strconv"
)

type Node struct {
	Val  int
	Next *Node
}

func NewNode(val int) *Node {
	return &Node{Val: val}
}

func (n *Node) String() string {
	return strconv.Itoa(n.Val)
}

func QueryNodeList(head *Node) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}

// InsertNode 在index节点后面插入node
func InsertNode(head *Node, index int, node *Node) {
	if node == nil || head == nil {
		return
	}

	for i := 0; i < index && head != nil; i++ {
		head = head.Next
	}

	// 说明index超过范围了
	if head == nil {
		return
	}

	tmp := head.Next
	head.Next = node
	node.Next = tmp
}

func DeleteNode(head *Node, index int, node *Node) {
	if node == nil || head == nil {
		return
	}

	for i := 0; i < index && head != nil; i++ {
		head = head.Next
	}

	if head == nil || head.Next == nil {
		return
	}

	tmp := head.Next
	head.Next = tmp.Next
	tmp.Next = nil
}
