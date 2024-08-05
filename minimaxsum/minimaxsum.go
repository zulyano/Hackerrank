package minimaxsum

import "fmt"

func MinMaxSum(arr []int32) {
	var max int64
	var min int64

	for i := 0; i < len(arr); i++ {
		sum := sumWithoutIndex(arr, i)
		if i == 0 {
			max = sum
			min = sum
		} else {
			if sum > max {
				max = sum
			}

			if sum < min {
				min = sum
			}
		}

	}

	fmt.Print(int64(min))
	fmt.Print(" ")
	fmt.Print(int64(max))
}

func sumWithoutIndex(arr []int32, idx int) int64 {
	var store []int64
	for i, num := range arr {
		if i != idx {
			store = append(store, int64(num))
		}
	}

	var result int64
	for _, new := range store {
		result += new
	}

	return result
}
