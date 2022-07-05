package lib

import (
	"fmt"
)

func GenerateTimeDisplay(count uint64) string {
	var display string

	mill := count % 1000
	count /= 1000

	seconds := count % 60
	count /= 60

	minutes := count

	if minutes > 0 {
		if minutes < 10 {
			display += "0"
		}
		display += fmt.Sprintf("%d:", minutes)
	}

	if seconds < 10 {
		display += "0"
	}
	display += fmt.Sprintf("%d.", seconds)

	if mill < 10 {
		display += "0"
	}
	if mill < 100 {
		display += "0"
	}
	display += fmt.Sprintf("%d", mill)
	return display
}
