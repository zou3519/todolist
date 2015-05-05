package main

import (
	"fmt"
	"math"
)

// TLNode is a linked list node
type TLNode struct {
	next *TLNode
	down *TLNode
	key  int
	elt  interface{}
}

type LinkedTodoList struct {
	Sentinel []*TLNode // head of each linked list
	lengths  []int     // length of each linked list
	epsilon  float64   // epsilon
	// depth    int       // how deep the sentinel is
}

func NewLinkedTodoList() *LinkedTodoList {
	ltl := LinkedTodoList{}
	ltl.Sentinel = make([]*TLNode, 1, 16)
	ltl.lengths = make([]int, 1, 16)
	ltl.Sentinel[0] = &TLNode{}
	ltl.lengths[0] = 0
	ltl.epsilon = 0.2
	return &ltl
}

func (lst *LinkedTodoList) findPredecessors(x int) []*TLNode {
	depth := len(lst.Sentinel)
	result := make([]*TLNode, depth)

	// find them
	result[0] = lst.Sentinel[0]
	for i := 0; i < depth; i++ {
		if result[i].next != nil && result[i].next.key < x {
			result[i] = result[i].next
		}
		if result[i] != nil && result[i].down != nil {
			result[i+1] = result[i].down
		}
	}
	return result
}

// rebuildLayer rebuilds one layer and assumes the other layers on top will also be rebuilt.
func (lst *LinkedTodoList) rebuildLayer(i int) {
	// reference layer
	uk := lst.Sentinel[i+1]
	ui := lst.Sentinel[i]
	ui.next = nil

	length := 0

	second := false
	for node := uk.next; node != nil; node = node.next {
		if second == true {
			ui.next = &TLNode{down: node, key: node.key, elt: node.elt}
			ui = ui.next
			length++
		}
		second = !second
	}

	// length variable changes too
	lst.lengths[i] = length
}

// rebuild does a partial rebuilding of lists L_i to L_0 modifying L_i as well
func (lst *LinkedTodoList) rebuild(i int) {
	// the last sentinel
	// ui := lst.Sentinel[i]
	for k := i; k >= 0; k-- {
		lst.rebuildLayer(k)
	}
}

func (lst *LinkedTodoList) newLayer() {
	depth := len(lst.Sentinel)
	uold := lst.Sentinel[depth-1]

	unew := &TLNode{}

	// connections
	lst.Sentinel[depth-1] = unew
	unew.down = uold
	uold.down = nil
	if depth >= 2 {
		lst.Sentinel[depth-2].down = unew
	}

	lst.Sentinel = append(lst.Sentinel, uold)
	lst.lengths = append(lst.lengths, lst.lengths[depth-1])
	lst.rebuild(depth - 1)
}

func (lst *LinkedTodoList) removeLayer() {
	depth := len(lst.Sentinel)
	ubottom := lst.Sentinel[depth-1]
	if depth >= 2 {
		lst.Sentinel[depth-3].down = ubottom
	}
	lst.Sentinel[depth-2] = ubottom
	lst.lengths[depth-2] = lst.lengths[depth-1]
	lst.lengths = lst.lengths[:depth-1]
	lst.Sentinel = lst.Sentinel[:depth-1] // reduce size
	lst.rebuild(len(lst.Sentinel) - 2)
}

func (lst *LinkedTodoList) Search(key int) (value interface{}, ok bool) {
	path := lst.findPredecessors(key)
	uh := path[len(path)-1]
	nextGuy := uh.next
	if nextGuy.key == key {
		return nextGuy.elt, true
	} else {
		return nil, false
	}
}

func (lst *LinkedTodoList) Delete(key int) (value interface{}, ok bool) {
	path := lst.findPredecessors(key)
	depth := len(lst.Sentinel)

	// thing wasn't found in the list
	foundNode := path[len(path)-1].next
	if foundNode == nil || foundNode.key != key {
		return nil, false
	}

	successorNode := foundNode.next

	// destroy all the found nodes and add sucessor nodes
	var prev *TLNode = nil
	for i := len(path) - 1; i >= 0; i-- {
		predecessor := path[i]

		// perform deletion
		if predecessor.next != nil && predecessor.next.key == key {
			predecessor.next = predecessor.next.next
			lst.lengths[i] -= 1
		}

		// now, add in successor where it should be (right after predecessor)
		if successorNode != nil {
			// successor node is present
			if predecessor.next != nil && predecessor.next.key == successorNode.key {
				prev = predecessor.next
			} else {
				// successor node is not present
				newNode := TLNode{next: predecessor.next, down: prev,
					key: successorNode.key, elt: successorNode.elt}
				predecessor.next = &newNode
				prev = &newNode
				lst.lengths[i] += 1
			}
		}
	}

	// check to see if need to delete layers (h = depth - 1)
	if float64(lst.lengths[depth-1]) < math.Ceil(math.Pow(2.0-lst.epsilon, float64(depth-2))) {
		fmt.Println("Remove Layer!")
		lst.removeLayer()
	}

	// rebalance TODO: WRAP IN FUNCTION
	// now, do partial rebuliding if there is more than 1 thing in L_0
	if lst.lengths[0] > 1 {
		// first, find the smallest index i such that |L_i| <= (2-ep)^i
		i := 0
		for ; float64(lst.lengths[i]) > math.Pow(2.-lst.epsilon, float64(i)); i++ {
			// fmt.Println(i, math.Pow(2.-lst.epsilon, float64(i)))
		}
		if float64(lst.lengths[i]) > math.Pow(2.-lst.epsilon, float64(i)) {
			fmt.Println("Something went wrong! In Insert!")
		}
		lst.rebuild(i - 1)
	}
	return foundNode.elt, true

}

