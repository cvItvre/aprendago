package cap06

import (
	"fmt"
)

func class03() {

	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Hora:", horas)
		for x := 0; x < 60; x++ {
			fmt.Print(x, " ")
		}
		fmt.Println("")
	}

}
