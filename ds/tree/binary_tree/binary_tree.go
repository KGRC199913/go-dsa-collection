package binary_tree

type Tree struct {
	root *Node
}

func NewBinaryTree() *Tree {
	return &Tree{}
}

func NewBinaryTreeWithRoot(root *Node) *Tree {
	return &Tree{root: root}
}

func (tree *Tree) GetRoot() *Node {
	return tree.root
}

func (tree *Tree) SetRoot(root *Node) {
	tree.root = root
}

func (tree *Tree) IsEmpty() bool {
	return tree.root == nil
}

func (tree *Tree) Height() int {
	return tree.height(tree.root)
}

func (tree *Tree) height(root *Node) int {
	if root == nil {
		return 0
	}
	leftHeight := tree.height(root.left)
	rightHeight := tree.height(root.right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (tree *Tree) Size() int {
	return tree.size(tree.root)
}

func (tree *Tree) size(root *Node) int {
	if root == nil {
		return 0
	}
	return 1 + tree.size(root.left) + tree.size(root.right)
}

func (tree *Tree) Map(f func(any) any) {
	tree.mapTree(tree.root, f)
}

func (tree *Tree) mapTree(root *Node, f func(any) any) {
	if root == nil {
		return
	}
	root.data = f(root.data)
	tree.mapTree(root.left, f)
	tree.mapTree(root.right, f)
}

func (tree *Tree) PreOrder() []any {
	var result []any
	tree.preOrder(tree.root, &result)
	return result
}

func (tree *Tree) preOrder(root *Node, i *[]any) {
	if root == nil {
		return
	}
	*i = append(*i, root.data)
	tree.preOrder(root.left, i)
	tree.preOrder(root.right, i)
}

func (tree *Tree) InOrder() []any {
	var result []any
	tree.inOrder(tree.root, &result)
	return result
}

func (tree *Tree) inOrder(root *Node, i *[]any) {
	if root == nil {
		return
	}
	tree.inOrder(root.left, i)
	*i = append(*i, root.data)
	tree.inOrder(root.right, i)
}

func (tree *Tree) PostOrder() []any {
	var result []any
	tree.postOrder(tree.root, &result)
	return result
}

func (tree *Tree) postOrder(root *Node, i *[]any) {
	if root == nil {
		return
	}
	tree.postOrder(root.left, i)
	tree.postOrder(root.right, i)
	*i = append(*i, root.data)
}

// Insert a new node into the next available position in the tree. Try to keep the tree balanced.
func (tree *Tree) Insert(data any) {
	tree.insert(tree.root, data)
}

func (tree *Tree) insert(root *Node, data any) {
	if tree.root == nil {
		tree.root = NewBinaryTreeNode(data)
		return
	}
	if root.left == nil {
		root.left = NewBinaryTreeNode(data)
		return
	}
	if root.right == nil {
		root.right = NewBinaryTreeNode(data)
		return
	}
	if tree.height(root.left) < tree.height(root.right) {
		tree.insert(root.left, data)
	} else {
		tree.insert(root.right, data)
	}
}

// Delete a node from the tree. Try to keep the tree balanced.
func (tree *Tree) Delete(data any) {
	tree.delete(tree.root, data)
}

func (tree *Tree) delete(root *Node, data any) {
	if root == nil {
		return
	}
	if root.left != nil && root.left.data == data {
		root.left = nil
		return
	}
	if root.right != nil && root.right.data == data {
		root.right = nil
		return
	}
	if tree.height(root.left) < tree.height(root.right) {
		tree.delete(root.right, data)
	} else {
		tree.delete(root.left, data)
	}
}

// Search for a node in the tree.
func (tree *Tree) Search(data any) bool {
	return tree.search(tree.root, data)
}

func (tree *Tree) search(root *Node, data any) bool {
	if root == nil {
		return false
	}
	if root.data == data {
		return true
	}
	return tree.search(root.left, data) || tree.search(root.right, data)
}
