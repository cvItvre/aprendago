package cap08

import (
	"fmt"
)

func class06() {

	slice := make([]int, 5, 10)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	for k, _ := range slice {
		slice[k] = k + 1
	}

	fmt.Println(slice)

	// slice[5] = 6     panic: runtime error: index out of range [5] with length 5

	slice = append(slice, 6)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 11)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}
