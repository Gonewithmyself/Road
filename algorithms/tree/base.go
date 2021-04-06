package tree

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

type (
	Node struct {
		Data        int
		left, right *Node
	}

	Tree struct {
		Root *Node
	}
)

func (t *Tree) print() {
	if t.Root == nil {
		return
	}
	type level struct {
		*Node
		lv int
	}

	newlvNode := func(n *Node, lv int) *level {
		return &level{
			Node: n,
			lv:   lv,
		}
	}
	q := list.New()
	curlv := 1
	q.PushBack(newlvNode(t.Root, 1))

	for q.Len() != 0 {
		e := q.Front()
		node := e.Value.(*level)
		lv := node.lv
		if lv != curlv {
			curlv = lv
			fmt.Printf("\n")
		}

		if node.left != nil {
			q.PushBack(newlvNode(node.left, lv+1))
			fmt.Printf("/")
		} else {
			fmt.Printf(" ")
		}
		fmt.Printf("%v", node.Data)

		if node.right != nil {
			q.PushBack(newlvNode(node.right, lv+1))
			fmt.Printf("\\ ")
		} else {
			fmt.Printf("  ")
		}
		q.Remove(e)
	}

	fmt.Println()
}

func (t *Tree) push(data int) {
	if t.Root == nil {
		t.Root = &Node{Data: data}
		return
	}
	var (
		curr = t.Root
		prev = t.Root
	)
	for {
		n := rand.Intn(100)
		if n%2 == 1 {
			curr = prev.left
			if curr == nil {
				prev.left = &Node{Data: data}
				break
			}
			prev = prev.left
		} else {
			curr = prev.right
			if curr == nil {
				prev.right = &Node{Data: data}
				break
			}
			prev = prev.right
		}
	}
}

func buildTree(datas ...int) *Tree {
	t := &Tree{}
	for _, num := range datas {
		t.push(num)
	}
	return t
}

func init() {
	rand.Seed(time.Now().Unix())
}
