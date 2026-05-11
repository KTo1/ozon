package main

import "fmt"

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	key  int
	prev *Node
	next *Node
}

func NewList() *List {
	return &List{
		head: nil,
		tail: nil,
	}
}

func (l *List) PushFront(node *Node) *Node {
	if l.head == nil {
		l.head = node
		l.tail = node

		return l.head
	}

	node.next = l.head
	l.head.prev = node
	l.head = node

	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) MoveToFront(node *Node) *Node {
	if node == l.head {
		return l.head
	}

	if node == l.tail {
		node.prev.next = node.next
		//node.next.prev = node.prev
		l.tail = node.prev
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	node.next = l.head
	l.head.prev = node
	l.head = node

	return l.head
}

func (l *List) Delete(node *Node) {
	if node == l.head {
		prev := l.head.prev
		l.head.next = nil
		prev.prev = nil
		l.head = prev

		return
	}

	if node == l.tail {
		next := l.tail.next
		next.next = nil
		l.tail.prev = nil
		l.tail = next

		return
	}

	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev

	node.next = nil
	node.prev = nil

	return
}

func (l *List) Print() {
	cur := l.head
	for cur.next != nil {
		fmt.Printf(" -> %v", cur.key)
		cur = cur.next
	}

	if cur != nil {
		fmt.Printf(" -> %v", cur.key)
	}

	fmt.Printf("\n")
}

func main() {
	list := NewList()

	node1 := list.PushFront(&Node{
		key: 1,
	})

	list.PushFront(&Node{
		key: 2,
	})

	node3 := list.PushFront(&Node{
		key: 3,
	})

	list.PushFront(&Node{
		key: 4,
	})

	list.Print()

	list.MoveToFront(node1)
	list.Print()

	list.MoveToFront(node3)
	list.Print()

	list.Delete(node1)
	list.Print()
}
