package queue

import "fmt"

type CircularQueue struct {
	data []any
	head int
	tail int
	size int
}

func NewCircularQueue() *CircularQueue {
	return &CircularQueue{
		data: make([]any, 0),
		head: -1,
		tail: -1,
		size: 0,
	}
}

func (queue *CircularQueue) GetSize() int {
	return queue.size
}

func (queue *CircularQueue) IsEmpty() bool {
	return queue.size == 0
}

func (queue *CircularQueue) Enqueue(data any) {
	if queue.size == 0 {
		queue.head = 0
	}
	queue.tail = (queue.tail + 1) % len(queue.data)
	if queue.tail == queue.head {
		queue.data = append(queue.data, data)
	} else {
		queue.data[queue.tail] = data
	}
	queue.size++
}

func (queue *CircularQueue) Dequeue() any {
	if queue.IsEmpty() {
		panic("Queue is empty")
	}
	data := queue.data[queue.head]
	queue.head = (queue.head + 1) % len(queue.data)
	queue.size--
	return data
}

func (queue *CircularQueue) Size() int {
	return queue.size
}

func (queue *CircularQueue) Peek() any {
	if queue.IsEmpty() {
		panic("Queue is empty")
	}
	return queue.data[queue.head]
}

func (queue *CircularQueue) Clear() {
	queue.data = nil
	queue.head = -1
	queue.tail = -1
	queue.size = 0
}

func (queue *CircularQueue) Contains(data any) bool {
	for i := 0; i < queue.size; i++ {
		if queue.data[(queue.head+i)%len(queue.data)] == data {
			return true
		}
	}
	return false
}

func (queue *CircularQueue) IndexOf(data any) int {
	for i := 0; i < queue.size; i++ {
		if queue.data[(queue.head+i)%len(queue.data)] == data {
			return i
		}
	}
	return -1
}

// Merge two queue
func (queue *CircularQueue) Merge(other *CircularQueue) {
	for i := 0; i < other.size; i++ {
		queue.Enqueue(other.data[(other.head+i)%len(other.data)])
	}
}

// Copy queue
func (queue *CircularQueue) Copy() *CircularQueue {
	newQueue := NewCircularQueue()
	newQueue.data = make([]any, queue.Size())
	copy(newQueue.data, queue.data)
	newQueue.head = queue.head
	newQueue.tail = queue.tail
	newQueue.size = queue.size
	return newQueue
}

// Split queue
func (queue *CircularQueue) Split() (*CircularQueue, *CircularQueue) {
	newQueue := queue.Copy()
	newQueue2 := NewCircularQueue()
	for i := 0; i < queue.size/2; i++ {
		newQueue2.Enqueue(newQueue.Dequeue())
	}
	return newQueue, newQueue2
}

// Reverse queue
func (queue *CircularQueue) Reverse() {
	newQueue := NewCircularQueue()
	for i := 0; i < queue.size; i++ {
		newQueue.Enqueue(queue.Dequeue())
	}
	queue.Merge(newQueue)
}

// Rotate queue
func (queue *CircularQueue) Rotate(n int) {
	if n < 0 {
		for i := 0; i < -n; i++ {
			queue.Enqueue(queue.Dequeue())
		}
	} else {
		for i := 0; i < n; i++ {
			queue.Enqueue(queue.Dequeue())
		}
	}
}

// String returns a string representation of the queue
func (queue *CircularQueue) String() string {
	str := "Queue: head ["
	for i := 0; i < queue.size; i++ {
		str += fmt.Sprint(queue.data[(queue.head+i)%len(queue.data)])
		if i != queue.size-1 {
			str += ", "
		}
	}
	str += "] tail"
	return str
}
