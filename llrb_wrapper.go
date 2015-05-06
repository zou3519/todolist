package main

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

type LLRB struct {
	tree *redblacktree.Tree
}

func NewLLRB() *LLRB {
	return &LLRB{tree: redblacktree.NewWithIntComparator()}
}

func (t *LLRB) Search(key int) (interface{}, bool) {
	return t.tree.Get(key)
}

func (t *LLRB) Delete(key int) (interface{}, bool) {
	t.tree.Remove(key)
	return nil, true
}

func (t *LLRB) Insert(key int, value interface{}) {
	t.tree.Put(key, value)
	return
}

func (t *LLRB) String() string {
	return "LLRB"
}
