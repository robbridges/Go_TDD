package ints

func AddSum(sum1 int, sum2 int) int {
	return sum1 + sum2
}

func AddAllNums(intSlice []int) int {
	num := 0
	for _, v := range intSlice {
		num += v
	}
	return num
}
