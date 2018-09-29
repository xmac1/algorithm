package insert

func Sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		this := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if this < arr[j] {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = this

	}
}
