package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Welcome to Go LinkedList")
	ll := LinkedList{}

	// Check if LinkedList is empty
	fmt.Println(ll.is_empty()) // true
	// Check the length of the LinkedList
	fmt.Println(ll.length()) // 0
	// Add nodes LinkedList
	ll.prepend(10)
	fmt.Println(ll.length()) // 1
	ll.prepend(9)
	fmt.Println(ll.length()) // 2
	ll.printLinkedList()     // [9 10]
	// Add a node at the end of the LinkedList
	ll.append(7)
	ll.printLinkedList() // [9 10 7]

	ll.insert(54, 1)
	ll.printLinkedList() // [9 54 10 7]
}

// Create a node : A node should have a data and a pointer to the next node if any else points to nil.
type Node struct {
	data      int   // Data that the node will contain
	next_node *Node // Pointer to the next node
}

// Create a LinkedList : A collection of nodes linked with pointers.
// A LinkedList must have a head of type Node to define the starting point.
type LinkedList struct {
	head *Node // The starting point of a linked list
}

// Check if a LinkedList is empty
func (ll LinkedList) is_empty() bool {
	/*
		Checks if a LinkedList is empty or not
		Returns a boolean true if empty else returns false
	*/
	return ll.head == nil // if LinkedList head is nill, the LinkedList is empty
}

// Get the length of a LinkedList
func (ll LinkedList) length() int {
	/*
		Returns the length(int) of items in a LinkedList
		This a linear operation 0(n)
	*/
	count := 0 //Initial length of a LinkedList

	current := ll.head // Keep track of the current node in a LinkedList

	// Loop through LinkedList until we get to the tail node and return count
	for current != nil {
		// As long as the next_node is not nil
		// Increament count
		count += 1
		// Move current node to the next node
		current = current.next_node
	}

	return count
}

// Add a node at the beginning of a LinkedList
func (ll *LinkedList) prepend(data int) {
	/*
		Adds a new node at the start of a LinkedList
		Takes data of type int, adds data at the start of the LinkedList
	*/

	// Create a node to add in the LinkedList
	new_node := Node{data: data}

	// Check if the head of the LinkedList is nil.
	// Add new node as head node of the LinkedList
	if ll.head == nil {
		ll.head = &new_node
	} else {
		// LinkedList has existing nodes, add node to the existing nodes in the LinkedList
		new_node.next_node = ll.head // Link new node to the head of the LinkedList
		ll.head = &new_node          // Define new head for the LinkedList
	}
}

// Print the contents of the LinkedList
func (ll LinkedList) printLinkedList() {
	/*
		Prints the content of a linkedlist
	*/
	// Define linkedlist to store the final content of a LinkedList
	linkedList := []int{}

	// Define the current node, to keep track of nodes
	current := ll.head

	// While loop , as long as the current node is not the tail node
	for current != nil {
		// Add node data into linkedList
		linkedList = append(linkedList, current.data)
		// Move to the next node in the LinkdedList
		current = current.next_node
	}
	// Log contents of the linkedList
	fmt.Println(linkedList)
}

// Add node at the end of a LinkedList
func (ll LinkedList) append(data int) {
	/*
		Adds a new node at the end of a LinkedList
		Takes data of type int, added the data at the end of a LinkedList
		Linear operation, 0(n)
	*/
	// Create new node
	new_node := Node{data: data}

	// Create current node to keep track of all nodes as we traverse the LinkedList
	current := ll.head

	// While current.next_node is not nil, continue looping
	for current.next_node != nil {
		// Move to the next node until we get to the end of the LinkedList
		current = current.next_node
	}
	// At this point we at the last node
	current.next_node = &new_node
}

// Add a new node at a given index
func (ll *LinkedList) insert(data, index int) {
	/*
		Adds a new node at a given index position
		data : data/item to add into the LinkedList
		index : Position to insert the new data in the node
	*/
	// Check if index == 0 and add data at the beginning of the LinkedList
	if index == 0 {
		ll.append(data)
	} else {
		// Check if index is out of bound
		if index > ll.length() {
			log.Fatal("Index out of bound")
			return
		}
		// Create new node to insert into the LinkedList
		new_node := Node{data: data}

		// Keep track of nodes in LinkedList
		current := ll.head

		position := index // keep track of the position in the LinkedList

		// Loop until position is 1, at this point index has been reached
		for position < 1 {
			position -= 1
			current = current.next_node
		}
		// Define node before current node
		prev_node := current
		// Create the next node to attach new node before it
		next_node := current.next_node
		// Link prev_node with the new node
		prev_node.next_node = &new_node
		// Link new node with the next node
		new_node.next_node = next_node
	}
}
