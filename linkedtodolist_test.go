package main

import (
	"testing"
)

func Test_LinkedTodo_Insert(t *testing.T) {
	GenericInsertTest(t, LinkedTodoListBuilder(0.2))
}

func Test_LinkedTodo_Search(t *testing.T) {
	GenericSearchTest(t, LinkedTodoListBuilder(0.2))
}

func Test_LinkedTodo_Delete(t *testing.T) {
	GenericDeleteTest(t, LinkedTodoListBuilder(0.2))
}
