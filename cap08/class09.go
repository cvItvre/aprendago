package cap08

import (
	"fmt"
)

func class09() {

	ctt := map[string]int{
		"alfred": 551234,
		"joana":  997342,
	}

	fmt.Println(ctt)
	fmt.Println(ctt["joana"])

	ctt["karol"] = 978730
	fmt.Println(ctt["karol"])

	fmt.Println(ctt["pedro"]) // 0 - inexistente ou valor real? -> comma ok idiom

	num, ok := ctt["pedro"]
	fmt.Println(num, ok) // 0 false

	//ctt["pedro"] = 875506

	//aplying
	if value, ok := ctt["pedro"]; !ok {
		fmt.Println("nao tem!!!!")
	} else {
		fmt.Println(value)
	}

}
