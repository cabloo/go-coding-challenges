package hello

import "fmt"

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func (tree *BinaryTreeNode) LeafValueSequence() []int {
	if tree.Left == nil && tree.Right == nil {
		return []int{tree.Value}
	}

	leftDone := make(chan bool)
	var left []int
	var right []int

	go func() {
		if tree.Left != nil {
			left = tree.Left.LeafValueSequence()
		}
		leftDone <- true
	}()

	if tree.Right != nil {
		right = tree.Right.LeafValueSequence()
	}

	<-leftDone

	return append(left, right...)
}

func TestLeafValueSequence() {
	fmt.Println((&BinaryTreeNode{
		Left:  &BinaryTreeNode{Left: &BinaryTreeNode{Value: 1}, Right: &BinaryTreeNode{Value: 2}},
		Right: &BinaryTreeNode{Right: &BinaryTreeNode{Value: 4}},
	}).LeafValueSequence())
}
