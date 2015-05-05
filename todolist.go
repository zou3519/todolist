package main

import (
	"fmt"
	"math"
)

type TNode struct {
	value interface{}
	key   int
	next  []*TNode  // List of next pointers from this.  Indexed from the bottom
	tl    *TodoList // TodoList that this is a part of
}

type TodoList struct {
	lengths  []int   // the size of the ith list
	height   int     // the height of the todolist. equal to depth - 1
	n        int     // number of things being held
	epsilon  float64 // constant
	Sentinel TNode   // sentinel[i] is the ith sentinel; access list with .next[0]
}

func NewTodoList() *TodoList {
	result := TodoList{lengths: make([]int, 1, 16)}
	result.Sentinel = TNode{next: make([]*TNode, 1, 16), tl: &result}
	result.lengths[0] = 0
	result.epsilon = 0.2
	return &result
}

func (node *TNode) Next(i int) *TNode {
	// we don't have enough allocation
	if node.tl.height-i+1 > len(node.next) {
		return nil
	}
	return node.next[node.tl.height-i]
}

func (node *TNode) SetNext(i int, next *TNode) {
	// need to allocate more by appending (maybe not the best way to do this)
	if node.tl.height-i+1 > len(node.next) {
		for len(node.next) < node.tl.height-i+1 {
			node.next = append(node.next, nil)
		}
	}
	node.next[node.tl.height-i] = next
}

// findPredecessors returns a list of *TLNodes that immediately proceed
// the key x.  result[i] = node means node is the predecessor to x in L_i
func (tl *TodoList) findPredecessors(x int) []*TNode {
	height := tl.height
	result := make([]*TNode, height+1)

	// find them
	result[0] = &tl.Sentinel // get the first sentinel
	for i := 0; i <= height; i++ {
		nextGuy := result[i].Next(i)
		if nextGuy != nil && nextGuy.key < x {
			result[i] = nextGuy
		}
		if i+1 <= height {
			result[i+1] = result[i] // go down (autoamatically)
		}
	}
	return result
}

// rebuildLayer rebuilds one layer and assumes the other layers on top will also be rebuilt.
func (tl *TodoList) rebuildLayer(i int) {
	length := 0 // new length

	// reference layer
	referenceNode := tl.Sentinel.Next(i + 1)
	iNode := &tl.Sentinel
	iNode.SetNext(i, nil) // delete everything after

	second := false
	for node := referenceNode; node != nil; node = node.Next(i + 1) {
		if second == true {
			iNode.SetNext(i, node)
			iNode = node
			iNode.SetNext(i, nil) // reset
			length++
		}
		second = !second
	}

	// length variable changes too
	tl.lengths[i] = length
}

// rebuild does a partial rebuilding of lists L_i to L_0 modifying L_i as well
func (td *TodoList) rebuild(i int) {
	for k := i; k >= 0; k-- {
		td.rebuildLayer(k)
	}
}

func (tl *TodoList) newLayer() {
	oldHeight := tl.height

	// add new layer for sentinel
	tl.Sentinel.next = append(tl.Sentinel.next, nil)
	tl.height++

	// update lengths
	tl.lengths = append(tl.lengths, tl.lengths[oldHeight])

	// rebuild from second to last layer up
	tl.rebuild(oldHeight)
}

func (tl *TodoList) removeLayer() {
	height := tl.height
	tl.Sentinel.next = tl.Sentinel.next[:height] // removes the last thing (the first layer)
	tl.lengths = tl.lengths[1:]                  // remove the first length
	tl.height--
	tl.rebuild(height - 2)
}

func (tl *TodoList) Search(key int) (value interface{}, ok bool) {
	path := tl.findPredecessors(key)
	height := tl.height
	uh := path[height]
	nextGuy := uh.Next(height)
	if nextGuy != nil && nextGuy.key == key {
		return nextGuy.value, true
	} else {
		return nil, false
	}
}

// fixes the invariant that the first layer must have <= 1 thing in it.
// caller must check if the first layer is indeed violated
func (tl *TodoList) fixFirstLayer() {
	// first, find the smallest index i such that |L_i| <= (2-ep)^i
	i := 0
	for ; float64(tl.lengths[i]) > math.Pow(2.-tl.epsilon, float64(i)); i++ {
		fmt.Println(tl.lengths[i], math.Pow(2.-tl.epsilon, float64(i)))
	}
	if float64(tl.lengths[i]) > math.Pow(2.-tl.epsilon, float64(i)) {
		fmt.Println("Something went wrong! In Insert!")
	}
	tl.rebuild(i - 1)
}

