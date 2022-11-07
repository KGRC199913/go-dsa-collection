package queue

import (
	"bytes"
	"fmt"
)

type SimpleQueue struct {
	data []any
}

func NewSimpleQueue() *SimpleQueue {
	return &SimpleQueue{}
}

func (queue *SimpleQueue) Enqueue(data any) {
	queue.data = append(queue.data, data)
}

func (queue *SimpleQueue) Dequeue() any {
	if queue.IsEmpty() {
		panic("Queue is empty")
	}
	data := queue.data[0]
	queue.data = queue.data[1:]
	return data
}

func (queue *SimpleQueue) IsEmpty() bool {
	return len(queue.data) == 0
}

func (queue *SimpleQueue) Size() int {
	return len(queue.data)
}

func (queue *SimpleQueue) Peek() any {
	if queue.IsEmpty() {
		panic("Queue is empty")
	}
	return queue.data[0]
}

func (queue *SimpleQueue) Clear() {
	queue.data = nil
}

func (queue *SimpleQueue) Contains(data any) bool {
	for _, d := range queue.data {
		if d == data {
			return true
		}
	}
	return false
}

func (queue *SimpleQueue) IndexOf(data any) int {
	for i, d := range queue.data {
		if d == data {
			return i
		}
	}
	return -1
}

// Merge two queue
func (queue *SimpleQueue) Merge(other *SimpleQueue) {
	queue.data = append(queue.data, other.data...)
}

// Copy the queue
func (queue *SimpleQueue) Copy() *SimpleQueue {
	newQueue := NewSimpleQueue()
	newQueue.data = make([]any, queue.Size())
	copy(newQueue.data, queue.data)
	return newQueue
}

// Split the queue into two queues
func (queue *SimpleQueue) Split(index int) (*SimpleQueue, *SimpleQueue) {
	if index < 0 || index >= queue.Size() {
		panic("Index out of bounds")
	}
	return &SimpleQueue{data: queue.data[:index]}, &SimpleQueue{data: queue.data[index:]}
}

// Reverse the queue
func (queue *SimpleQueue) Reverse() {
	for i, j := 0, queue.Size()-1; i < j; i, j = i+1, j-1 {
		queue.data[i], queue.data[j] = queue.data[j], queue.data[i]
	}
}

// String returns a string representation of the queue
// The string representation is "[v1, v2, v3, ...]"
func (queue *SimpleQueue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i, data := range queue.data {
		if i != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(fmt.Sprint(data))
	}
	buffer.WriteString("]")
	return buffer.String()

}