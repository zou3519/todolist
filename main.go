package main

import (
	"fmt"
	_ "math/rand"
	"time"
)

func main() {
	nums := []int{92, 4, 2, 49, 47, 32, 90, 95, 35, 85, 41} // rand.Perm(100)
	start := time.Now()

	var d Dict = NewTodoList()
	for _, v := range nums {
		fmt.Println("Inserting", v)
		d.Insert(v, v)
		fmt.Println(d.String())
	}

	elapsed := time.Since(start)
	fmt.Println("Took: %s", elapsed)
}
