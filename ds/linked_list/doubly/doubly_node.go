package doubly

type LinkedListNode struct {
	data any
	prev *LinkedListNode
	next *LinkedListNode
}

func NewDoublyLinkedListNode(data any) *LinkedListNode {
	return &LinkedListNode{data: data}
}

func (node *LinkedListNode) GetData() any {
	return node.data
}

func (node *LinkedListNode) SetData(data any) {
	node.data = data
}

func (node *LinkedListNode) GetPrev() *LinkedListNode {
	return node.prev
}

func (node *LinkedListNode) SetPrev(prev *LinkedListNode) {
	node.prev = prev
}

func (node *LinkedListNode) GetNext() *LinkedListNode {
	return node.next
}

func (node *LinkedListNode) SetNext(next *LinkedListNode) {
	node.next = next
}
