// Package bst provides BinaryTree Node & Tree implementation
package bst

import (
	"fmt"
	"strconv"
)

type Node struct {
	Value  uint
	Parent *Node
	Left   *Node
	Right  *Node
}

func NewNode(val uint) *Node {
	return &Node{Value: val}
}

func (n Node) Compare(other Node) int {
	if n.Value < other.Value {
		return -1
	} else if n.Value > other.Value {
		return 1
	}
	return 0
}

func (n Node) String() string {
	var left, right string
	if n.Left != nil {
		left = strconv.Itoa(int(n.Left.Value))
	}
	if n.Right != nil {
		right = strconv.Itoa(int(n.Right.Value))
	}
	return fmt.Sprintf("value=%d left=%s right=%s parent=%d \n", n.Value, left, right, n.Parent.Value)
}

type Tree struct {
	Head *Node
	Size uint
}

func NewTree(n *Node) *Tree {
	return &Tree{Head: n, Size: 1}
}

func (t *Tree) Insert(val uint) {
	head := t.Head
	for {
		if head.Value < val && head.Right != nil {
			head = head.Right
		} else if head.Value > val && head.Left != nil {
			head = head.Left
		} else {
			break
		}
	}

	if head.Value == val {
		return
	}

	newNode := &Node{Value: val, Parent: head}
	if head.Value < val {
		head.Right = newNode
	} else {
		head.Left = newNode
	}
	t.Size++
}

func (t *Tree) Search(val uint) *Node {
	n := t.Head
	for {
		if n.Value < val && n.Right != nil {
			n = n.Right
		} else if n.Value > val && n.Left != nil {
			n = n.Left
		} else {
			break
		}
	}
	if n.Value == val {
		return n
	}
	return nil
}

func (t *Tree) Delete(val uint) {
	n := t.Search(val)
	if n == nil {
		return
	}

	parent := n.Parent
	var replaceNode *Node

	if n.Left != nil {
		replaceNode = n.Left
	} else if n.Right != nil {
		replaceNode = n.Right
	}

	// parent & replace node rearrangement
	replaceNode.Parent = parent
	if parent.Right.Compare(*n) == 0 {
		parent.Right = replaceNode
	} else {
		parent.Left = replaceNode
	}
	t.Size--

	// right node rearrangement
	rightNode := n.Right
	if rightNode == nil {
		return
	}
	head := replaceNode
	for {
		// only need to lookup for right branch
		if head.Right != nil {
			head = head.Right
		} else {
			break
		}
	}
	head.Right = rightNode
	rightNode.Parent = head
}
