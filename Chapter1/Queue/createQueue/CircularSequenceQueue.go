// 循环队列

package createQueue

import "fmt"

type CircularQueue interface {
	Enqueue(val interface{})
	Dequeue() interface{}
	IsEmpty() bool
	GetFront() interface{}
	IsFull() bool
	GetSize() int
	resize(size int)
}

type circularQueue struct {
	queue []interface{}
	count int
	font  int
	tail  int
}

func (q *circularQueue) Enqueue(val interface{}) {
	if q.IsFull() {
		fmt.Println("队列已满，无法添加元素")
	} else {
		if (q.tail+1)%len(q.queue) == q.font {
			q.resize(len(q.queue) * 2)
		}
		q.queue[q.tail] = val
		q.tail = (q.tail + 1) % len(q.queue)
		q.count++
	}
}

func (q *circularQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		fmt.Println("队列已空")
		return nil
	} else {
		dequeueVar := q.GetFront()
		q.queue[q.font] = nil
		q.font = (q.font + 1) % len(q.queue)
		q.count--
		if q.count == cap(q.queue)/4 && cap(q.queue)/2 != 0 {
			q.resize(cap(q.queue) / 2)
		}
		return dequeueVar
	}
}

func (q *circularQueue) GetFront() interface{} {
	if q.IsEmpty() {
		fmt.Println("队列已空")
		return nil
	} else {
		return q.queue[q.font]
	}
}

func (q *circularQueue) GetSize() int {
	return q.count
}

func (q *circularQueue) IsFull() bool {
	return q.font == q.tail+1
}

func (q *circularQueue) IsEmpty() bool {
	return q.font == q.tail
}

func (q *circularQueue) resize(newSize int) {
	newQueue := make([]interface{}, newSize)
	for i := 0; i < q.count; i++ {
		newQueue[i] = q.queue[(i+q.font)%len(q.queue)]
	}
	q.queue = newQueue
	q.font = 0
	q.tail = q.count
}

func CreateCircularQueue(size int) CircularQueue {
	return &circularQueue{
		queue: make([]interface{}, size+1),
		count: 0,
		font:  0,
		tail:  0,
	}
}
