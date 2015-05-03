package main

import ()

// Dict is an interface for data structures that can Insert, Search, and Delete
type Dict interface {
	// Insert element into the set
	Insert(int, interface{})

	// Search the set for the element by ID.  Return the element and T/F
	Search(int) (interface{}, bool)

	// Delete the element from the set by ID.  Returns the element and T/F
	Delete(int) (interface{}, bool)

	// String gives a string representation
	String() string
}
