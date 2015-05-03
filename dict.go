package main

import ()

// Comparable is an interface for things that are comparable to each other.
type Comparable interface {
	// Returns 0 if equal, > 0 if this is greater than the other
	// < 0 if the other is greater
	CompareTo(Comparable) int
}

// Element is to be stored inside a Set
type Element struct {

	// The value stored with this element.
	Value *Comparable
	// contains filtered or unexported fields
}

// Dict is an interface for data structures that can Insert, Search, and Delete
type Dict interface {
	// Insert element into the set
	Insert(Element)

	// Search the set for the element.
	Search(Element) error

	// Delete the element from the set
	Delete(Element) error
}
