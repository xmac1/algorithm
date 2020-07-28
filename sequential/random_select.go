package sequential

import "math/rand"

// 找出数组中第i小的元素
func RandomSelect(arr []int, i int) int {
	if len(arr) <= 0 {
		return 0

	}

	if len(arr) <= 1 {
		return arr[0]
	}

	q := randomPartition(arr)

	left := arr[:q]
	if len(left) == i-1 {
		return arr[q]
	}

	if len(left) > i-1 {
		return RandomSelect(left, i)
	} else {
		return RandomSelect(arr[q+1:], i-q-1)
	}

}

func randomPartition(arr []int) (q int) {
	r := rand.Intn(len(arr))
	pivot := arr[r]
	arr[r] = arr[len(arr)-1]
	arr[len(arr)-1] = pivot

	i, j := 0, -1
	for ; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			j++
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	j++
	arr[j], arr[len(arr)-1] = arr[len(arr)-1], arr[j]

	return j
}
