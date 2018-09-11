package main

import (
	"github.com/xmac1/example/heap"
	"fmt"
)

func main() {
	arr := []int{3, 6, 8, 1, 2, 5}

	hp := heap.BuildFromArray(arr)
	fmt.Print(hp)
}
