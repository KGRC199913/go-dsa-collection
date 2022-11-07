package heap

import "golang.org/x/exp/constraints"

type MinHeap[T constraints.Ordered] struct {
	data []T
	size int
}

func NewMinHeap[T constraints.Ordered]() *MinHeap[T] {
	return &MinHeap[T]{
		data: make([]T, 0),
		size: 0,
	}
}

func (heap *MinHeap[T]) GetSize() int {
	return heap.size
}

func (heap *MinHeap[T]) IsEmpty() bool {
	return heap.size == 0
}

func (heap *MinHeap[T]) GetMin() any {
	if heap.IsEmpty() {
		panic("Heap is empty")
	}
	return heap.data[0]
}

func (heap *MinHeap[T]) ExtractMin() any {
	if heap.IsEmpty() {
		panic("Heap is empty")
	}
	min := heap.data[0]
	heap.data[0] = heap.data[heap.size-1]
	heap.size--
	heap.heapifyDown()
	return min
}

func (heap *MinHeap[T]) Insert(data any) {
	heap.data = append(heap.data, data)
	heap.size++
	heap.heapifyUp()
}

func (heap *MinHeap[T]) GetMax() any {
	if heap.IsEmpty() {
		panic("Heap is empty")
	}

	max := heap.data[heap.size/2]
	for i := heap.size/2 + 1; i < heap.size; i++ {
		if max < heap.data[i] {
			max = heap.data[i]
		}
	}

	return max
}

func (heap *MinHeap[T]) heapifyUp() {
	index := heap.size - 1
	for index > 0 && heap.data[index] < heap.data[heap.parent(index)] {
		heap.swap(index, heap.parent(index))
		index = heap.parent(index)
	}
}

func (heap *MinHeap[T]) parent(index int) int {
	if index == 0 {
		panic("Index out of bounds")
	}
	return (index - 1) / 2
}

// Array implementation of heapifyDown
func (heap *MinHeap[T]) heapifyDown() {
	index := 0
	for index < heap.size && heap.hasLeftChild(index) {
		smallerChildIndex := heap.getLeftChildIndex(index)
		if heap.hasRightChild(index) && heap.getRightChild(index) < heap.getLeftChild(index) {
			smallerChildIndex = heap.getRightChildIndex(index)
		}
		if heap.data[index] < heap.data[smallerChildIndex] {
			break
		} else {
			heap.swap(index, smallerChildIndex)
		}
		index = smallerChildIndex
	}

}

func (heap *MinHeap[T]) hasLeftChild(index int) bool {
	return heap.getLeftChildIndex(index) < heap.size
}

func (heap *MinHeap[T]) getLeftChildIndex(index int) int {
	return 2*index + 1
}

func (heap *MinHeap[T]) hasRightChild(index int) bool {
	return heap.getRightChildIndex(index) < heap.size
}

func (heap *MinHeap[T]) getRightChildIndex(index int) int {
	return 2*index + 2
}

func (heap *MinHeap[T]) getRightChild(index int) T {
	return heap.data[heap.getRightChildIndex(index)]
}

func (heap *MinHeap[T]) getLeftChild(index int) T {
	return heap.data[heap.getLeftChildIndex(index)]
}

func (heap *MinHeap[T]) swap(index int, index2 int) {
	heap.data[index], heap.data[index2] = heap.data[index2], heap.data[index]
}
