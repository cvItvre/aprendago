package cap08

import (
	"fmt"
)

func class02() {

	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	//second_array := append(array, 6) can't do this to an array
	second_slice := append(slice, 6)
	fmt.Println(second_slice)
	fmt.Printf("%T - %T\n", array, slice)

	fmt.Println(slice[3])
	slice[3] = 861523
	fmt.Println(slice[3])

	slice[15] = 1 //index out of range

}
