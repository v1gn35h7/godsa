package ds

import (
	"fmt"
	//"os"

	List "container/list"
	"strconv"
	"strings"
)

type Node struct {
	data        int
	left, right *Node
}

func NewNode(val int) *Node {
	return &Node{
		val,
		nil,
		nil,
	}
}

func (n *Node) Data() int {
	return n.data
}

var head *List.Element

func GetNewBt() {
	//ioreader :=  bufio.NewReader(os.Stdin)
	//inp := scanLine(ioreader)
	//difs := strings.Split(inp, ",")
	//len(difs)
	//list := List.New()
	//root := formTree(difs, len(difs))
	root := &Node{0, nil, nil}
	root.right = &Node{2, nil, nil}
	root.left = &Node{1, nil, nil}
	root.left.left = &Node{3, nil, nil}
	root.left.right = &Node{4, nil, nil}
	root.left.left.left = &Node{5, nil, nil}

	morrisTraversal(root)

	//     for _, v := range difs {
	//         x, _ := strconv.Atoi(strings.TrimSpace(v))
	//         list.PushBack(x)
	//     }

	// //     for e := list.Front(); e != nil; e = e.Next() {
	// // 		fmt.Println(e.Value)
	// // 	}

	//	inOrder(root)
	// 	n := list.Len()
	// 	head = list.Front()
	// 	thead := ListtoBST(n)

	// 	fmt.Println(thead.right)

	//levelOrder(thead)

}

func BtListtoBST(n int) *Node {
	if n <= 0 {
		return nil
	}

	left := BtListtoBST(n / 2)

	Thead := &Node{head.Value.(int), nil, nil}
	Thead.left = left

	head = head.Next()

	right := BtListtoBST(n - (n / 2) - 1)
	Thead.right = right

	return Thead
}

func BtformTree(parent []string, n int) Node {
	var root Node
	ref := make([]*Node, n)

	for i := 0; i < n; i++ {
		ref[i] = &Node{i, nil, nil}
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

func inOrder(root *Node) {
	if root == nil {
		return
	}

	inOrder(root.left)
	fmt.Printf("%d\t", root.data)
	inOrder(root.right)

	return
}

func levelOrder(root *Node) {

	lqueue := Queue{make([]*Node, 0)}

	lqueue.Push(root)
	temp := lqueue.Peek()

	for temp != nil {
		if temp.left != nil {
			lqueue.Push(temp.left)
		}
		if temp.right != nil {
			lqueue.Push(temp.right)
		}
		fmt.Println(temp.data)
		lqueue.Pop()
		// fmt.Println(lqueue.data)
		temp = lqueue.Peek()
	}
}

func morrisTraversal(root *Node) {

	crnt := root

	for crnt != nil {

		if crnt.left == nil {
			fmt.Println(crnt.data)
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
			fmt.Println(crnt.data)
			crnt = crnt.right
		} else {
			pre.right = crnt
			crnt = crnt.left
		}
	}

	return
}
