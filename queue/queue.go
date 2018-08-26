package queue

// An FIFO queue.
type Queue []interface{}

// Pushes the element into the queue.
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

// Pops element from head.
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// Returns if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
