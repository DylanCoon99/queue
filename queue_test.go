package queue

import (
	//"log"
	"fmt"
	"testing"
)

/*
func InitExampleQueue() {

	return
}
*/

func TestInitQueue(t *testing.T) {

	var queue *Queue

	queue = InitQueue()

	if queue.Head != nil {
		t.Errorf("InitQueue failed, expected nil got %v", queue.Head)
	}

	if queue.Tail != nil {
		t.Errorf("InitQueue failed, expected nil got %v", queue.Tail)
	}

	fmt.Println("Successfully init queue.")

}

func TestInsertEqualPriority(t *testing.T) {


	var queue *Queue

	queue = InitQueue()

	queue.Insert(5)
	queue.Insert(6)
	queue.Insert(7)
	queue.Insert(8)


	queue.PrintQueue()

}
