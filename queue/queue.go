package queue

// An FIFO queue.
type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)

}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEMpty() bool {
	return len(*q) == 0
}

// test
// test
