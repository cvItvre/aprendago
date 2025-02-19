package cap06

import (
	"fmt"
)

func class06() {

	for i := 33; i <= 122; i++ {
		fmt.Printf("%d\t%#x\t%#U\t%v\n", i, i, i, string(rune(i)))
	}

}
