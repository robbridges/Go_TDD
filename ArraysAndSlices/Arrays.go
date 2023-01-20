package ArraysAndSlices

func GetSumFromArray(arr []int) int {
	sum := 0
	for _, number := range arr {
		sum += number
	}

	return sum
}

func SumAllSlices(slice [][]int) []int {
	var returnedSlice []int
	for _, value := range slice {
		// zero check so that there is not a panic in the software
		if len(value) == 0 {
			returnedSlice = append(returnedSlice, 0)
		} else {
			returnedSlice = append(returnedSlice, GetSumFromArray(value))
		}
	}
	return returnedSlice
}
