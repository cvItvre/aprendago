package cap13

import (
	"fmt"
)

func ex03() {

	defer fmt.Println("=====INICIANDO PROGRAMA DE CONTAGEM=====")

	for i := 0; i < 100; i++ {
		fmt.Println("Contando...", i)
	}

}
