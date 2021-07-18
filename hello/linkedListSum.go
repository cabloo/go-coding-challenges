package hello

import (
	"fmt"
	"strings"
)

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

func next(n *LinkedListNode) *LinkedListNode {
	if n != nil {
		return n.Next
	}
	return nil
}

func LinkedListSum(a, b *LinkedListNode) (result *LinkedListNode) {
	var curr *LinkedListNode
	var carry int
	for a != nil || b != nil || carry > 0 {
		sum := carry + sumNodes(a, b)
		carry = 0
		if sum >= 10 {
			carry = int(sum / 10)
			sum = sum % 10
		}
		if curr == nil {
			curr = &LinkedListNode{Value: sum}
			result = curr
		} else {
			curr.Next = &LinkedListNode{Value: sum}
			curr = curr.Next
		}
		a, b = next(a), next(b)
	}
	return result
}

func sumNodes(a, b *LinkedListNode) (sum int) {
	if a != nil {
		sum += a.Value
	}
	if b != nil {
		sum += b.Value
	}
	return
}

func (n *LinkedListNode) String() string {
	var b strings.Builder
	for n != nil {
		fmt.Fprintf(&b, "%d -> ", n.Value)
		n = n.Next
	}
	return b.String()[:b.Len()-4]
}

func TestLinkedListSum() {
	fmt.Println(LinkedListSum(&LinkedListNode{
		2, &LinkedListNode{4, &LinkedListNode{Value: 3}},
	}, &LinkedListNode{
		5, &LinkedListNode{6, &LinkedListNode{Value: 4}},
	}))
}
