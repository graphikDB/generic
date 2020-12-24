package generic

// create a node that holds the graphs vertex as data
type queueNode struct {
	value interface{}
	next  *queueNode
}

// create a queue data structure
type simpleQueue struct {
	head   *queueNode
	tail   *queueNode
	length int
}

func NewQueue() Queue {
	return &simpleQueue{}
}

func (q *simpleQueue) Enqueue(value interface{}) {
	q.length++
	n := &queueNode{value: value}
	// if the queue is empty, set the head and tail as the node value
	if q.tail == nil {
		q.head = n
		q.tail = n
		return
	}
	q.tail.next = n
	q.tail = n
}

//
func (q *simpleQueue) Dequeue() interface{} {
	n := q.head
	// return nil, if head is empty
	if n == nil {
		return nil
	}

	q.head = q.head.next

	// if there wasn't any next node, that
	// means the queue is empty, and the tail
	// should be set to nil
	if q.head == nil {
		q.tail = nil
	}
	q.length--
	return n.value
}

func (q *simpleQueue) Len() int {
	return q.length
}
