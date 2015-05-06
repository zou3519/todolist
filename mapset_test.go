package main

import (
	"testing"
)

func Test_MapSet_Insert(t *testing.T) {
	GenericInsertTest(t, MapSetBuilder)
}

func Test_MapSet_Search(t *testing.T) {
	GenericSearchTest(t, MapSetBuilder)
}

func Test_MapSet_Delete(t *testing.T) {
	GenericDeleteTest(t, MapSetBuilder)
}
