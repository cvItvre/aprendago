package cap09

import (
	"fmt"
)

func ex02() {
	slice := [10]int{}

	for k, _ := range slice {
		slice[k] = k + 1
	}

	for k, v := range slice {
		fmt.Println(k, v)
	}

	fmt.Printf("%T", slice)
}
