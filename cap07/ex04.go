package cap07

import (
	"fmt"
)

func ex04() {
	futureYear := 2025
	currentYear := 1998
	for {
		fmt.Println(currentYear)
		if currentYear == futureYear {
			break
		}
		currentYear++
	}
}
