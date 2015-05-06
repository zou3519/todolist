package main

import (
	"testing"
)

func Test_TodoList2_Insert(t *testing.T) {
	GenericInsertTest(t, TodoList2Builder(0.2))
}

func Test_TodoList2_Search(t *testing.T) {
	GenericSearchTest(t, TodoList2Builder(0.2))
}

func Test_TodoList2_Delete(t *testing.T) {
	GenericDeleteTest(t, TodoList2Builder(0.2))
}
