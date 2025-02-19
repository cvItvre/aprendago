package cap08

import (
	"fmt"
)

func class07() {

	ss := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	fmt.Println(ss)
	fmt.Println(ss[1])
	fmt.Println(ss[1][1])

	matrix := make([][]int, 5, 10)
	fmt.Println(matrix)         // [[] [] [] [] []]
	fmt.Println(matrix[0])      // []
	fmt.Println(len(matrix))    // 5
	fmt.Println(len(matrix[0])) // 0
	fmt.Println(cap(matrix))    // 10
	fmt.Println(cap(matrix[0])) // 0

	slices := [][][][]int{
		[][][]int{
			[][]int{
				[]int{1, 2, 3, 4, 5, 6},
			},
			[][]int{
				[]int{10, 20, 30, 40, 50},
			},
		},

		[][][]int{
			[][]int{
				[]int{11, 13, 15, 17, 19},
			},
			[][]int{
				[]int{-3},
			},
		},
	}

	fmt.Println(slices)

}
