package main

import (
	"algorithmSystem/Chapter1/Queue/createQueue"
	"fmt"
)

func main() {
	//queue := createQueue.NewQueue()
	//queue.Enqueue(100)
	//queue.Enqueue(101)
	//queue.Enqueue(102)
	//queue.Dequeue()
	//fmt.Println(queue.GetFront())

	circularQueue := createQueue.CreateCircularQueue(3)
	circularQueue.Enqueue(10)
	circularQueue.Enqueue(11)
	circularQueue.Enqueue(12)
	circularQueue.Enqueue(13)
	circularQueue.Enqueue(14)
	circularQueue.Enqueue(15)
	circularQueue.Enqueue(16)
	circularQueue.Enqueue(17)
	circularQueue.Enqueue(18)
	circularQueue.Enqueue(19)
	circularQueue.Enqueue(20)
	circularQueue.Enqueue(21)
	circularQueue.Dequeue()
	circularQueue.Dequeue()
	circularQueue.Dequeue()
	circularQueue.Dequeue()
	circularQueue.Dequeue()
	circularQueue.Dequeue()
	circularQueue.Dequeue()
	circularQueue.Dequeue()
	fmt.Println(circularQueue)
}
