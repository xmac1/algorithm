package randomize

import (
	"math/rand"
	"time"
)

func Shuffle(arr []int) {
	l := len(arr)

	seed := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < l-1; i++ {

		rn := seed.Intn(l-i-1) + i
		tmp := arr[i]
		arr[i] = arr[rn]
		arr[rn] = tmp
	}
	return
}
