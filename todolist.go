package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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

	// txt := "nil"
	// if next != nil {
	// 	txt = fmt.Sprintf("%v", next.key)
	// }
	// fmt.Printf("L[%v]: %v -> %v created\n", i, node.key, txt)
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
		// fmt.Println(i, tl.lengths[i], math.Pow(2.-tl.epsilon, float64(i)))
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
	if float64(tl.lengths[height]) >= math.Ceil(math.Pow(2.0-tl.epsilon, float64(height))) {
		// fmt.Println("New Layer!")
		tl.newLayer()
	}

	// now, do partial rebuilding if there is more than 1 thing in L_0
	if tl.lengths[0] > 1 {
		// fmt.Println("Rebuild L_0 cond!")
		tl.fixFirstLayer()
	}
}

func (tl *TodoList) Delete(key int) (value interface{}, ok bool) {
	path := tl.findPredecessors(key)
	height := tl.height

	//fmt.Println("Delete", key)

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

		// tNext may or may not be the found node
		tNext := predecessor.Next(i)

		// case 1
		if tNext == nil {
			predecessor.SetNext(i, successorNode)
			if successorNode != nil {
				tl.lengths[i] += 1
			}

			// case 2
		} else if successorNode != nil && tNext.key == successorNode.key {
			// do nothing

			// cases 3-5
		} else if tNext.key == foundNode.key {
			tNextNext := tNext.Next(i)
			// case 3
			if tNextNext == nil {
				predecessor.SetNext(i, successorNode)
				if successorNode == nil {
					tl.lengths[i] -= 1
				}
				// case 5
			} else if tNextNext.key == successorNode.key {
				predecessor.SetNext(i, successorNode)
				tl.lengths[i] -= 1
				// case 4
			} else {
				predecessor.SetNext(i, successorNode)
				successorNode.SetNext(i, tNextNext)
			}
			// case 6, pre -> -> -> 11
		} else {
			predecessor.SetNext(i, successorNode)
			successorNode.SetNext(i, tNext)
			tl.lengths[i] += 1
		}
	}

	//fmt.Println("after sky high", tl.String())
	// check to see if need to delete layers (h = depth - 1)
	if float64(tl.lengths[height]) < math.Ceil(math.Pow(2.0-tl.epsilon, float64(height-1))) {
		//fmt.Println("Remove Layer!")
		tl.removeLayer()
		//fmt.Println("after remove layer", tl.String())
	}

	// now, do partial rebuilding if there is more than 1 thing in L_0
	if tl.lengths[0] > 1 {
		//fmt.Println("Rebuild L_0 cond!")
		tl.fixFirstLayer()
		//fmt.Println("after rebuild L_0", tl.String())
	}

	return foundNode.value, true
}

func (tl *TodoList) DebugString() string {
	result := "TodoList (debug view)\n"
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

func main() {
	fmt.Printf("Hello, world!\n")

	tl := NewTodoList()
	fmt.Println(tl.String())
	// tl.Insert(0, 1)
	// fmt.Println(tl.String())
	// tl.Insert(2, 1)
	// fmt.Println(tl.String())
	// tl.Insert(3, 1)
	// fmt.Println(tl.String())
	// tl.Insert(4, 4)
	// fmt.Println(tl.String())
	// tl.Insert(5, 1)
	// fmt.Println(tl.String())
	// tl.Insert(6, 1)
	// fmt.Println(tl.String())
	// tl.Insert(7, 1)
	// fmt.Println(tl.String())
	// tl.Delete(6)
	// fmt.Println(tl.String())
	// tl.Delete(5)
	// fmt.Println(tl.String())
	// tl.Delete(7)
	// fmt.Println(tl.String())
	// tl.Delete(1)
	// fmt.Println(tl.String())
	// tl.Delete(0)
	// fmt.Println(tl.String())
	nums1 := []int{11, 0, 6, 9, 7, 4, 2, 3, 1, 8, 10, 5}
	nums2 := []int{9, 7, 6, 8, 3, 4, 11, 1, 0, 10, 5, 2}
	rand.Seed(time.Now().UTC().UnixNano())
	// for i := 0; i < 10; i++ {
	// 	nums := rand.Perm(12)
	// 	fmt.Println(nums)
	// 	for _, v := range nums {
	// 		tl.Insert(v, v)
	// 		//fmt.Println(tl.String())
	// 	}
	// 	nums = rand.Perm(12)
	// 	fmt.Println(nums)
	// 	for _, v := range nums {
	// 		tl.Delete(v)
	// 		// fmt.Println(tl.String())
	// 	}
	// }
	// nums = rand.Perm(12)
	// fmt.Println(nums)
	for _, v := range nums1 {
		tl.Insert(v, v)
		fmt.Println(tl.String())
	}
	// nums = rand.Perm(12)
	// fmt.Println(nums)
	for _, v := range nums2 {
		tl.Delete(v)
		fmt.Println(tl.String())
	}

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
