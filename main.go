package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func main() {
	// set seed
	rand.Seed(time.Now().UTC().UnixNano())

	validModes := "epsilongraph, redblack, todolist, todolist2, " +
		"linkedtodolist, mapset, treap, skiplist, optimalbst"
	modePtr := flag.String("mode", "<please specify a mode>",
		"The data structure you want to generate statistics for. Valid modes: "+validModes)
	nPtr := flag.Int("n", 100000, "The maximum number of items to insert/search/delete")
	repsPtr := flag.Int("trials", 1, "Multiple trials - results will be averaged over the trials")
	epsPtr := flag.Float64("epsilon", 0.2, "Value of epsilon constant, if data structure is a todolist variant")
	flag.Parse()

	N := *nPtr
	reps := *repsPtr
	mode := *modePtr
	epsilon := *epsPtr

	if mode == "epsilongraph" {
		fmt.Printf("Generating an epsilon graph for TodoList with reps: %v\n", reps)
	} else if strings.HasPrefix(mode, "todolist") || strings.HasSuffix(mode, "todolist") {
		fmt.Printf("Running benchmarks for %v (epsilon: %v) on maximum N: %v, trials for each N: %v\n", mode, epsilon, N, reps)
	} else {
		fmt.Printf("Running benchmarks for %v on maximum N: %v, trials for each N: %v\n", mode, N, reps)
	}

	// make the output directory on current working directory
	os.Mkdir("."+string(filepath.Separator)+"Output", 0777)

	switch mode {
	case "epsilongraph":
		TodolistEpsilonGraphs(reps)
	case "redblack":
		ExpAll(LLRBBuilder, N, reps)
	case "todolist":
		ExpAll(TodoListBuilder(epsilon), N, reps)
	case "todolist2":
		ExpAll(TodoList2Builder(epsilon), N, reps)
	case "linkedtodolist":
		ExpAll(LinkedTodoListBuilder(epsilon), N, reps)
	case "mapset":
		ExpAll(MapSetBuilder, N, reps)
	case "treap":
		ExpAll(TreapBuilder, N, reps)
	case "skiplist":
		ExpAll(SkipListBuilder, N, reps)
	case "optimalbst":
		fmt.Println("reps defaulting to 5 (shouldn't matter for an optimal bst), sorry!")
		perm := rand.Perm(N)
		sort.Ints(perm) // in place
		ExpSearch(NewOptimalBST(perm), N)
	default:
		fmt.Printf("Mode %v not recognized. Terminating.\n", mode)
		fmt.Println("Valid modes are:", validModes)
	}
}
