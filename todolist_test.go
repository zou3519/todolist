package main

import (
	"testing"
)

func Test_TodoList_Insert(t *testing.T) {
	GenericInsertTest(t, TodoListBuilder(0.2))
}

func Test_TodoList_Search(t *testing.T) {
	GenericSearchTest(t, TodoListBuilder(0.2))
}

func Test_TodoList_Delete(t *testing.T) {
	GenericDeleteTest(t, TodoListBuilder(0.2))
}
