package DoubleList

import "fmt"

//func main() {
//	doubleList := CreateDoubleList()
//	doubleList.AddNodeFromHead(1)
//	doubleList.AddNodeFromHead(2)
//	doubleList.AddNodeBehindTail(3)
//	//doubleList.DeleteTail()
//	//doubleList.DeleteHead()
//	length := getLength(doubleList.Head)
//
//	dummy := &ListNode{nil, doubleList.Head, 0}
//	for i := 0; i < length; i++ {
//		fmt.Println(dummy.Next.Val)
//		dummy = dummy.Next
//	}
//}

func GetLength(list *ListNode) int {
	length := 0
	for ; list != nil; list = list.Next {
		length++
	}
	return length
}

type ListNode struct {
	Prev *ListNode
	Next *ListNode
	Val  int
}

type List struct {
	Head *ListNode
	Tail *ListNode
	Size int
}

// 从头节点添加

func (list *List) AddNodeFromHead(value int) int {
	newNode := new(ListNode)
	newNode.Val = value

	if list.Size < 1 {
		list.Head = newNode
		list.Tail = newNode
	} else {
		headNode := list.Head
		headNode.Prev = newNode
		newNode.Next = headNode
		list.Head = newNode
	}
	list.Size += 1
	return list.Size
}

// 从链表尾部添加

func (list *List) AddNodeBehindTail(value int) int {
	newNode := new(ListNode)
	newNode.Val = value

	// 链表为空时
	if list.Size < 1 {
		list.Head = newNode
		list.Tail = newNode
	} else {
		// 链表的长度>=1时
		tailNode := list.Tail
		tailNode.Next = newNode
		newNode.Prev = tailNode
		newNode.Next = nil
		list.Tail = newNode
	}

	list.Size += 1
	return list.Size
}

// 删除头节点

func (list *List) DeleteHead() *ListNode {
	currentHead := list.Head
	if list.Size < 1 {
		fmt.Errorf("链表为空")
	} else {
		newHead := list.Head.Next
		newHead.Prev = nil
		list.Head = newHead
		list.Size -= 1
	}
	return currentHead
}

// 删除尾部节点

func (list *List) DeleteTail() *ListNode {
	currentTail := list.Head
	if list.Size < 1 {
		fmt.Errorf("链表为空")
	} else {
		newTail := list.Tail.Prev
		newTail.Next = nil
		list.Tail = newTail
		list.Size -= 1
	}
	return currentTail
}

// 删除任意节点
func (list *List) Remove(node *ListNode) *ListNode {
	// 如果是node == nil
	if node == nil {
		fmt.Errorf("节点不存在")
	} else if node == list.Head {
		list.DeleteHead()
	} else if node == list.Tail {
		list.DeleteTail()
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		list.Size -= 1
	}
	return node
}

// 添加任意节点
func (list *List) Add(node *ListNode) int {
	if list.Size < 1 {
		list.AddNodeFromHead(node.Val)
	} else {
		node.Prev.Next = node
		node.Next.Prev = node
	}
	list.Size += 1
	return list.Size
}

func CreateDoubleList() (list *List) {
	return &List{}
}
