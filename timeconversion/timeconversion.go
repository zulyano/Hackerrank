package timeconversion

import (
	"fmt"
	"strconv"
	"strings"
)

func TimeConversion(s string) string {
	delimiter := ":"
	parts := strings.Split(s, delimiter)

	formatted := ""

	if len(parts) == 3 {
		if parts[2][2] == 'A' {
			if parts[0] == "12" {
				formatted = "00:" + parts[1] + ":" + parts[2][:2]
			} else {
				formatted = parts[0] + ":" + parts[1] + ":" + parts[2][:2]
			}
		} else {
			if parts[0] == "12" {
				formatted = "12:" + parts[1] + ":" + parts[2][:2]
			} else {
				num, _ := strconv.Atoi(parts[0])
				formatted = strconv.Itoa(num+12) + ":" + parts[1] + ":" + parts[2][:2]
			}
		}
	}

	fmt.Println(formatted)

	return formatted
}
