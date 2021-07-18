package hello

import (
	"fmt"
	"math"
)

// type BinaryTreeNode struct {
// 	Value int
// 	Left  *BinaryTreeNode
// 	Right *BinaryTreeNode
// }

func (n BinaryTreeNode) IsValidBST() bool {
	return isValidBST(n.Left, math.Inf(-1), float64(n.Value)) &&
		isValidBST(n.Right, float64(n.Value), math.Inf(1))
}

func isValidBST(n *BinaryTreeNode, low, high float64) bool {
	if n == nil {
		return true
	}
	if float64(n.Value) >= high || float64(n.Value) <= low {
		return false
	}
	return isValidBST(n.Left, low, float64(n.Value)) &&
		isValidBST(n.Right, float64(n.Value), high)
}

func TestValidBST() {
	n := BinaryTreeNode{
		7,
		&BinaryTreeNode{3, &BinaryTreeNode{Value: 2}, &BinaryTreeNode{Value: 4}},
		&BinaryTreeNode{9, &BinaryTreeNode{Value: 8}, &BinaryTreeNode{Value: 10}},
	}
	fmt.Println(n, n.IsValidBST())
	n = BinaryTreeNode{
		7,
		&BinaryTreeNode{3, &BinaryTreeNode{Value: 2}, &BinaryTreeNode{Value: 8}},
		&BinaryTreeNode{9, &BinaryTreeNode{Value: 8}, &BinaryTreeNode{Value: 10}},
	}
	fmt.Println(n, n.IsValidBST())
	n = BinaryTreeNode{
		7,
		&BinaryTreeNode{3, &BinaryTreeNode{Value: 2}, &BinaryTreeNode{Value: 4}},
		&BinaryTreeNode{9, &BinaryTreeNode{Value: 5}, &BinaryTreeNode{Value: 10}},
	}
	fmt.Println(n, n.IsValidBST())
}
