package main

import (
	// "fmt"
	// "github.com/davecheney/profile"
	"math/rand"
	"time"
	// "sort"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	N := 300000
	ExpInsert(NewLLRB(), N)
	ExpSearch(NewLLRB(), N)
	ExpDelete(NewLLRB(), N)

	ExpAll(LLRBBuilder, N)

	// N := 400000
	// ExpInsert(NewTodoList(0.2), N)
	// ExpSearch(NewTodoList(0.2), N)
	// ExpDelete(NewTodoList(0.2), N)

	// N := 500000
	// ExpAll(SkipListBuilder, N)

	// TodolistEpsilonGraphs()
	// e.g. ExpAll(TodoListBuilder(0.2), N)

	// N := 1000000
	// perm := rand.Perm(N)
	// sort.Ints(perm)
	// ExpSearch(NewOptimalBST(perm), N)

}
