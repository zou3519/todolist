package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"
)

func ExpInsert(d Dict, N int) {
	Exper(d, N, "Insert")
}

func ExpSearch(d Dict, N int) {
	Exper(d, N, "Search")
}

func ExpDelete(d Dict, N int) {
	Exper(d, N, "Delete")
}

func Exper(d Dict, N int, operation string) {

	multiplier := 1.1
	trials := int(math.Log2(float64(N)) / math.Log2(multiplier)) //log (base multiplier) of n
	times := make([][]string, trials)

	count := 0
	min_N := int(math.Max(float64(N / 100), 10))

	switch operation {
	case "Insert":
		for n := min_N; n < N; n = int(math.Ceil(float64(n) * multiplier)) {
			perm := rand.Perm(n)
			start := time.Now()
			for j := 0; j < n; j++ {
				d.Insert(perm[j], perm[j])
			}
			elapsed := (time.Since(start)).Seconds()

			times[count] = []string{fmt.Sprintf("%v", n), fmt.Sprintf("%v", elapsed)}
			count++
		}
	case "Delete":
		//n is the number of items to use in a given trial
		for n := min_N; n < N; n = int(math.Ceil(float64(n) * multiplier)) {
			perm_ins := rand.Perm(n)
			for j := 0; j < n; j++ {
				d.Insert(perm_ins[j], perm_ins[j])
			}
			perm := rand.Perm(n)
			start := time.Now()
			for j := 0; j < n; j++ {
				d.Delete(perm[j])
			}
			elapsed := (time.Since(start)).Seconds()

			times[count] = []string{fmt.Sprintf("%v", n), fmt.Sprintf("%v", elapsed)}
			count++
		}
	case "Search":
		//n is the number of items to use in a given trial
		for n := min_N; n < N; n = int(math.Ceil(float64(n) * multiplier)) {
			perm_ins := rand.Perm(n)
			for j := 0; j < n; j++ {
				d.Insert(perm_ins[j], perm_ins[j])
			}
			perm := rand.Perm(n)
			start := time.Now()
			for j := 0; j < n; j++ {
				d.Search(perm[j])
			}
			elapsed := (time.Since(start)).Seconds()

			times[count] = []string{fmt.Sprintf("%v", n), fmt.Sprintf("%v", elapsed)}
			count++
		}
	default:
		{
			fmt.Println("Not a valid operation!")
		}
	}

	d_type := strings.Replace(reflect.TypeOf(d).String(), "*main.", "", 1)
	filename := fmt.Sprintf("Outputs/output%s%s.csv", d_type, operation)
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
