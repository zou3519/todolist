package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// TLNode is a linked list node
type SLNode struct {
	next  []*SLNode
	key   int
	value interface{}
	level int
}

type SkipList struct {
	Sentinel     *SLNode // head of each linked list
	maxheight    int
	height       int     // how deep the sentinel is
	prob         float64 // probability of promoting an element
	num_elements int     // number of elements in the skiplist
}

func NewSLNode(key int, value interface{}, level int) *SLNode {
	node := SLNode{}
	node.key = key
	node.value = value
	//HARD-CODED - WHATT????
	node.level = level
	node.next = make([]*SLNode, level, 16)
	return &node
}

func NewSkipList() *SkipList {
	rand.Seed(time.Now().UTC().UnixNano())

	sl := SkipList{}
	sl.maxheight = 16 //appropriate for N <= (1/p)^(sl.maxheight)
	sl.height = 1
	// sl.Sentinel = make([]*SLNode, sl.maxheight, sl.maxheight)
	sl.Sentinel = &SLNode{next: make([]*SLNode, sl.maxheight, sl.maxheight),
		key: math.MinInt32, value: math.MinInt32}
	sl.prob = 0.5
	sl.num_elements = 0
	return &sl
}

//WHERE to start the search??
func (sl *SkipList) findPredecessors(x int) []*SLNode {
	height := sl.height
	result := make([]*SLNode, height, height)

	//result indexing goes top to bottom
	result[height-1] = sl.Sentinel
	for i := height - 1; i >= 0; i-- {
		if i < height-1 {
			result[i] = result[i+1]
		}
		for result[i].next[i] != nil && result[i].next[i].key < x {
			result[i] = result[i].next[i]
		}
		// fmt.Println(result, i, result[i])
	}

	return result
}

func (sl *SkipList) FullSearch(key int) (value interface{}, ok bool) {
	path := sl.findPredecessors(key)
	x := path[0].next[0]
	if x.key == key {
		return x.value, true
	} else {
		return nil, false
	}
	return
}

func (sl *SkipList) Search(key int) (value interface{}, ok bool) {
	height := sl.height

	//result indexing goes top to bottom
	current := sl.Sentinel
	for i := height - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].key <= key {
			current = current.next[i]
			if current.key == key {
				return current.value, true
			}
		}
	}

	return nil, false
}

func (sl *SkipList) RandLevel() int {
	level := 1
	for rand.Float64() < sl.prob {
		level++
	}
	return level
}

func (sl *SkipList) Insert(key int, value interface{}) {
	//Can probably generate random numbers faster than this
	level := sl.RandLevel()
	fmt.Println("Inserting:", key, "into L", level-1)
	if level > sl.maxheight {
		fmt.Println("OH NOES, it's too high!!!!")
	}
	if level > sl.height {
		sl.height = level
	}

	//splice key into the skiplist
	preds := sl.findPredecessors(key)
	// fmt.Println(sl.height)
	// for i := level - 1; i >= 0; i-- {
	// 	fmt.Println(preds[i].value)
	// 	if preds[i].next[i] == nil {
	// 		fmt.Println("nil")
	// 	} else {
	// 		fmt.Println(preds[i].next[i].value)
	// 	}
	// }
	//update the existing node whose key matches 'key'
	if preds[0].next[0] != nil && preds[0].next[0].key == key {
		preds[0].next[0].value = value
	} else {
		x := NewSLNode(key, value, level)
		for i := 0; i < level; i++ {
			// fmt.Println(i)
			x.next[i] = preds[i].next[i]
			preds[i].next[i] = x
		}
		sl.num_elements++
	}
	return
}

func (sl *SkipList) Delete(key int) (value interface{}, ok bool) {
	fmt.Println("Deleting:", key)

	//splice key into the skiplist
	preds := sl.findPredecessors(key)
	x := preds[0].next[0]
	if x.key == key {
		for i := x.level - 1; i >= 0; i-- {
			//reached the top of this pile
			// fmt.Println(i)
			if preds[i].next[i].key != key {
				break
			} else {
				preds[i].next[i] = x.next[i]
			}
		}
		// free(x)
		for j := sl.height - 1; j >= 0 && sl.Sentinel.next[j] == nil; j-- {
			sl.height--
		}
		sl.num_elements--
		return x.value, true
	} else {
		return nil, false
	}
}

func (sl *SkipList) String() string {
	result := ""
	keys := make([]int, sl.num_elements)
	node := sl.Sentinel.next[0]
	for i := 0; i < sl.num_elements; i++ {
		keys[i] = node.key
		node = node.next[0]
	}

	for i := sl.height - 1; i >= 0; i-- {
		node = sl.Sentinel.next[i]
		str := fmt.Sprintf("L[%v]--", i)

		for j := 0; j < sl.num_elements; j++ {
			if node != nil && node.key == keys[j] {
				str += fmt.Sprintf("%v--", node.key)
				node = node.next[i]
			} else {
				str += "_--"
			}
		}

		result = result + str + "\n"
	}

	result = "LinkedTodoList\n" + result
	return result
}

func (sl *SkipList) test_insert(items []int) {
	for i := 0; i < len(items); i++ {
		sl.Insert(items[i], items[i])
		fmt.Println(sl)
	}
	return
}

func (sl *SkipList) test_search(items []int) {
	for i := 0; i < len(items); i++ {
		a, ok := sl.Search(items[i])
		if ok {
			fmt.Println("Found:", a)
		} else {
			fmt.Println("Couldn't find", items[i], "!")
		}
	}
	return
}

func (sl *SkipList) test_delete(items []int) {
	for i := 0; i < len(items); i++ {
		a, ok := sl.Delete(items[i])
		if ok {
			fmt.Println("Deleted:", a)
			fmt.Println(sl)
		} else {
			fmt.Println("Couldn't delete", items[i], "!")
		}
	}
	return
}

// func main() {
// 	fmt.Printf("Hello, world!\n")

// 	sl := NewSkipList()
// 	sl.test_insert([]int{0, 8, 9, 1, 7, -5, 11, 4, 3, 5, 10, 2, 6})
// 	fmt.Printf("Inserting done\n")

// 	sl.test_search([]int{4, 5, 6, -1, 12, -5})
// 	fmt.Printf("Searching done\n")

// 	sl.test_delete([]int{3, 7, 8, 9, 6, -1, 0})
// 	fmt.Printf("Deleting done\n")

// }
