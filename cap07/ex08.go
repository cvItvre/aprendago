package cap07

import (
	"fmt"
)

func ex08() {
	sexo := "NB"
	switch {
	case sexo == "M":
		fmt.Println("bem-vindo")
	case sexo == "F":
		fmt.Println("bem-vinda")
	default:
		fmt.Println("OLHA A CULTURA WOKE BUUUUUUUUU")
	}
}
