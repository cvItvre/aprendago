package cap07

import (
	"fmt"
)

func ex07() {
	idade := 20
	if idade >= 18 {
		fmt.Println("Pode dirigir")
	} else if idade >= 16 {
		fmt.Println("Nos eua voce pode dirigir!")
	} else {
		fmt.Println("Voce é apenas uma criança")
	}
}
