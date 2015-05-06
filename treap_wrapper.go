package main

import (
	"stathat.com/c/treap"
)

type Treap struct {
	tree *treap.Tree
}

func IntLess(p, q interface{}) bool {
    return p.(int) < q.(int)
}

func NewTreap() *Treap {
	return &Treap{tree: treap.NewTree(IntLess)}
}

func TreapBuilder() Dict {
	return NewTreap()
}

func (t *Treap) Insert(key int, value interface{}) {
	t.tree.Insert(key, value)
	return
}

func (t *Treap) Search(key int) (interface{}, bool) {
	item_got := t.tree.Get(key)
	if item_got != nil {
		return item_got, true
	} else {
		return nil, false
	}
}

func (t *Treap) Delete(key int) (interface{}, bool) {
	t.tree.Delete(key)
	return nil, true
}


func (t *Treap) String() string {
	return "Treap"
}
