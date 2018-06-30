package bst

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	n := NewNode(1)
	m := NewNode(2)
	if n.Compare(*m) != -1 || m.Compare(*n) != 1 || n.Compare(*n) != 0 {
		fmt.Printf("n.Compare(n) = %+v\n", n.Compare(*n))
		t.Error()
	}
}

func TestTree(t *testing.T) {
	n := NewNode(1)
	tree := NewTree(n)

	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(6)

	if tree.Size != 6 {
		t.Errorf("Tree.Size = %d, expect %d", tree.Size, 6)
	}

	five := tree.Search(5)
	if five == nil || five.Value != 5 || five.Parent.Value != 4 || five.Left != nil || five.Right.Value != 6 {
		t.Errorf("tree.Search 5, got %#v", five)
	}

	// // after delete 5
	tree.Delete(5)
	if tree.Size != 5 {
		t.Errorf("Expect size after delete = 5, but got %d", tree.Size)
	}

	// // node rearrange
	four := tree.Search(4)
	if four.Right.Value != 6 || four.Left.Value != 2 || four.Parent.Value != 1 {
		t.Errorf("tree.Search(4) %q", four)
	}
}
