package cap16

import (
	"fmt"
	"slices"
	"sort"
)

func class05() {

	ss := []string{"pera", "uva", "banana", "jaca", "morango"}
	fmt.Println(ss)
	sort.Strings(ss) // as of Go 1.22, this function simply calls slices.Sort
	// slices.Sort(ss)
	fmt.Println(ss)

	si := []int{4, 7, 2, 9, 1, 0, 3}
	fmt.Println(si)
	//sort.Ints(si) // as of Go 1.22, this function simply calls slices.Sort
	slices.Sort(si)
	fmt.Println(si)

}
