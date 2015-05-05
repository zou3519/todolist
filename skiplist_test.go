package main

import (
	"testing"
)

func SkipListBuilder() Dict {
	return NewSkipList()
}

func Test_SkipList_Insert(t *testing.T) {
	GenericInsertTest(t, SkipListBuilder)
}

func Test_SkipList_Search(t *testing.T) {
	GenericSearchTest(t, SkipListBuilder)
}

func Test_SkipList_Delete(t *testing.T) {
	GenericDeleteTest(t, SkipListBuilder)
}
