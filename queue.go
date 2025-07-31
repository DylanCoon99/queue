package queue

import (
	"errors"
	"fmt"
	"strconv"
)

type Node struct {
	Value    int
	Next     *Node
	Priority int
}

type Queue struct {
	Head *Node // front of the queue; first to be dequeued
	Tail *Node // back of the queue; last to be dequeued
}

func InitQueue() *Queue {
	return &Queue{
		Head: nil,
		Tail: nil,
	}
}

func (q *Queue) PrintQueue() {

	output_string := ""

	cur := q.Tail

	for cur != nil {
		if cur.Next != nil {
			output_string += (strconv.Itoa(cur.Value) + " -> ")
		} else {
			output_string += (strconv.Itoa(cur.Value))
		}
		cur = cur.Next
	}

	fmt.Println(output_string)
}

// Note: 1 is highest priority, higher numbers are lower

func (q *Queue) Insert(params ...int) error {
	// params is just [val] for normal queue insert (no priority) and [val, priority] for specifc priority

	if len(params) == 2 {
		// there is a specified priority
		// iterate over the queue until you find a higher priority node
		priority := params[1]
		cur := q.Tail

		node := &Node{
			Value:    params[0],
			Next:     nil,
			Priority: priority,
		}

		if q.Head == nil {
			q.Tail = node
			q.Head = node
			return nil
		}


		if q.Tail.Priority <= priority {
			node.Next = q.Tail
			q.Tail = node
			return nil
		}

		for {
			if cur == q.Head {
				cur.Next = node
				q.Head = node
				break
			}
			if priority < cur.Next.Priority {
				cur = cur.Next
			} else {
				temp := cur.Next
				cur.Next = node
				node.Next = temp
				break
			}
		}

		return nil
	} else if len(params) == 1 {
		// else; just insert to the tail of the queue
		node := &Node{
			Value:    params[0],
			Next:     nil,
			Priority: 1,
		}

		if q.Head == nil {
			q.Tail = node
			q.Head = node
			return nil
		}

		node.Next = q.Tail

		q.Tail = node
		return nil
	} else {
		return errors.New("Invalid Params in Insert function")
	}

}
