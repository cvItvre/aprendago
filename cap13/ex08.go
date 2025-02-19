package cap13

import (
	"fmt"
)

func ex08() {

	x := retornaFunc()
	x()

}

func retornaFunc() func() {

	return func() {
		fmt.Println("Olha eu aqui!")
	}

}
