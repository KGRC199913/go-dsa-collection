package singly

type LinkedListNode struct {
	data any
	next *LinkedListNode
}

func NewSinglyLinkedListNode(data any) *LinkedListNode {
	return &LinkedListNode{data: data}
}

func (node *LinkedListNode) GetData() any {
	return node.data
}

func (node *LinkedListNode) SetData(data any) {
	node.data = data
}

func (node *LinkedListNode) GetNext() *LinkedListNode {
	return node.next
}

func (node *LinkedListNode) SetNext(next *LinkedListNode) {
	node.next = next
}
