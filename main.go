package main

import (
	// "fmt"
	// "math/rand"
	// "time"
)

func main() {
	// rand.Seed(time.Now().UTC().UnixNano())

	// nums := rand.Perm(100000)
	// start := time.Now()

	// var d Dict = NewLLRB()
	// for _, v := range nums {
	// 	//fmt.Println("Inserting", v)
	// 	d.Insert(v, v)
	// 	//fmt.Println(d.String())
	// 	//fmt.Println(d.(*TodoList).DebugString())
	// }
	// elapsed := time.Since(start)
	// fmt.Printf("inserts took: %s\n", elapsed)

	// start = time.Now()
	// nums = rand.Perm(100000)
	// for _, v := range nums {
	// 	//fmt.Println("Inserting", v)
	// 	_, _ = d.Search(v)
	// 	//fmt.Println(d.String())
	// 	//fmt.Println(d.(*TodoList).DebugString())
	// }
	// elapsed = time.Since(start)
	// fmt.Printf("searches took: %s\n", elapsed)

	// fmt.Println(d.String())

	N := 1000000
	ExpInsert(NewSkipList(), N)
	ExpSearch(NewSkipList(), N)
	ExpDelete(NewSkipList(), N)

}

