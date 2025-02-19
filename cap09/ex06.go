package cap09

import (
	"fmt"
)

func ex06() {
	estados := make([]string, 26, 27)

	estados[0] = "Pernambuco"
	estados[1] = "Bahia"
	estados[2] = "Paraiba"
	estados[3] = "Rio Grande do Norte"
	estados[4] = "Maranhão"

	fmt.Println(estados, len(estados), cap(estados))

	for i := 0; i < len(estados); i++ {
		fmt.Println(estados[i])
	}
}
