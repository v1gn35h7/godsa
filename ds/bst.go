package ds

import (
	"fmt"
	//"os"

	List "container/list"
	"strconv"
	"strings"
)

type BstNode struct {
	data        int
	left, right *BstNode
}

var bstHead *List.Element

var tempArray []BstNode

func (bst *BstNode) PrintAllLeafs(bstHead *BstNode, array []int, len int) {
	if bstHead == nil {
		return
	}

	array = append(array, bstHead.data)
	len++
	if bstHead.left == nil && bstHead.right == nil {
		bst.printArry(array, len)
	}

	if bstHead.left != nil {
		bst.PrintAllLeafs(bstHead.left, array, len)
	}

	if bstHead.right != nil {
		bst.PrintAllLeafs(bstHead.right, array, len)
	}

	return
}

func (bst *BstNode) ReversepathUtil(bstHead *BstNode, array []*BstNode, len, k int) {
	if bstHead == nil {
		return
	}

	if bstHead.left == nil && bstHead.right == nil {
		return
	}

	array = append(array, bstHead)
	len++
	if k == bstHead.data {
		bst.ReversPath(bstHead, array)
	}

	if bstHead.left != nil {
		bst.ReversepathUtil(bstHead.left, array, len, k)
	}

	if bstHead.right != nil {
		bst.ReversepathUtil(bstHead.right, array, len, k)
	}

	return
}

func (bst *BstNode) ReversPath(BstNode *BstNode, path []*BstNode) {
	reversedPath := make([]int, 0)

	for i := (len(path) - 1); i >= 0; i-- {
		reversedPath = append(reversedPath, path[i].data)
	}

	for i := 0; i < len(path); i++ {
		fmt.Println(reversedPath[i], path[i].data)
		path[i].data = reversedPath[i]
	}
}

func (bst *BstNode) printArry(arr []int, len int) {
	for i := 0; i < len; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Printf("\n")
}

func (bst *BstNode) ListtoBST(n int) *BstNode {
	if n <= 0 {
		return nil
	}

	left := bst.ListtoBST(n / 2)

	TbstHead := &BstNode{bstHead.Value.(int), nil, nil}
	TbstHead.left = left

	bstHead = bstHead.Next()

	right := bst.ListtoBST(n - (n / 2) - 1)
	TbstHead.right = right

	return TbstHead
}

func (bst *BstNode) formTree(parent []string, n int) BstNode {
	var root BstNode
	ref := make([]*BstNode, n)

	for i := 0; i < n; i++ {
		ref[i] = &BstNode{i, nil, nil}
	}

	for i := 0; i < n; i++ {
		r, _ := strconv.Atoi(strings.TrimSpace(parent[i]))
		if r == -1 {
			root = *ref[i]
		} else {
			if ref[r].left == nil {
				ref[r].left = ref[i]
			} else {
				ref[r].right = ref[i]
			}
		}
	}

	return root
}

func (bst *BstNode) inOrder(root *BstNode) {
	if root == nil {
		return
	}

	bst.inOrder(root.left)
	fmt.Printf("%d\t", root.data)
	bst.inOrder(root.right)

	return
}

type BstQueue struct {
	data []*BstNode
}

func (queue *BstQueue) push(x *BstNode) []*BstNode {
	queue.data = append(queue.data, x)
	return queue.data
}

func (queue *BstQueue) pop() []*BstNode {
	queue.data = queue.data[1:]
	return queue.data
}

func (queue *BstQueue) peek() *BstNode {
	if len(queue.data) == 0 {
		return nil
	}
	return queue.data[0]
}

func (bst *BstNode) levelOrder(root *BstNode) {

	lqueue := BstQueue{make([]*BstNode, 0)}

	lqueue.push(root)
	temp := lqueue.peek()

	for temp != nil {
		if temp.left != nil {
			lqueue.push(temp.left)
		}
		if temp.right != nil {
			lqueue.push(temp.right)
		}
		fmt.Printf("%d ", temp.data)
		lqueue.pop()
		// fmt.Println(lqueue.data)
		temp = lqueue.peek()
	}
}

/*
* Binary Tree traversal using Morris algorithm
 */
func (bst *BstNode) MorrisTraversal(root *BstNode) {

	crnt := root

	for crnt != nil {

		if crnt.left == nil {
			fmt.Printf("%d ", crnt.data)
			crnt = crnt.right
			continue
		}

		pre := crnt.left
		for pre.right != nil {
			if pre.right.data == crnt.data {
				break
			}
			pre = pre.right
		}

		if pre.right == crnt {
			pre.right = nil
			fmt.Printf("%d ", crnt.data)
			crnt = crnt.right
		} else {
			pre.right = crnt
			crnt = crnt.left
		}
	}

	return
}
