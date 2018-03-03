package list

import (
	"fmt"
	"log"
)

// node (Private) - Defines the structure for each individual node in a linked list
type node struct {
	data interface{} // Value of Node
	next *node       // Pointer to the next Node
	list *List       // Pointer to the list it is attached to
}

// nextNode (Private) - Returns the next node in the list
func (n node) nextNode() *node {
	// returns nil if there is not list AND if the pointer to the next
	// node is the same as the head's next node there for there is next node
	if next := n.next; n.list != nil && next != &n.list.head {
		return next
	}
	return nil
}

// List (Public) - The container for all the linked nodes in a set
type List struct {
	head node // the begining node
	tail node
	size int // size of the list
}

// init (Private) - Generates a linked list with Size=0 and head pointing to itself
func (l *List) init() *List {
	l.head.next = &l.head
	l.tail.next = l.head.next
	l.size = 0
	return l
}

// New (Public) - Returns an initialized list.
func New() *List { return new(List).init() }

// AddFirst (Public) - Adds an item to the begining of the list, pushes the
// list over 1 to make room.
func (l *List) AddFirst(T interface{}) {
	if T == nil {
		panic("Cannot add a nil value to deque")
	}
	new := &node{data: T, list: l}
	if l.size != 0 {
		prev := l.head.next
		new.next = prev
		l.tail.next = new
	}
	l.head.next = new
	l.size++
}

// AddLast (Public) - traverses to the end of the list and adds the item to the end.
func (l *List) AddLast(T interface{}) {
	if T == nil {
		log.Println("Cannot add a nil value to deque")
		return
	}
	new := node{data: T, list: l}
	(l.tail.next.next) = &new
	(l.tail.next) = &new
	/*
		for current := l.head.next; current != nil; current = current.nextNode() {
			if current.next == nil {
				current.next = new
				break
			}
		}*/
	l.size++
}

// RemoveFirst (Public) - removes and shifts the list left. Returns the value of the removed item
func (l *List) RemoveFirst() interface{} {
	if l.size == 0 {
		return nil
	}
	removed := l.head.next
	l.head.next = removed.next
	l.size--
	return removed.data
}

// RemoveLast (Public) - Traverse the list and removes the last item in the list
// So there was a memory leak in my initial program, I had to check for the case
// when I have 1 item in the list and also more than one.
func (l *List) RemoveLast() interface{} {
	if l.size == 0 {
		return nil
	}
	var removed interface{}
	if l.size == 1 {
		removed = l.head.next.data
		l.head.next = nil
		l.size--
		return removed
	}
	for current := l.head.next; current != nil; current = current.nextNode() {
		if current.next.next == nil {
			removed = current.next.data
			current.next = nil
			break
		}
	}
	l.size--
	return removed
}

// GetFirst (Public) - Returns the data from the first node in the list.
func (l *List) GetFirst() interface{} {
	return l.head.next.data
}

// GetLast (Public) - loops over the list to find and return the data from the last item.
func (l *List) GetLast() interface{} {
	for current := l.head.next; current != nil; current = current.nextNode() {
		if current.next == nil {
			return current.data
		}
	}
	return nil
}

// Contains (Public) - Returns true or false whether an item was contained in the list
func (l *List) Contains(T interface{}) bool {
	for current := l.head.next; current != nil; current = current.nextNode() {
		if current.data == T {
			return true
		}
	}
	return false
}

// Size (Public) - Returns the size of the list
func (l *List) Size() int {
	return l.size
}

// String (Public) - Allows for the fmt.Print* functions to print the list struct as a string.
func (l *List) String() string {
	if l.size == 0 {
		return "[ ]"
	}
	result := "[ "
	for current := l.head.next; current != nil; current = current.nextNode() {
		result += fmt.Sprintf("%v ", current.data)
	}
	return result + "]"
}
