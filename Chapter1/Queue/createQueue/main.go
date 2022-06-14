package createQueue

import "fmt"

type CreateQueue interface {
	Enqueue(val interface{})
	Dequeue() interface{}
	GetFront() interface{}
	GetSize() int
	IsEmpty() bool
}

type queue struct {
	queue []interface{}
	size  int
}

func (q *queue) Enqueue(val interface{}) {
	q.queue = append(q.queue, val)
	q.size++
}

func (q *queue) Dequeue() interface{} {
	if q.IsEmpty() {
		fmt.Println("队列为空")
		return nil
	} else {
		dequeueVar := q.GetFront()
		q.queue = q.queue[1:q.GetSize()]
		q.size--
		return dequeueVar
	}
}

func (q *queue) GetFront() interface{} {
	if q.IsEmpty() {
		return nil
	} else {
		return q.queue[0]
	}
}

func (q *queue) GetSize() int {
	return q.size
}

func (q *queue) IsEmpty() bool {
	return q.size == 0
}

func NewQueue() CreateQueue {
	return &queue{
		queue: []interface{}{},
		size:  0,
	}

}
