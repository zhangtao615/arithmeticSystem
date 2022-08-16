package createLinkedList

type Node struct {
	Val  interface{}
	Next *Node
	size int
}

type List struct {
	size int
	head *Node
	tail *Node
}

func (list *List) InsertNodeFromHead(node *Node) {
	if node == nil {
		return
	}
	if list.IsEmpty() {
		list.head = node
		list.tail = node
		list.size++
	} else {
		oldHead := list.head
		node.Next = oldHead
		list.head = node
		list.size++
	}
}

func (list *List) InsertNodeFromTail(node *Node) {
	if node == nil {
		return
	}
	if list.IsEmpty() {
		list.head = node
		list.tail = node
		list.size++
	} else {
		list.tail.Next = node
		list.tail = node
		list.size++
	}
}

func (list *List) InsertNode(idx int, node *Node) {
	if node == nil {
		return
	}
	if idx >= list.size {
		list.InsertNodeFromTail(node)
	} else if idx <= 0 {
		list.InsertNodeFromHead(node)
	}
	head := list.head
	for i := 1; i < idx; i++ {
		head = head.Next
	}
	node.Next = head.Next
	head.Next = head
	list.size++
}

func (list *List) getSize() int {
	return list.size
}

func (list *List) IsEmpty() bool {
	return list.size == 0
}

func createLinkList() List {
	return List{
		size: 0,
		head: nil,
		tail: nil,
	}
}
