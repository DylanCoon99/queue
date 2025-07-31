package queue


type Node struct {
	Value    int
	Next     *Node
	Priority int
}



type Queue interface {
	Insert(params ...int) error  // params is just [val] for normal queue insert (no priority) and [val, priority] for specifc priority
	Dequeue() int
}



type Queue struct {
	Head *Node    // front of the queue; first to be dequeued
	Tail *Node    // back of the queue; last to be dequeued
}





func InitNormalQueue() *Queue {
	return &Queue{
		Head: nil,
		Tail: nil,
	}
}





func (q *Queue) Insert(params ...int) error {
	// params is just [val] for normal queue insert (no priority) and [val, priority] for specifc priority

	if len(params) == 2 {
		// there is a specified priority
	}


}

