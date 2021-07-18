package hello

import (
	"fmt"
)

// type LinkedListNode struct {
// 	Value int
// 	Next *LinkedListNode
// }

func ReverseLinkedList(n *LinkedListNode) *LinkedListNode {
	var result *LinkedListNode
	for n != nil {
		result = &LinkedListNode{n.Value, result}
		n = n.Next
	}
	return result
}

func printLinkedList(n *LinkedListNode) {
	for n != nil {
		fmt.Println(n.Value)
		n = n.Next
	}
}

func TestReverseLinkedList() {
	printLinkedList(ReverseLinkedList(
		&LinkedListNode{1, &LinkedListNode{2, &LinkedListNode{3, nil}}},
	))
}