func (tl *TodoList) Insert(key int, value interface{}) {
	path := tl.findPredecessors(key)
	height := tl.height

	// key is already here, update the value
	if next := path[height].Next(height); next != nil && next.key == key {
		next.value = value
		return
	}

	// update internal structures
	tl.n++

	// otherwise, insert (key, value) everywhere right after path
	n := &TNode{key: key, value: value, tl: tl, next: make([]*TNode, height+1)}
	for i := height; i >= 0; i-- {
		n.SetNext(i, path[i].Next(i))
		path[i].SetNext(i, n)
		tl.lengths[i] += 1 // increase length count by 1
	}

	// rebalancing condition
	if float64(tl.lengths[height]) > math.Ceil(math.Pow(2.0-tl.epsilon, float64(height))) {
		fmt.Println("New Layer!")
		tl.newLayer()
	}

	// now, do partial rebuilding if there is more than 1 thing in L_0
	if tl.lengths[0] > 1 {
		fmt.Println("Rebuild L_0 cond!")
		tl.fixFirstLayer()
	}
}

func (tl *TodoList) Delete(key int) (value interface{}, ok bool) {
	path := tl.findPredecessors(key)
	height := tl.height

	foundNode := path[height].Next(height)

	// thing wasn't found in the list
	if foundNode == nil || foundNode.key != key {
		return nil, false
	}

	// thing was found, subtract 1 from n
	tl.n--
	successorNode := foundNode.Next(height) // can be nil

	// destroy found node, create all successors
	for i := 0; i <= height; i++ {
		predecessor := path[i]

		// set the successor
		tNext := predecessor.Next(i)
		predecessor.SetNext(i, successorNode)

		// figure out whether or not to decrement or add for the number
		if tNext != nil && tNext.key == key {
			tl.lengths[i] -= 1

			// make the successor appear
			tNextNext := tNext.Next(i)
			if tNextNext != nil && tNextNext.key != successorNode.key {
				successorNode.SetNext(i, tNextNext)
				tl.lengths[i] += 1
			}
		}
		if tNext == nil {
			tl.lengths[i] += 1
		}

	}

	// check to see if need to delete layers (h = depth - 1)
	if float64(tl.lengths[height]) < math.Ceil(math.Pow(2.0-tl.epsilon, float64(height-1))) {
		fmt.Println("Remove Layer!")
		tl.removeLayer()
	}

	// now, do partial rebuilding if there is more than 1 thing in L_0
	if tl.lengths[0] > 1 {
		fmt.Println("Rebuild L_0 cond!")
		tl.fixFirstLayer()
	}

	return foundNode.value, true
}

func (tl *TodoList) String() string {
	result := "TodoList\n"
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
				build += "   --"
			} else {
				build += fmt.Sprintf("%3d--", v)
				node = node.Next(i)
			}
		}
		result += build + fmt.Sprintf(" (%v)", tl.lengths[i]) + "\n"
	}

	return result
}

func main() {
	fmt.Printf("Hello, world!\n")

	tl := NewTodoList()
	fmt.Println(tl.String())
	tl.Insert(0, 1)
	fmt.Println(tl.String())
	tl.Insert(2, 1)
	fmt.Println(tl.String())
	tl.Insert(3, 1)
	fmt.Println(tl.String())
	tl.Insert(4, 4)
	fmt.Println(tl.String())
	tl.Insert(5, 1)
	fmt.Println(tl.String())
	tl.Insert(6, 1)
	fmt.Println(tl.String())
	tl.Insert(7, 1)
	fmt.Println(tl.String())
	tl.Delete(6)
	fmt.Println(tl.String())
	tl.Delete(5)
	fmt.Println(tl.String())
	tl.Delete(7)
	fmt.Println(tl.String())
	tl.Delete(1)
	fmt.Println(tl.String())
	tl.Delete(0)
	fmt.Println(tl.String())

	// a, ok := tl.Search(4)
	// if ok {
	// 	fmt.Println("Search returned", a)
	// } else {
	// 	fmt.Println("Alert!")
	// }
	// a, ok = tl.Search(8)
	// if !ok {
	// 	fmt.Println("Search did not return")
	// } else {
	// 	fmt.Println("Alert!")
	// }
	// fmt.Println(tl.String())
}
