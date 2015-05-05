package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	nums := rand.Perm(100000)
	start := time.Now()

	var d Dict = NewLinkedTodoList()
	for _, v := range nums {
		//fmt.Println("Inserting", v)
		d.Insert(v, v)
		//fmt.Println(d.String())
		//fmt.Println(d.(*TodoList).DebugString())
	}

	elapsed := time.Since(start)
	fmt.Printf("Took: %s\n", elapsed)
	// fmt.Println(d.String())
}
