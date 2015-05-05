package main

import (
	"testing"
)

func LLRBBuilder() Dict {
	return NewLLRB()
}

func Test_LLRB_Insert(t *testing.T) {
	GenericInsertTest(t, LLRBBuilder)
}

func Test_LLRB_Search(t *testing.T) {
	GenericSearchTest(t, LLRBBuilder)
}

func Test_LLRB_Delete(t *testing.T) {
	GenericDeleteTest(t, LLRBBuilder)
}
