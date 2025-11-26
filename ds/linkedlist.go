package ds

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

// 1
// 2 7
// 1 2 5 7 10 ...

func merge(head1, head2 *ListNode) *ListNode {

	var resHead *ListNode
	var temphead *ListNode

	for head1.next != nil && head2.next != nil {

		if head1.data < head2.data {
			if resHead == nil {
				resHead = head1
				temphead = resHead
			} else {
				temphead.next = head1
			}
			head1 = head1.next
		} else {
			if resHead == nil {
				resHead = head2
				temphead = resHead
			} else {
				temphead.next = head2
			}
			head2 = head2.next
		}
		temphead = temphead.next
	}

	if head1 != nil {
		temphead.next = head1

	}

	if head2 != nil {
		temphead.next = head2
	}

	return resHead

}

// O(N + M)
func NewLinkedLIst() *ListNode {
	head := &ListNode{data: 1}

	temp := head
	for i := 0; i < 10; i++ {
		node := &ListNode{
			data: i + 2,
		}
		temp.next = node
		temp = temp.next
		fmt.Println(temp)
	}

	return head
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Print(head.data, "-> ")
		head = head.next
	}

}
