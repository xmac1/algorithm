package main

import (
	"fmt"

	"math/rand"
	"time"

	"github.com/xmac1/example/sequential"
)

func main() {
	n := 2200000
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(n))
	}

	start := time.Now()
	fmt.Println(sequential.RandomSelect(arr, 2200001))
	dur := time.Since(start)
	fmt.Println(dur.String())
	fmt.Println(arr)
}
