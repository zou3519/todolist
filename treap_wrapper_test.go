package main

import (
	"testing"
)

func Test_Treap_Insert(t *testing.T) {
	GenericInsertTest(t, TreapBuilder)
}

func Test_Treap_Search(t *testing.T) {
	GenericSearchTest(t, TreapBuilder)
}

func Test_Treap_Delete(t *testing.T) {
	GenericDeleteTest(t, TreapBuilder)
}
