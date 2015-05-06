package main

import (
	"fmt"
	"math"
)

type TContainer struct {
	node *TNode2
	key  int
}

type TNode2 struct {
	value interface{}
	key   int
	next  []TContainer // List of next pointers from this.  Indexed from the bottom
	tl    *TodoList2   // TodoList2 that this is a part of
}

type TodoList2 struct {
	lengths    []int     // the size of the ith list
	height     int       // the height of the todolist. equal to depth - 1
	n          int       // number of things being held
	epsilon    float64   // constant
	Sentinel   TNode2    // sentinel[i] is the ith sentinel; access list with .next[0]
	limits     []float64 // limits[i] = (2-epsilon)^h
	ceilLimits []float64 // ceil_limits[i] = ceil((2-epsilon)^h)
}

func NewTodoList2() *TodoList2 {
	result := TodoList2{lengths: make([]int, 1, 16)}
	result.Sentinel = TNode2{next: make([]TContainer, 1, 16), tl: &result}
	result.lengths[0] = 0
	result.epsilon = 0.2
	return &result
}

func TodoList2Builder() Dict {
	return NewTodoList2()
}

func (tl *TodoList2) lengthLimit(h int, ceil bool) float64 {
	for len(tl.limits)-1 < h {
		newH := float64(len(tl.limits))
		x := math.Pow(2.-tl.epsilon, newH)
		tl.limits = append(tl.limits, x)
		tl.ceilLimits = append(tl.ceilLimits, math.Ceil(x))
	}
	if ceil {
		return tl.ceilLimits[h]
	} else {
		return tl.limits[h]
	}
}

func (node *TNode2) Next(i int) *TNode2 {
	// we don't have enough allocation
	if node.tl.height-i+1 > len(node.next) {
		return nil
	}
	return node.next[node.tl.height-i].node
}

func (node *TNode2) NextKey(i int) int {
	// we don't have enough allocation
	if node.tl.height-i+1 > len(node.next) || node.next == nil {
		fmt.Println("Something went wrong! NextKey")
		return 0
	}
	return node.next[node.tl.height-i].key
}

func (node *TNode2) SetNext(i int, next *TNode2) {
	// need to allocate more by appending (maybe not the best way to do this)
	if node.tl.height-i+1 > len(node.next) {
		for len(node.next) < node.tl.height-i+1 {
			node.next = append(node.next, TContainer{})
		}
	}
	node.next[node.tl.height-i].node = next
	if next != nil {
		node.next[node.tl.height-i].key = next.key
	}

	// txt := "nil"
	// if next != nil {
	//  txt = fmt.Sprintf("%v", next.key)
	// }
	// fmt.Printf("L[%v]: %v -> %v created\n", i, node.key, txt)
}

// findPredecessors returns a list of *TLNodes that immediately proceed
// the key x.  result[i] = node means node is the predecessor to x in L_i
func (tl *TodoList2) findPredecessors(x int) []*TNode2 {
	height := tl.height
	result := make([]*TNode2, height+1)

	// find them
	result[0] = &tl.Sentinel // get the first sentinel
	for i := 0; i <= height; i++ {
		nextGuy := result[i].Next(i)
		if nextGuy != nil && result[i].NextKey(i) < x {
			result[i] = nextGuy
		}
		if i+1 <= height {
			result[i+1] = result[i] // go down (autoamatically)
		}
	}
	return result
}

// rebuildLayer rebuilds one layer and assumes the other layers on top will also be rebuilt.
func (tl *TodoList2) rebuildLayer(i int) {
	length := 0 // new length

	// reference layer
	referenceNode := tl.Sentinel.Next(i + 1)
	iNode := &tl.Sentinel

	// delete the connections in the current layer
	for node := iNode; node != nil; {
		next := node.Next(i)
		node.SetNext(i, nil)
		node = next
	}

	second := false
	for node := referenceNode; node != nil; node = node.Next(i + 1) {
		// if we're a second thing, then build
		if second == true {
			iNode.SetNext(i, node)
			iNode = node
			length++
			//fmt.Print(node.key, " ")
		}
		second = !second
	}
	//fmt.Print(node.key, "\n")

	// length variable changes too
	tl.lengths[i] = length
}

// rebuild does a partial rebuilding of lists L_i to L_0 modifying L_i as well
func (td *TodoList2) rebuild(i int) {
	for k := i; k >= 0; k-- {
		td.rebuildLayer(k)
	}
}

func (tl *TodoList2) newLayer() {
	oldHeight := tl.height

	// add new layer for sentinel
	tl.Sentinel.next = append(tl.Sentinel.next, TContainer{})
	tl.height++

	// update lengths
	tl.lengths = append(tl.lengths, tl.lengths[oldHeight])

	// rebuild from second to last layer up
	tl.rebuild(oldHeight)
}

func (tl *TodoList2) removeLayer() {
	height := tl.height
	tl.Sentinel.next = tl.Sentinel.next[:height] // removes the last thing (the first layer)
	tl.lengths = tl.lengths[1:]                  // remove the first length
	tl.height--
	tl.rebuild(height - 2)
}

func (tl *TodoList2) Search(key int) (value interface{}, ok bool) {
	path := tl.findPredecessors(key)
	height := tl.height
	uh := path[height]
	nextGuy := uh.Next(height)
	if nextGuy != nil && uh.NextKey(height) == key {
		return nextGuy.value, true
	} else {
		return nil, false
	}
}

