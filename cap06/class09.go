package cap06

import (
	"fmt"
)

func class09() {

	x := 5

	switch {
	case x < 5:
		fmt.Println("x é menor que 5")
	case x == 5:
		fmt.Println("x é igual a 5")
	case x > 5:
		fmt.Println("x é maior que 5")
	default:
		fmt.Println("cade x?")
	}

	inOffice := "karol"

	switch inOffice {
	case "pedro":
		fmt.Println("bem-vindo pedro")
	case "karol":
		fmt.Println("bem-vinda karol")
	default:
		fmt.Println("fudeo, nao tem ninguém!!!")
	}

	itsTwo := 3

	switch itsTwo {
	case 2:
		fmt.Println("sim!!! é dois!")
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fallthrough
	case 6:
		fallthrough
	default:
		fmt.Println("nao é dois... fodase")
	}

	switch inOffice {
	case "pedro", "karol":
		fmt.Println("r305")
	case "gustavo", "kesia":
		fmt.Println("a302")
	}

	switch {
	case (x < 10), (x == 10):
		fmt.Println("1 ou 2")
	case (x > 10):
		fmt.Println("3 ou 4")
	}

}
