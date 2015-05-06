package main

import (
	"fmt"
	"math"
)

type Node struct {
	key   int
	val interface{}
	left *Node
	right *Node
}

type OptimalBST struct {
	root *Node
	size int
}

func NewOptimalBST(nodes []int) *OptimalBST {
	num_nodes := len(nodes)
	max_height := int(math.Log2(float64(num_nodes + 1)) - 1)
	t := &OptimalBST{}
	median_ind := (num_nodes-1)/2
	t.root = &Node{key: nodes[median_ind], val: nodes[median_ind], 
		left: build(nodes, max_height - 1, median_ind, 0),
		right: build(nodes, max_height - 1, median_ind, 1)}

	return t
}

func build(nodes []int, level int, current int, right int) *Node {
	// fmt.Println(current)
	left_ind := current - int(math.Pow(2, float64(level)))
	right_ind := current + int(math.Pow(2, float64(level)))
	//left
	if right == 0 {
		if level > 0 {
			return &Node{key: nodes[left_ind], val: nodes[left_ind],
				left: build(nodes, level - 1, left_ind, 0),
				right: build(nodes, level - 1, left_ind, 1)}
		} else {
			return &Node{key: nodes[left_ind], val: nodes[left_ind], left: nil, right: nil}
		}
	} else {
		if level > 0 {
			return &Node{key: nodes[right_ind], val: nodes[right_ind],
				left: build(nodes, level - 1, right_ind, 0),
				right: build(nodes, level - 1, right_ind, 1)}
		} else {
			return &Node{key: nodes[right_ind], val: nodes[right_ind], left: nil, right: nil}
		}
	}
}

func (t *OptimalBST) Search(key int) (interface {}, bool){
	for node := t.root; node != nil; {
		if key < node.key {
			node = node.left
		} else if key > node.key {
			node = node.right
		} else {
			return node.val, true
		}
	}
	return nil, false
}

func (t *OptimalBST) Delete(key int) (interface {}, bool){
	return nil, false
}

func (t *OptimalBST) Insert(key int, value interface{}) {
	return
}

func (t *OptimalBST) String() string{
	return "BST"
}

// func (t *OptimalBST) traverse(n *Node, s string) string{
// 	if n.left != nil && n.right != nil {
// 		return s += n.key + find(n.left) + find(n.right)
// 	} else {
// 		return s += n.key
// 	}

// }

func (t *OptimalBST) Print(current *Node) {
	fmt.Printf("%d ", current.val)
	if current.left != nil {
		t.Print(current.left)
	}
	if current.right != nil {
		t.Print(current.right)
	}
}

// n := 15
// perm := rand.Perm(n)
// sort.Ints(perm)
// fmt.Println(perm)
// t := NewOptimalBST(perm)

// for i := 0; i < n; i++ {
// 	fmt.Println(t.Search(perm[i]))
// }
// t.Print(t.root)














