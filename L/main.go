package main

type Node struct {
	Value int
	Next  *Node
}

type List struct {
	head *Node
}

func (l *List) Add(value int) {
	newElem := &Node{
		Value: value,
		Next:  nil,
	}

	current := l.head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newElem
}

func (l *List) Reverse() {
	current := l.head
	var prev *Node = nil

	for current != nil {
		next := current.Next
		current.Next = prev

		prev = current
		current = next
	}

	l.head = prev
}

func main() {
	list := List{head: &Node{
		Value: 0,
		Next:  nil,
	}}

	list.Add(1)
	list.Add(2)

	list.Reverse()
}
