package main

import (
	"testing"
)

func LinkedTodoListBuilder() Dict {
	return NewLinkedTodoList()
}

func Test_LinkedTodo_Insert(t *testing.T) {
	GenericInsertTest(t, LinkedTodoListBuilder)
}

func Test_LinkedTodo_Search(t *testing.T) {
	GenericSearchTest(t, LinkedTodoListBuilder)
}

func Test_LinkedTodo_Delete(t *testing.T) {
	GenericDeleteTest(t, LinkedTodoListBuilder)
}
