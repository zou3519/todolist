package main

// dict_test.go contains generic functions for testing dict data structures

import (
	"math"
	"math/rand"
	"testing"
	"time"
	"fmt"
)

// GenericInsertTest tests random inserts against dicts. May not be comprehensive.
func GenericInsertTest(t *testing.T, build dictBuilder) {
	rand.Seed(time.Now().UTC().UnixNano())
	repeat := 7
	sizes := []int{1, 10, 100, 1000, 10000}

	for _, size := range sizes {
		for iter := 0; iter < repeat; iter++ {
			d := build()
			nums := rand.Perm(size)
			for _, v := range nums {
				d.Insert(v, v)
			}
		}
	}
}

// GenericSearchTest tests random searches against dicts. May not be comprehensive.
func GenericSearchTest(t *testing.T, build dictBuilder) {
	rand.Seed(time.Now().UTC().UnixNano())
	repeat := 5
	sizes := []int{1, 10, 100, 1000, 10000}

	for _, size := range sizes {
		for iter := 0; iter < repeat; iter++ {
			d := build()
			nums := rand.Perm(size)
			for _, v := range nums {
				d.Insert(v, v)
			}

			// now, perform the searches
			// first, searches for things that aren't here
			for ns := 0; ns < repeat; ns++ {
				key := rand.Intn(size) + size
				_, ok := d.Search(key)
				if ok != false {
					t.Error("Search for non-existent element returned true")
				}
			}

			for search := 0; search < repeat; search++ {
				s_nums := rand.Perm(size)
				n := rand.Intn(size)
				for i := 0; i < n; i++ {
					key := s_nums[i]
					a, ok := d.Search(key)
					if ok != true {
						t.Error("Search for existent element returned false")
					}
					if a != key {
						t.Error("Search for key returned wrong value")
						fmt.Println(a)
					}
				}
			}
		}
	}
}

// GenericDeleteTest tests random deletes against dicts. May not be comprehensive.
func GenericDeleteTest(t *testing.T, build dictBuilder) {
	rand.Seed(time.Now().UTC().UnixNano())
	repeat := 5
	sizes := []int{1, 10, 100, 1000}

	// first, the insert all then delete all tests (in random permutation)
	for _, size := range sizes {
		for iter := 0; iter < repeat; iter++ {
			d := build()
			nums := rand.Perm(size)
			for _, v := range nums {
				d.Insert(v, v)
			}

			// try deleting everything
			nums = rand.Perm(size)
			for _, v := range nums {
				d.Delete(v)
			}

			// d should be empty at this point
		}
	}

	// now, try inserts then deletes then inserts
	for _, size := range sizes {
		for iter := 0; iter < repeat; iter++ {
			d := build()
			for i := 0; i < repeat; i++ {
				nums := rand.Perm(size)
				for _, v := range nums {
					d.Insert(v, v)
				}

				// try deleting everything
				nums = rand.Perm(size)
				for _, v := range nums {
					d.Delete(v)
				}
			}
		}
	}
}

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
}
