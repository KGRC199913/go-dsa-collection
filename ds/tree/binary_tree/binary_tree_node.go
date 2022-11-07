package binary_tree

type Node struct {
	data  any
	left  *Node
	right *Node
}

func NewBinaryTreeNode(data any) *Node {
	return &Node{data: data}
}

func (node *Node) GetData() any {
	return node.data
}

func (node *Node) GetLeft() *Node {
	return node.left
}

func (node *Node) GetRight() *Node {
	return node.right
}

func (node *Node) SetData(data any) {
	node.data = data
}

func (node *Node) SetLeft(left *Node) {
	node.left = left
}

func (node *Node) SetRight(right *Node) {
	node.right = right
}

func (node *Node) IsLeaf() bool {
	return node.left == nil && node.right == nil
}

func (node *Node) IsFull() bool {
	return node.left != nil && node.right != nil
}