func (lst *LinkedTodoList) Insert(key int, value interface{}) {
	path := lst.findPredecessors(key)
	depth := len(lst.Sentinel)

	// key is already here
	if next := path[depth-1].next; next != nil && next.key == key {
		next.elt = value
		return
	}

	// otherwise, insert (key, value) everywhere right after path
	var prev *TLNode = nil
	for i := depth - 1; i >= 0; i-- {
		n := TLNode{key: key, elt: value, down: prev, next: path[i].next}
		prev = &n
		path[i].next = &n
		lst.lengths[i] += 1 // increase length count by 1
	}

	// check to see if need to add more layers
	if float64(lst.lengths[depth-1]) >= math.Ceil(math.Pow(2.0-lst.epsilon, float64(depth-1))) {
		// fmt.Println("New Layer!")
		lst.newLayer()
	}

	// now, do partial rebuliding if there is more than 1 thing in L_0
	if lst.lengths[0] > 1 {
		// first, find the smallest index i such that |L_i| <= (2-ep)^i
		i := 0
		for ; float64(lst.lengths[i]) > math.Pow(2.-lst.epsilon, float64(i)); i++ {
			// fmt.Println(i, math.Pow(2.-lst.epsilon, float64(i)))
		}
		if float64(lst.lengths[i]) > math.Pow(2.-lst.epsilon, float64(i)) {
			fmt.Println("Something went wrong! In Insert!")
		}
		lst.rebuild(i - 1)
	}
}

func (lst *LinkedTodoList) String() string {
	depth := len(lst.Sentinel)
	result := ""

	maxLength := lst.lengths[depth-1]
	keys := make([]int, maxLength)

	for i := depth - 1; i >= 0; i-- {
		ui := lst.Sentinel[i]
		str := fmt.Sprintf("L[%2d]--", i)

		if i == depth-1 {
			count := 0
			for j := ui.next; j != nil; j = j.next {
				keys[count] = j.key
				count++
			}
		}

		j := ui.next
		for count := 0; count < maxLength; count++ {
			// keys[count] = j.key

			if j != nil && j.key == keys[count] {
				str += fmt.Sprintf("%d--", j.key)
				j = j.next
			} else {
				// create string with digits equal to the number of digits in v
				digits := int(math.Floor(math.Log10(float64(keys[count])))) + 1
				for c := 0; c < digits; c++ {
					str += " "
				}
				str += "--"
			}
		}

		result = str + fmt.Sprintf(" (%d)", lst.lengths[i]) + "\n" + result
	}

	result = "LinkedTodoList\n" + result
	return result
}

// func main() {
// 	fmt.Printf("Hello, world!\n")

// 	ltl := NewLinkedTodoList()

// 	ltl.Insert(8, 8)
// 	ltl.Insert(9, 9)
// 	ltl.Insert(1, 1)
// 	ltl.Insert(7, 7)
// 	ltl.Insert(11, 11)
// 	ltl.Insert(4, 4)
// 	ltl.Insert(3, 3)

// 	a, ok := ltl.Search(4)
// 	if ok {
// 		fmt.Println("Search returned", a)
// 	} else {
// 		fmt.Println("Alert!")
// 	}
// 	a, ok = ltl.Search(5)
// 	if !ok {
// 		fmt.Println("Search did not return")
// 	} else {
// 		fmt.Println("Alert!")
// 	}
// 	fmt.Println(ltl.String())

// 	ltl.Delete(3)
// 	fmt.Println(ltl)
// 	ltl.Delete(7)
// 	fmt.Println(ltl)
// 	ltl.Insert(7, 7)
// 	fmt.Println(ltl)
// 	ltl.Delete(7)
// 	fmt.Println(ltl)
// 	ltl.Delete(8)
// 	fmt.Println(ltl)
// 	ltl.Delete(9)
// 	fmt.Println(ltl)
// 	// var m Dict = NewMapSet()
// 	//m.Insert(1, "stuff")
// 	//fmt.Println(m)

// 	var _ Dict = NewLLRB()

// }
