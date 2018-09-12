package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

func main() {
	var arr []int

	for i := 0; i < 1000000; i++ {
		arr = append(arr, rand.Intn(1000000))
	}

	start  := time.Now()
	sort.Ints(arr)
	dur := time.Since(start)
	fmt.Print(dur.String())
}
