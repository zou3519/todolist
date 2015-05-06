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

	// N := 1000000
	// ExpInsert(NewSkipList(), N)
	// ExpSearch(NewSkipList(), N)
	// ExpDelete(NewSkipList(), N)

	N := 2000000
	ExpAll(TreapBuilder, N)

	TodolistEpsilonGraphs()
	// e.g. ExpAll(TodoListBuilder(0.2), N)

	// N := 1000000
	// perm := rand.Perm(N)
	// sort.Ints(perm)
	// ExpSearch(NewOptimalBST(perm), N)

}
