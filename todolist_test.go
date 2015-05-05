// Copyright 2010 Petar Maymounkov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"math"
	"math/rand"
	"testing"
)

// rand.Seed(time.Now().UTC().UnixNano())



func benchmarkRandomPerm(n int, b *testing.B) {
	rand.Perm(n)
}

func benchmarkRandomNums(n int, b *testing.B) {
	b.StopTimer()
	arr := make([]int, n, n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			arr[j] = rand.Intn(n)
		}
	}
}

var N = int(math.Pow(10, 6))

func BenchmarkInsert(b *testing.B) {
	benchmarkRandomInsert(NewSkipList(), N, b)
}
func BenchmarkSearch(b *testing.B) {
	benchmarkRandomSearch(NewSkipList(), N, b)
}
func BenchmarkDelete(b *testing.B) {
	benchmarkRandomDelete(NewSkipList(), N, b)
}

func benchmarkSeqInsert(d Dict, n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			d.Insert(j, j)
		}
	}
}
func benchmarkRandomInsert(d Dict, n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		perm := rand.Perm(n)
		b.StartTimer()
		for j := 0; j < n; j++ {
			d.Insert(perm[j], perm[j])
		}
	}
}



func benchmarkRandomSearch(d Dict, n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a1 := rand.Perm(n)
		a2 := rand.Perm(n)
		for j := 0; j < n; j++ {
			d.Insert(a1[j], a1[j])
		}
		b.StartTimer()
		for j := 0; j < n; j++ {
			d.Search(a2[j])
		}
	}
}

func benchmarkRandomDelete(d Dict, n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a1 := rand.Perm(n)
		a2 := rand.Perm(n)
		for j := 0; j < n; j++ {
			d.Insert(a1[j], a1[j])
		}
		b.StartTimer()
		for j := 0; j < n; j++ {
			d.Delete(a2[j])
		}
	}

func TodoListBuilder() Dict {
	return NewTodoList()
}

func Test_TodoList_Insert(t *testing.T) {
	GenericInsertTest(t, TodoListBuilder)
}

func Test_TodoList_Search(t *testing.T) {
	GenericSearchTest(t, TodoListBuilder)
}

func Test_TodoList_Delete(t *testing.T) {
	GenericDeleteTest(t, TodoListBuilder)
}
