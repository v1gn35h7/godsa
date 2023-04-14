package ds

type Queue struct {
	data []*Node
}

func (queue *Queue) push(x *Node) []*Node {
	queue.data = append(queue.data, x)
	return queue.data
}

func (queue *Queue) pop() []*Node {
	queue.data = queue.data[1:]
	return queue.data
}

func (queue *Queue) peek() *Node {
	if len(queue.data) == 0 {
		return nil
	}
	return queue.data[0]
}
