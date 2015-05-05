package main

import (
	"github.com/petar/GoLLRB/llrb"
)

type LLRB struct {
	tree *llrb.LLRB
}

type Number struct {
	nkey int
	nval interface{}
}

func NewLLRB() *LLRB {
	return &LLRB{tree: llrb.New()}
}

func (a Number) Less(b llrb.Item) bool { return a.nkey < b.(Number).nkey }

func (t *LLRB) Search(key int) (interface{}, bool) {
	item_got := t.tree.Get(Number{nkey: key, nval: 0})
	if item_got == nil {
		return nil, false
	}
	return item_got.(Number).nval, true
}

func (t *LLRB) Delete(key int) (interface{}, bool) {
	item_got := t.tree.Delete(Number{nkey: key, nval: 0})
	if item_got == nil {
		return nil, false
	}
	return item_got.(Number).nval, true
}

func (t *LLRB) Insert(key int, value interface{}) {
	t.tree.ReplaceOrInsert(Number{nkey: key, nval: value})
	return
}

func (t *LLRB) String() string {
	return "LLRB"
}
