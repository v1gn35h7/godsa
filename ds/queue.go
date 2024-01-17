package ds

import "fmt"

type Queue struct {
	data []*Node
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]*Node, 0),
	}
}

func (queue *Queue) Push(x *Node) []*Node {
	queue.data = append(queue.data, x)
	return queue.data
}

func (queue *Queue) Pop() []*Node {
	queue.data = queue.data[1:]
	return queue.data
}

func (queue *Queue) Peek() *Node {
	if len(queue.data) == 0 {
		return nil
	}
	return queue.data[0]
}

func (q *Queue) Load(index int) *Node {
	if index >= q.Len() {
		return nil
	}

	return q.data[index]
}

func (q *Queue) Len() int {
	return len(q.data)
}

func (q *Queue) Flush() {
	for _, v := range q.data {
		fmt.Println(v)
	}
}
