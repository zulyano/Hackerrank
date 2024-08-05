package plusminus

import "fmt"

func PlusMinus(arr []int32) {
	total := float64(len(arr))
	positive := 0.0
	negative := 0.0
	zero := 0.0

	for _, num := range arr {
		if num == 0 {
			zero += 1
		} else if num > 0 {
			positive += 1
		} else if num < 0 {
			negative += 1
		}
	}

	fmt.Println(positive / total)
	fmt.Println(negative / total)
	fmt.Println(zero / total)
}
