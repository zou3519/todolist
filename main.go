package main

import (
	_ "fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	N := 300000
	ExpInsert(NewTodoList(), N)
	ExpSearch(NewTodoList(), N)
	ExpDelete(NewTodoList(), N)
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
