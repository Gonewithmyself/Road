package tree

import (
	"container/list"
	"fmt"
)

func (t Tree) preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.Data, " ")
	t.preOrder(node.left)
	t.preOrder(node.right)
}

func (t Tree) preOrderIter() {
	if t.Root == nil {
		return
	}
	stack := list.New()
	stack.PushBack(t.Root)

	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)
		top := e.Value.(*Node)
		fmt.Print(top.Data, " ")

		if top.right != nil {
			stack.PushBack(top.right)
		}

		if top.left != nil {
			stack.PushBack(top.left)
		}
	}

	fmt.Println()
}

func (t Tree) inOrder(node *Node) {
	if node == nil {
		return
	}

	t.inOrder(node.left)
	fmt.Print(node.Data, " ")
	t.inOrder(node.right)
}

func (t Tree) inOrderIter() {
	if t.Root == nil {
		return
	}
	stack := list.New()
	curr := t.Root

	for curr != nil || stack.Len() != 0 {
		if curr != nil {
			stack.PushBack(curr)
			curr = curr.left
		} else {
			e := stack.Back()
			stack.Remove(e)
			curr = e.Value.(*Node)
			fmt.Print(curr.Data, " ")
			curr = curr.right
		}
	}

	fmt.Println()
}

func (t Tree) postOrder(node *Node) {
	if node == nil {
		return
	}

	t.postOrder(node.left)
	t.postOrder(node.right)
	fmt.Print(node.Data, " ")
}

func (t Tree) postOrderIter() {
	if t.Root == nil {
		return
	}
	stack := list.New()
	stack2 := list.New()
	curr := t.Root
	stack.PushBack(curr)

	for stack.Len() != 0 {
		e := stack.Back()
		stack.Remove(e)
		curr = e.Value.(*Node)
		stack2.PushBack(curr)
		if curr.left != nil {
			stack.PushBack(curr.left)
		}

		if curr.right != nil {
			stack.PushBack(curr.right)
		}
	}

	for stack2.Len() != 0 {
		e := stack2.Back()
		stack2.Remove(e)
		fmt.Print(e.Value.(*Node).Data, " ")
	}

	fmt.Println()
}
