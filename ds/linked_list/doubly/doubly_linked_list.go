package doubly

import "fmt"

type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
	size int
}

func NewDoublyLinkedList() *LinkedList {
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
	node := NewDoublyLinkedListNode(data)
	node.SetNext(list.head)

	if list.head != nil {
		list.head.SetPrev(node)
	}
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
	node := NewDoublyLinkedListNode(data)
	node.SetPrev(list.tail)
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
	node := NewDoublyLinkedListNode(data)
	current := list.head
	for i := 0; i < index-1; i++ {
		current = current.GetNext()
	}
	next := current.GetNext()
	current.SetNext(node)
	node.SetPrev(current)
	node.SetNext(next)
	next.SetPrev(node)
	list.size++
}

func (list *LinkedList) RemoveHead() any {
	if list.IsEmpty() {
		panic("List is empty")
	}
	data := list.head.GetData()
	list.head = list.head.GetNext()
	if list.head != nil {
		list.head.SetPrev(nil)
	}
	list.size--
	return data
}

func (list *LinkedList) RemoveTail() any {
	if list.IsEmpty() {
		panic("List is empty")
	}
	data := list.tail.GetData()
	list.tail = list.tail.GetPrev()
	if list.tail != nil {
		list.tail.SetNext(nil)
	}
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
	current := list.head
	for i := 0; i < index-1; i++ {
		current = current.GetNext()
	}
	data := current.GetNext().GetData()
	current.SetNext(current.GetNext().GetNext())
	current.GetNext().SetPrev(current)
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
	current := list.tail
	for i := list.size - 1; i >= 0; i-- {
		if current.GetData() == data {
			return i
		}
		current = current.GetPrev()
	}
	return -1
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
	arr := make([]any, list.size)
	current := list.head
	for i := 0; i < list.size; i++ {
		arr[i] = current.GetData()
		current = current.GetNext()
	}
	return arr
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
	current := list.head
	for current != nil {
		next := current.GetNext()
		current.SetNext(current.GetPrev())
		current.SetPrev(next)
		current = next
	}
	temp := list.head
	list.head = list.tail
	list.tail = temp
}

// Merge two lists
func (list *LinkedList) Merge(other *LinkedList) *LinkedList {
	if list.IsEmpty() {
		return other
	}
	if other.IsEmpty() {
		return list
	}
	list.tail.SetNext(other.head)
	other.head.SetPrev(list.tail)
	list.tail = other.tail
	list.size += other.size
	return list
}

// Split a list into two lists
func (list *LinkedList) Split(index int) (*LinkedList, *LinkedList) {
	if index < 0 || index >= list.size {
		panic("Index out of bounds")
	}
	if index == 0 {
		return NewDoublyLinkedList(), list
	}
	if index == list.size-1 {
		return list, NewDoublyLinkedList()
	}
	left := NewDoublyLinkedList()
	right := NewDoublyLinkedList()
	current := list.head
	for i := 0; i < index; i++ {
		left.Append(current.GetData())
		current = current.GetNext()
	}
	for i := index; i < list.size; i++ {
		right.Append(current.GetData())
		current = current.GetNext()
	}
	return left, right
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

// Filter returns a new list containing only the elements that
// satisfy the given predicate
func (list *LinkedList) Filter(f func(any) bool) *LinkedList {
	filtered := NewDoublyLinkedList()
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
