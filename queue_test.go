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


func TestDequeueEqualPriority(t *testing.T) {

	var queue *Queue

	queue = InitQueue()

	
	queue.Insert(5)
	queue.Insert(6)
	queue.Insert(7)
	queue.Insert(8)
	


	v, _ := queue.Dequeue()

	if v != 5 {
		t.Errorf("Dequeue failed, expected 5 got %v", v)
	}

	fmt.Println("Successfully dequeued 5")

	v, _ = queue.Dequeue()

	if v != 6 {
		t.Errorf("Dequeue failed, expected 6 got %v", v)
	}

	fmt.Println("Successfully dequeued 6")

	v, _ = queue.Dequeue()

	if v != 7 {
		t.Errorf("Dequeue failed, expected 7 got %v", v)
	}

	fmt.Println("Successfully dequeued 7")


	v, _ = queue.Dequeue()

	if v != 8 {
		t.Errorf("Dequeue failed, expected 8 got %v", v)
	}

	fmt.Println("Successfully dequeued 8")


	_, err := queue.Dequeue()

	if err == nil {
		t.Errorf("Dequeue failed, expected error got %v", err)
	}


}




func TestInsertNotEqualPriority(t *testing.T) {

	var queue *Queue

	queue = InitQueue()

	
	queue.Insert(5, 5)
	queue.Insert(6, 4)
	queue.Insert(7, 1)
	queue.Insert(8, 2)
	


	queue.PrintQueue()

}



func TestDequeueNotEqualPriority(t *testing.T) {

	var queue *Queue

	queue = InitQueue()

	
	queue.Insert(5, 5)
	queue.Insert(6, 4)
	queue.Insert(7, 1)
	queue.Insert(8, 2)
	


	v, _ := queue.Dequeue()

	if v != 7 {
		t.Errorf("Dequeue failed, expected 7 got %v", v)
	}

	fmt.Println("Successfully dequeued 7")

	v, _ = queue.Dequeue()

	if v != 8 {
		t.Errorf("Dequeue failed, expected 8 got %v", v)
	}

	fmt.Println("Successfully dequeued 8")

	v, _ = queue.Dequeue()

	if v != 6 {
		t.Errorf("Dequeue failed, expected 6 got %v", v)
	}

	fmt.Println("Successfully dequeued 6")


	v, _ = queue.Dequeue()

	if v != 5 {
		t.Errorf("Dequeue failed, expected 5 got %v", v)
	}

	fmt.Println("Successfully dequeued 5")


	_, err := queue.Dequeue()

	if err == nil {
		t.Errorf("Dequeue failed, expected error got %v", err)
	}


}

