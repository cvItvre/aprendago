package cap07

import (
	"fmt"
)

func ex02() {
	for i := 65; i < 91; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\t", i)
		}
		fmt.Println("")
	}
}