// fixes the invariant that the first layer must have <= 1 thing in it.
// caller must check if the first layer is indeed violated
func (tl *TodoList2) fixFirstLayer() {
	// first, find the smallest index i such that |L_i| <= (2-ep)^i
	i := 0
	for ; float64(tl.lengths[i]) > tl.lengthLimit(i, false); i++ {
		// fmt.Println(i, tl.lengths[i], math.Pow(2.-tl.epsilon, float64(i)))
	}
	tl.rebuild(i - 1)
}

func (tl *TodoList2) Insert(key int, value interface{}) {
	path := tl.findPredecessors(key)
	height := tl.height

	// key is already here, update the value
	if next := path[height].Next(height); next != nil && path[height].NextKey(height) == key {
		next.value = value
		return
	}

	// update internal structures
	tl.n++

	// otherwise, insert (key, value) everywhere right after path
	n := &TNode2{key: key, value: value, tl: tl, next: make([]TContainer, height+1)}
	for i := height; i >= 0; i-- {
		n.SetNext(i, path[i].Next(i))
		path[i].SetNext(i, n)
		tl.lengths[i] += 1 // increase length count by 1
	}

	// rebalancing condition
	if float64(tl.lengths[height]) >= tl.lengthLimit(height, true) {
		// fmt.Println("New Layer!")
		tl.newLayer()
	}

	// now, do partial rebuilding if there is more than 1 thing in L_0
	if tl.lengths[0] > 1 {
		// fmt.Println("Rebuild L_0 cond!")
		tl.fixFirstLayer()
	}
}

func (tl *TodoList2) Delete(key int) (value interface{}, ok bool) {
	path := tl.findPredecessors(key)
	height := tl.height

	foundNode := path[height].Next(height)

	// thing wasn't found in the list
	if foundNode == nil || path[height].NextKey(height) != key {
		return nil, false
	}

	// thing was found, subtract 1 from n
	tl.n--
	successorNode := foundNode.Next(height) // can be nil

	// destroy found node, create all successors
	for i := 0; i <= height; i++ {
		predecessor := path[i]

		// tNext may or may not be the found node
		tNext := predecessor.Next(i)

		// now, there are 3 cases of the node configuration, with many subcases
		// let pre be the previous node, 9 be the found node, and suc be
		// the successor node. Keep in mind that the successor node can be nil
		// case 1: pre -> nil
		// case 2: pre -> suc ->
		// case 3: pre -> 9 -> nil
		// case 4: pre -> 9 -> 11 (11 > suc)
		// case 5: pre -> 9 -> suc -> 11
		// case 6: pre -> 11

		if tNext == nil { // case 1
			predecessor.SetNext(i, successorNode)
			if successorNode != nil {
				tl.lengths[i] += 1
			}
		} else if successorNode != nil && predecessor.NextKey(i) == successorNode.key {
			// case 2, do nothing

		} else if predecessor.NextKey(i) == key { // cases 3-5
			tNextNext := tNext.Next(i)

			if tNextNext == nil { // case 3
				predecessor.SetNext(i, successorNode)
				if successorNode == nil {
					tl.lengths[i] -= 1
				}
			} else if tNext.NextKey(i) == successorNode.key { // case 5
				predecessor.SetNext(i, successorNode)
				tl.lengths[i] -= 1
			} else { // case 4
				predecessor.SetNext(i, successorNode)
				successorNode.SetNext(i, tNextNext)
			}
		} else { // case 6
			predecessor.SetNext(i, successorNode)
			successorNode.SetNext(i, tNext)
			tl.lengths[i] += 1
		}
	}

	// check to see if need to delete layers (h = depth - 1)
	if float64(tl.lengths[height]) < tl.lengthLimit(height-1, true) {
		tl.removeLayer()
	}

	// now, do partial rebuilding if there is more than 1 thing in L_0
	if tl.lengths[0] > 1 {
		tl.fixFirstLayer()
	}

	return foundNode.value, true
}

func (tl *TodoList2) DebugString() string {
	result := "TodoList2 (debug view)\n"
	// for each L_i, print out stuff
	for i := 0; i <= tl.height; i++ {
		build := fmt.Sprintf("L[%2d]--", i)
		for node := tl.Sentinel.Next(i); node != nil; node = node.Next(i) {
			build += fmt.Sprintf("%v--", node.key)
		}
		result += build + fmt.Sprintf(" (%v)", tl.lengths[i]) + "\n"
	}
	return result
}
func (tl *TodoList2) String() string {
	result := "TodoList2\n"
	keys := make([]int, tl.n)

	// populate keys
	count := 0
	for node := tl.Sentinel.Next(tl.height); node != nil; node = node.Next(tl.height) {
		keys[count] = node.key
		count++
	}

	// for each L_i, print out stuff
	for i := 0; i <= tl.height; i++ {
		build := fmt.Sprintf("L[%2d]--", i)
		node := tl.Sentinel.Next(i)
		for _, v := range keys {
			if node == nil || v != node.key {
				// create string with digits equal to the number of digits in v
				digits := int(math.Floor(math.Log10(float64(v)))) + 1
				if v == 0 {
					digits = 1
				}
				for c := 0; c < digits; c++ {
					build += " "
				}
				build += "--"
			} else {
				build += fmt.Sprintf("%d--", v)
				node = node.Next(i)
			}
		}
		result += build + fmt.Sprintf(" (%v)", tl.lengths[i]) + "\n"
	}

	return result
}
