package cap06

import (
	"fmt"
)

func class05() {

	for i := 0; i <= 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

}
