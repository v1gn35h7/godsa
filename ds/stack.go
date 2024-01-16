package ds

import "fmt"

type Snode struct {
	Data int
	Next *Snode
}

// Asked in appian
func GetNthNodeFromTail() {
	head := Snode{0, nil}
	tmpHead := &head
	var res Snode

	for i := 1; i < 10; i++ {
		tmpHead.Next = &Snode{i, nil}
		tmpHead = tmpHead.Next
	}

	recFunc(&res, &head, 3)

	fmt.Println(res)
}

func recFunc(res, head *Snode, n int) int {

	if head.Next == nil {
		return 1
	}

	v := 1 + recFunc(res, head.Next, n)

	if v == n {
		res = head
		return v
	}

	return v
}
