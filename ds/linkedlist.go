package ds

type ListNode struct {
	data int
	next *Node
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
