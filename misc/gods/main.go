package main

import (
	"fmt"

	"github.com/Gonewithmyself/gomod"
	"github.com/coreos/etcd/pkg/adt"
	"github.com/emirpasic/gods/maps/treemap"
	"github.com/emirpasic/gods/trees/btree"
)

func main() {
	ivtree()
}

func bTree() {
	tree := btree.NewWithIntComparator(3) // empty (keys are of type int)

	tree.Put(1, "x") // 1->x
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(1, "a") // 1->a, 2->b (in order, replacement)
	tree.Put(3, "c") // 1->a, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->a, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	tree.Put(6, "f") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)
	tree.Put(7, "g") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f, 7->g (in order)
	gomod.Greeting()

	fmt.Println(tree)
}

func rbtree() {
	m := treemap.NewWithIntComparator()
	for i := 0; i < 50; i++ {
		m.Put(i, i)
	}

	it := m.Iterator()
	c := 0
	for it.Next() {
		k := it.Key().(int)
		fmt.Println(k)
	}
	_ = c
}

func ivtree() {
	ivt := adt.NewIntervalTree()
	ivt.Insert(adt.NewInt64Interval(1, 3), 123)
	ivt.Insert(adt.NewInt64Interval(9, 13), 456)
	ivt.Insert(adt.NewInt64Interval(7, 20), 789)

	rs := ivt.Stab(adt.NewInt64Point(4))
	for _, v := range rs {
		fmt.Printf("Overlapping range: %+v\n", v)
	}

	v := ivt.Find(adt.NewInt64Interval(15, 19))
	fmt.Println(v)
}
