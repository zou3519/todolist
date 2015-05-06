package main

import (
	"fmt"
)

type MapSet struct {
	Map map[int]interface{}
}

func NewMapSet() *MapSet {
	m := make(map[int]interface{})
	return &MapSet{Map: m}
}

func MapSetBuilder() Dict {
	return NewMapSet()
}

func (set *MapSet) Insert(key int, value interface{}) {
	set.Map[key] = value
}

func (set *MapSet) Search(key int) (interface{}, bool) {
	value, ok := set.Map[key]
	if !ok {
		return nil, ok
	} else {
		return value, true
	}
}

func (set *MapSet) Delete(key int) (interface{}, bool) {
	elt, ok := set.Map[key]
	if ok {
		delete(set.Map, key)
		return elt, true
	} else {
		return nil, false
	}
}

func (set *MapSet) String() string {
	return fmt.Sprintln(set.Map)
}
