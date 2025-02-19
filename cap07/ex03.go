package cap07

import (
	"fmt"
)

const (
	currentYear int = 2025
	bornYear    int = 1998
)

func ex03() {
	time := bornYear
	for time <= currentYear {
		fmt.Println(time)
		time++
	}
}
