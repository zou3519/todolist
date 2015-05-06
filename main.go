package main

import (
	//"fmt"
	"github.com/davecheney/profile"
	"math/rand"
	"time"
)

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	rand.Seed(time.Now().UTC().UnixNano())
	N := 100000
	ExpInsert(NewTodoList(), N)
	ExpSearch(NewTodoList(), N)
	ExpDelete(NewTodoList(), N)
	// ExpInsert(NewMapSet(), N)
	// ExpSearch(NewMapSet(), N)
	// ExpDelete(NewMapSet(), N)
	// ExpInsert(NewLLRB(), N)
	// ExpSearch(NewLLRB(), N)
	// ExpDelete(NewLLRB(), N)
	// ExpInsert(NewLinkedTodoList(), N)
	// ExpSearch(NewLinkedTodoList(), N)
	// ExpDelete(NewLinkedTodoList(), N)

	// tl := NewLinkedTodoList()
	// nums := rand.Perm(1)
	// fmt.Println(nums)
	// for _, v := range nums {
	// 	tl.Insert(v, v)
	// 	fmt.Println(tl.String())
	// }
	// nums = rand.Perm(1)
	// fmt.Println(nums)
	// for _, v := range nums {
	// 	tl.Delete(v)
	// 	fmt.Println(tl.String())
	// }
}
