package main

import (
	// "llrb"
	// "github.com/zou3519/llrb"
	"github.com/petar/GoLLRB/llrb"
)

type LLRB struct {
	tree *llrb.LLRB
}

type Number struct {
	nkey int
	nval int
}

func NewLLRB() *LLRB {
	return &LLRB{}
}

func (a Number) Less(b llrb.Item) bool { return a.nkey < b.(Number).nkey }

func (t *LLRB) Search(key int) (interface{}, bool) {
	var found bool
	item_got := t.tree.Get(Number{nkey: key, nval: 0})
	if item_got != nil {
		found = true
	}
	return item_got.(Number).nval, found
}

func (t *LLRB) Delete(key int) (interface{}, bool) {
	var found bool
	item_got := t.tree.Delete(Number{nkey: key, nval: 0})
	if item_got != nil {
		found = true
	}
	return item_got.(Number).nval, found
} 

func (t *LLRB) Insert(key int, value interface{}) {
	t.tree.ReplaceOrInsert(Number{nkey: key, nval: value.(int)})
	return
}

func (t *LLRB) String() string {
	return "LLRB"
}