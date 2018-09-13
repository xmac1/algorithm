package quicksort

import "math/rand"

func Sort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	q := partition(arr)


	Sort(arr[:q])
	Sort(arr[q+1:])

}

func RandomSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	q := random_partition(arr)


	RandomSort(arr[:q])
	RandomSort(arr[q+1:])
}

func random_partition(arr []int) int {
	i := rand.Intn(len(arr) - 1)

	tmp := arr[len(arr) - 1]
	arr[len(arr) - 1] = arr[i]
	arr[i] = tmp

	return partition(arr)
}

func partition(arr []int) int {
	max := len(arr) -1
	x := arr[max]
	var i = -1

	for j:=0; j < max; j++ {
		if arr[j] <= x {
			i++
			if i != j {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}

	tmp := arr[i+1]
	arr[i+1] = x
	arr[max] = tmp
	return i+1
}