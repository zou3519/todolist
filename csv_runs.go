package main

import (
	"encoding/csv"
	"fmt"
	// "math"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"
)

func TodolistEpsilonGraphs(reps int) {
	inserts := 1000000
	searches := 500000

	times := make([][]string, 0, 16)

	for epsilon := 0.02; epsilon <= 0.7; epsilon += 0.01 {
		fmt.Println(epsilon)
		insert_time := 0.
		search_time := 0.

		d := NewTodoList(epsilon)

		//inserting
		for r := 0; r < reps; r++ {
			perm := rand.Perm(inserts)
			start := time.Now()
			for _, v := range perm {
				d.Insert(v, v)
			}
			insert_time += (time.Since(start)).Seconds()

			perm = rand.Perm(inserts)
			start = time.Now()
			for j := 0; j < searches; j++ {
				d.Search(j)
			}
			search_time += (time.Since(start)).Seconds()
		}
		insert_time = insert_time / float64(reps)
		search_time = search_time / float64(reps)
		times = append(times, []string{fmt.Sprint(epsilon), fmt.Sprint(insert_time), fmt.Sprint(search_time)})
	}

	d := NewTodoList(0.1)
	d_type := strings.Replace(reflect.TypeOf(d).String(), "*main.", "", 1)
	filename := fmt.Sprintf("Outputs/epsilon%s.csv", d_type)
	fmt.Println(filename)
	csvfile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)
	err = writer.WriteAll(times) // flush everything into csvfile
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func ExpInsert(d Dict, N int) {
	Exper(d, N, "Insert")
}

func ExpSearch(d Dict, N int) {
	Exper(d, N, "Search")
}

func ExpDelete(d Dict, N int) {
	Exper(d, N, "Delete")
}

func ExpAll(DB dictBuilder, N int, reps int) {

	// multiplier := 1.1
	// trials := int(math.Log2(float64(N)) / math.Log2(multiplier)) //log (base multiplier) of n
	// times := make([][]string, trials + 1)
	// min_N := int(math.Max(float64(N / 100), 10))

	difference := N / 50
	min_N := difference
	trials := (N-min_N)/difference + 1
	times := make([][]string, trials+1)

	var insert_time float64
	var search_time float64
	var delete_time float64
	var perm []int
	var start time.Time
	var d Dict

	times[0] = []string{"Number of Data", "Insert", "Search", "Delete"}
	count := 1
	// for n := min_N; n < N; n = int(math.Ceil(float64(n) * multiplier)) {
	for n := min_N; n <= N; n += difference {
		fmt.Println(n)
		insert_time = 0
		search_time = 0
		delete_time = 0
		for r := 0; r < reps; r++ {
			d = DB()

			//inserting
			perm = rand.Perm(n)
			start = time.Now()
			for j := 0; j < n; j++ {
				d.Insert(perm[j], perm[j])
			}
			insert_time += (time.Since(start)).Seconds()

			//searching
			perm = rand.Perm(n)
			start = time.Now()
			for j := 0; j < n; j++ {
				d.Search(perm[j])
			}
			search_time += (time.Since(start)).Seconds()

			//deleting
			perm = rand.Perm(n)
			start = time.Now()
			for j := 0; j < n; j++ {
				d.Delete(perm[j])
			}
			delete_time += (time.Since(start)).Seconds()
		}
		insert_time = insert_time / float64(reps)
		search_time = search_time / float64(reps)
		delete_time = delete_time / float64(reps)

		times[count] = []string{fmt.Sprintf("%v", n), fmt.Sprintf("%v", insert_time),
			fmt.Sprintf("%v", search_time), fmt.Sprintf("%v", delete_time)}
		count++
	}

	d_type := strings.Replace(reflect.TypeOf(d).String(), "*main.", "", 1)
	h, m, s := time.Now().Local().Clock()
	filename := fmt.Sprintf("Outputs/output%sAll%dTime%d%d%d.csv", d_type, N, h, m, s)
	fmt.Println(filename)
	csvfile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)
	err = writer.WriteAll(times) // flush everything into csvfile
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}

func Exper(d Dict, N int, operation string) {
	reps := 5

	// multiplier := 1.1
	// trials := int(math.Log2(float64(N)) / math.Log2(multiplier)) //log (base multiplier) of n
	// times := make([][]string, trials)
	// min_N := int(math.Max(float64(N / 100), 10))

	difference := N / 50
	min_N := difference
	trials := (N-min_N)/difference + 1
	times := make([][]string, trials)

	var elapsed float64
	count := 0
	switch operation {
	case "Insert":
		// for n := min_N; n < N; n = int(math.Ceil(float64(n) * multiplier)) {
		for n := min_N; n < N; n += difference {
			elapsed = 0
			for r := 0; r < reps; r++ {
				perm := rand.Perm(n)
				start := time.Now()
				for j := 0; j < n; j++ {
					d.Insert(perm[j], perm[j])
				}
				elapsed += (time.Since(start)).Seconds()
			}

			times[count] = []string{fmt.Sprintf("%v", n),
				fmt.Sprintf("%v", elapsed/float64(reps))}
			count++
		}
	case "Delete":
		// for n := min_N; n < N; n = int(math.Ceil(float64(n) * multiplier)) {
		for n := min_N; n < N; n += difference {
			elapsed = 0
			for r := 0; r < reps; r++ {
				perm_ins := rand.Perm(n)
				for j := 0; j < n; j++ {
					d.Insert(perm_ins[j], perm_ins[j])
				}

				perm := rand.Perm(n)
				start := time.Now()
				for j := 0; j < n; j++ {
					d.Delete(perm[j])
				}
				elapsed += (time.Since(start)).Seconds()
			}

			times[count] = []string{fmt.Sprintf("%v", n),
				fmt.Sprintf("%v", elapsed/float64(reps))}
			count++
		}
	case "Search":
		// for n := min_N; n < N; n = int(math.Ceil(float64(n) * multiplier)) {
		for n := min_N; n < N; n += difference {
			elapsed = 0
			for r := 0; r < reps; r++ {
				perm_ins := rand.Perm(n)
				for j := 0; j < n; j++ {
					d.Insert(perm_ins[j], perm_ins[j])
				}

				perm := rand.Perm(n)
				start := time.Now()
				for j := 0; j < n; j++ {
					d.Search(perm[j])
				}
				elapsed += (time.Since(start)).Seconds()
			}

			times[count] = []string{fmt.Sprintf("%v", n),
				fmt.Sprintf("%v", elapsed/float64(reps))}
			count++
		}
	default:
		{
			fmt.Println("Not a valid operation!")
		}
	}

	d_type := strings.Replace(reflect.TypeOf(d).String(), "*main.", "", 1)
	h, m, s := time.Now().Local().Clock()
	filename := fmt.Sprintf("Outputs/output%s%sTime%d%d%d.csv", d_type, operation, h, m, s)
	fmt.Println(filename)
	csvfile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)
	err = writer.WriteAll(times) // flush everything into csvfile
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}
