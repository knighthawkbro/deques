package main

import (
	"deques/array"
	"deques/list"
	"fmt"
)

// Deque (Private) - Interface functionality for deques
type deque interface {
	AddFirst(T interface{})
	AddLast(T interface{})

	RemoveFirst() interface{}
	RemoveLast() interface{}

	GetFirst() interface{}
	GetLast() interface{}

	Contains(T interface{}) bool
	Size() int
	String() string
}

func main() {
	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning driver function as an array...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	array := array.New()
	driver(array)

	fmt.Println("\n*************************************************")
	fmt.Print("*\tRunning driver function as a list...")
	fmt.Println("\n*************************************************")
	fmt.Println("")
	list := list.New()
	driver(list)
}

func driver(d deque) {

	// Adding to front (empty list)
	fmt.Println("\nadding 'hello' to front")
	d.AddFirst("hello")
	fmt.Println(d)

	// Adding to front
	fmt.Println("\nadding 'well' to front")
	d.AddFirst("well")
	fmt.Println(d)

	// Adding to end
	fmt.Println("\nadding 'world' to end")
	d.AddLast("world")
	fmt.Println(d)

	// Retrieving from front
	fmt.Println("\ngetting first item")
	fmt.Println(d.GetFirst())

	// Retrieving from end
	fmt.Println("\ngetting last item")
	fmt.Println(d.GetLast())

	// Removing first item
	fmt.Println("\nremoving first item")
	d.RemoveFirst()
	fmt.Println(d)

	// Removing last item
	fmt.Println("\nremoving last item")
	d.RemoveLast()
	fmt.Println(d)

	// Removing last item (list of one)
	fmt.Println("\nremoving last item")
	d.RemoveLast()
	fmt.Println(d)

	// Removing last item (empty list)
	fmt.Println("\ntrying to removing last item")
	fmt.Println(d.RemoveLast())

	// Adding to end (empty list)
	fmt.Println("\nadding 'hello' to end")
	d.AddLast("hello")
	fmt.Println(d)

	// Removing first item (list of one)
	fmt.Println("\nremoving first item")
	d.RemoveFirst()
	fmt.Println(d)

	// removing first item (empty list)
	fmt.Println("\ntrying to remove first item")
	fmt.Println(d.RemoveFirst())
}
