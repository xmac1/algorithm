package countsort

func CountSort(arr []int, sorted []int, max int) {
	locArr := make([]int, max+1)
	for _, a := range arr {
		locArr[a]++
	}

	for i := 1; i < max+1; i++ {
		locArr[i] += locArr[i-1]
	}

	for i := len(arr) - 1; i >= 0; i-- {
		sorted[locArr[arr[i]]-1] = arr[i]
		locArr[arr[i]]--
	}

	return
}
