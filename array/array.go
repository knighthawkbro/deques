package array

import (
	"fmt"
)

// Array (Public) - Structure that defines
type Array struct {
	size       int
	collection []interface{}
}

// Init (Public) - initializes the array with whatever size is provided, This is what can be overrided by the user.
func (a *Array) Init(capacity int) *Array {
	if capacity < 0 {
		return nil
	}
	a.collection = make([]interface{}, capacity)
	a.size = 0
	return a
}

// New (Public) - Returns an initialized array with default size of 10.
func New() *Array { return new(Array).Init(10) }

// AddFirst (Public) - ensures the array is big enough for the new item, then
// shifts all the items right one space, then adds the first item to the front,
// lastly it increments the size
func (a *Array) AddFirst(T interface{}) {
	a.ensureSpace()
	a.shiftRight(0)
	a.collection[0] = T
	a.size++
}

// AddLast (Public) - ensures the array is big enough for the new item, then
// adds the item to the end of the array list, lastly it increments the size
func (a *Array) AddLast(T interface{}) {
	a.ensureSpace()
	a.collection[a.size] = T
	a.size++
}

// RemoveFirst (Public) - checks if array is empty, then stores the result in a
// variable, then it shifts everything but the first item left, then it decrements
// the size and finally returns the removed item.
func (a *Array) RemoveFirst() interface{} {
	if a.size == 0 {
		return nil
	}
	result := a.collection[0]
	a.shiftLeft(0)
	a.size--
	return result
}

// RemoveLast (Public) - Checks for an empty array, then stores the last item in
// the array, then it nils the last value, then decrements the size, finaly
// returns the removed item.
func (a *Array) RemoveLast() interface{} {
	if a.size == 0 {
		return nil
	}
	result := a.collection[a.size-1]
	a.collection[a.size-1] = nil
	a.size--
	return result
}

// GetFirst (Public) - Returns nil if array is empty, returns first item
func (a *Array) GetFirst() interface{} {
	if a.size == 0 {
		return nil
	}
	return a.collection[0]
}

// GetLast (Public) - Returns nil if array is empty, returns last item
func (a *Array) GetLast() interface{} {
	if a.size == 0 {
		return nil
	}
	return a.collection[a.size-1]
}

// Contains (Public) - Checks to see if item is in array, returns true or false
func (a *Array) Contains(T interface{}) bool {
	for _, item := range a.collection {
		if item == T {
			return true
		}
	}
	return false
}

// Size (Public) - returns the size of the Array
func (a *Array) Size() int {
	return a.size
}

// String (Public) - formats the array when fmt.Print is called.
func (a *Array) String() string {
	if a.size == 0 {
		return "[ ]"
	}
	s := "[ "
	for x := 0; x < a.size; x++ {
		s += fmt.Sprintf("%v ", a.collection[x])
	}
	return s + "]"
}

// ensureSpace (Private) - Sees if the size and capacity of the array are the same. If so,
// It creates a new array with double the capacity and overwrites the old array with a new
// array, then clears the new array for the GC.
func (a *Array) ensureSpace() {
	if a.size == cap(a.collection) {
		new := new(Array).Init(cap(a.collection) * 2)
		new.size = a.size
		for i := 0; i < a.size; i++ {
			new.collection[i] = a.collection[i]
		}
		*a = *new
		new = nil
	}
}

// shiftLeft (Private) - Moves all the items left after index (Destructive)
func (a *Array) shiftLeft(index int) {
	for i := index; i < a.size-1; i++ {
		a.collection[i] = a.collection[i+1]
	}
}

// shiftRight (Private) - Moves all the items to the right after the index (non-destructive)
func (a *Array) shiftRight(index int) {
	for i := a.size; i > index; i-- {
		a.collection[i] = a.collection[i-1]
	}
}
