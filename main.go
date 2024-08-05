package main

import (
	"hackerrank/minimaxsum"
	// "hackerrank/plusminus"
)

func main() {

	// expected output: 0.4, 0.4, 0.2
	arr := []int32{1, 1, 0, -1, -1}
	// plusminus.PlusMinus(arr)

	// expected output: 10, 14
	arr = []int32{256741038, 623958417, 467905213, 714532089, 938071625}
	minimaxsum.MinMaxSum(arr)
}
