package main

import "fmt"

func main() {
	doubleList := CreateDoubleList()
	doubleList.AddNodeFromHead(1)
	doubleList.AddNodeFromHead(2)
	doubleList.AddNodeBehindTail(3)
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

func (list *List) AddNodeFromHead(value int) {
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
}

// 从链表尾部添加

func (list *List) AddNodeBehindTail(value int) {
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
}

// 删除头节点

func (list *List) DeleteHead() {
	if list.Size < 1 {
		fmt.Errorf("链表为空")
		return
	} else {
		newHead := list.Head.Next
		newHead.Prev = nil
		list.Head = newHead
	}
	list.Size -= 1
}

// 删除尾部节点

func (list *List) DeleteTail() {
	if list.Size < 1 {
		fmt.Errorf("链表为空")
		return
	} else {
		newTail := list.Tail.Prev
		newTail.Next = nil
		list.Tail = newTail
	}
	list.Size -= 1
}

func CreateDoubleList() (list *List) {
	return &List{}
}
