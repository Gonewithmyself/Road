package tree

import (
	"fmt"
	"testing"
)

func Test_buildTree(t *testing.T) {
	tr := buildTree([]int{2, 3, 4, 9, 10, 11, 30, 50, 70, 88, 72}...)
	tr.print()

	tr.bfsOrder()

	fmt.Println(tr.preorderWithDivide())
	tr.preOrder(tr.Root)
	fmt.Println()
	tr.preOrderIter()

	tr.inOrder(tr.Root)
	fmt.Println()
	tr.inOrderIter()

	baseMp := make(map[string]interface{}, 0)
	transferSub, ok := baseMp["transfer_sub"].(string)
	_ = transferSub
	_ = ok

	tr.postOrder(tr.Root)
	fmt.Println()
	tr.postOrderIter()
	t.Error()
}
