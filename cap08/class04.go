package cap08

import (
	"fmt"
)

func class04() {

	sabores := []string{"pepperoni", "muzzarela", "frango com catupiry", "napolitana", "abacaxi"}

	fatia := sabores[0:2]
	fmt.Println(fatia)

	fatia = sabores[1:3]
	fmt.Println(fatia)

	fatia = sabores[2:5]
	fmt.Println(fatia)

	fatia = sabores[2:len(sabores)]
	fmt.Println(fatia)

	fatia = sabores[2:]
	fmt.Println(fatia)

	fatia = sabores[:2]
	fmt.Println(fatia)

	fatia = sabores[:]
	fmt.Println(fatia)

	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}

	sabores = append(sabores[0:2], sabores[3:]...)
	fmt.Println(sabores)

}
