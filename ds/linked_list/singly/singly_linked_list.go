package singly

import "fmt"

type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
	size int
}

func NewSinglyLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) GetHead() *LinkedListNode {
	return list.head
}

func (list *LinkedList) GetTail() *LinkedListNode {
	return list.tail
}

func (list *LinkedList) GetSize() int {
	return list.size
}

func (list *LinkedList) IsEmpty() bool {
	return list.size == 0
}

func (list *LinkedList) Prepend(data any) {
	node := NewSinglyLinkedListNode(data)
	node.SetNext(list.head)
	list.head = node
	if list.tail == nil {
		list.tail = node
	}
	list.size++
}

func (list *LinkedList) Append(data any) {
	if list.IsEmpty() {
		list.Prepend(data)
		return
	}
	node := NewSinglyLinkedListNode(data)
	list.tail.SetNext(node)
	list.tail = node
	list.size++
}

func (list *LinkedList) Insert(data any, index int) {
	if index < 0 || index > list.size {
		panic("Index out of bounds")
	}
	if index == 0 {
		list.Prepend(data)
		return
	}
	if index == list.size {
		list.Append(data)
		return
	}
	previous := list.head
	for i := 0; i < index-1; i++ {
		previous = previous.GetNext()
	}
	node := NewSinglyLinkedListNode(data)
	node.SetNext(previous.GetNext())
	previous.SetNext(node)
	list.size++
}

func (list *LinkedList) RemoveHead() any {
	if list.IsEmpty() {
		panic("List is empty")
	}
	data := list.head.GetData()
	list.head = list.head.GetNext()
	list.size--
	if list.IsEmpty() {
		list.tail = nil
	}
	return data
}

func (list *LinkedList) RemoveTail() any {
	if list.IsEmpty() {
		panic("List is empty")
	}
	data := list.tail.GetData()
	if list.size == 1 {
		list.head = nil
		list.tail = nil
		list.size--
		return data
	}
	previous := list.head
	for previous.GetNext() != list.tail {
		previous = previous.GetNext()
	}
	previous.SetNext(nil)
	list.tail = previous
	list.size--
	return data
}

func (list *LinkedList) RemoveAt(index int) any {
	if index < 0 || index >= list.size {
		panic("Index out of bounds")
	}
	if index == 0 {
		return list.RemoveHead()
	}
	if index == list.size-1 {
		return list.RemoveTail()
	}
	previous := list.head
	for i := 0; i < index-1; i++ {
		previous = previous.GetNext()
	}
	data := previous.GetNext().GetData()
	previous.SetNext(previous.GetNext().GetNext())
	list.size--
	return data
}

func (list *LinkedList) FirstIndexOf(data any) int {
	current := list.head
	for i := 0; i < list.size; i++ {
		if current.GetData() == data {
			return i
		}
		current = current.GetNext()
	}
	return -1
}

func (list *LinkedList) LastIndexOf(data any) int {
	current := list.head
	index := -1
	for i := 0; i < list.size; i++ {
		if current.GetData() == data {
			index = i
		}
		current = current.GetNext()
	}
	return index
}

func (list *LinkedList) Contains(data any) bool {
	return list.FirstIndexOf(data) != -1
}

func (list *LinkedList) PeekHead() any {
	if list.IsEmpty() {
		panic("List is empty")
	}
	return list.head.GetData()
}

func (list *LinkedList) PeekTail() any {
	if list.IsEmpty() {
		panic("List is empty")
	}
	return list.tail.GetData()
}

func (list *LinkedList) PeekAt(index int) any {
	if index < 0 || index >= list.size {
		panic("Index out of bounds")
	}
	current := list.head
	for i := 0; i < index; i++ {
		current = current.GetNext()
	}
	return current.GetData()
}

func (list *LinkedList) Clear() {
	// Traverse the list and set each node's next to nil
	current := list.head
	for current != nil {
		next := current.GetNext()
		current.SetNext(nil)
		current = next
	}
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list *LinkedList) ToArray() []any {
	array := make([]any, list.size)
	current := list.head
	for i := 0; i < list.size; i++ {
		array[i] = current.GetData()
		current = current.GetNext()
	}
	return array
}

// [1 -> 2 -> 3 -> 4 -> 5]
func (list *LinkedList) String() string {
	if list.IsEmpty() {
		return "[]"
	}
	str := "["
	current := list.head
	for current != nil {
		str += fmt.Sprintf("%v", current.GetData())
		if current.GetNext() != nil {
			str += " -> "
		}
		current = current.GetNext()
	}
	str += "]"
	return str
}

func (list *LinkedList) Reverse() {
	if list.IsEmpty() {
		return
	}
	var previous *LinkedListNode
	current := list.head
	for current != nil {
		next := current.GetNext()
		current.SetNext(previous)
		previous = current
		current = next
	}
	list.tail = list.head
	list.head = previous
}

// Merge another list into this list
func (list *LinkedList) Merge(other *LinkedList) {
	if other.IsEmpty() {
		return
	}
	if list.IsEmpty() {
		list.head = other.head
		list.tail = other.tail
		list.size = other.size
		return
	}
	list.tail.SetNext(other.head)
	list.tail = other.tail
	list.size += other.size
}

// Split this list into two lists at the given index
func (list *LinkedList) Split(index int) (*LinkedList, *LinkedList) {
	if index < 0 || index > list.size {
		panic("Index out of bounds")
	}
	if index == 0 {
		return nil, list
	}
	if index == list.size {
		return list, nil
	}
	previous := list.head
	for i := 0; i < index-1; i++ {
		previous = previous.GetNext()
	}
	other := &LinkedList{previous.GetNext(), list.tail, list.size - index}
	previous.SetNext(nil)
	list.tail = previous
	list.size = index
	return list, other
}

// Map applies the given function to each element in the list
func (list *LinkedList) Map(f func(any) any) {
	current := list.head
	for current != nil {
		current.SetData(f(current.GetData()))
		current = current.GetNext()
	}
}

// Fold applies the given function to each element in the list
// and returns the result
func (list *LinkedList) Fold(f func(any, any) any, initial any) any {
	current := list.head
	for current != nil {
		initial = f(initial, current.GetData())
		current = current.GetNext()
	}
	return initial
}

// Filter removes all elements from the list that do not satisfy the given predicate using two pointers
func (list *LinkedList) Filter(f func(any) bool) *LinkedList {
	filtered := NewSinglyLinkedList()
	current := list.head
	for current != nil {
		if f(current.GetData()) {
			filtered.Append(current.GetData())
		}
		current = current.GetNext()
	}
	return filtered
}

// Any returns true if any element in the list satisfies the given predicate
func (list *LinkedList) Any(f func(any) bool) bool {
	current := list.head
	for current != nil {
		if f(current.GetData()) {
			return true
		}
		current = current.GetNext()
	}
	return false
}

// Every returns true if all elements in the list satisfy the given predicate
func (list *LinkedList) Every(f func(any) bool) bool {
	current := list.head
	for current != nil {
		if !f(current.GetData()) {
			return false
		}
		current = current.GetNext()
	}
	return true
}

// HasCycle detect circular linked list using turtle and hare algorithm
func (list *LinkedList) HasCycle() bool {
	if list.IsEmpty() {
		return false
	}
	turtle := list.head
	hare := list.head
	for hare != nil && hare.GetNext() != nil {
		turtle = turtle.GetNext()
		hare = hare.GetNext().GetNext()
		if turtle == hare {
			return true
		}
	}
	return false
}
