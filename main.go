package main

import (
	//"fmt"
	// "github.com/davecheney/profile"
	// "math/rand"
	// "time"
)

func main() {

	N := 1000000
	ExpInsert(NewSkipList(), N)
	ExpSearch(NewSkipList(), N)
	ExpDelete(NewSkipList(), N)

}
