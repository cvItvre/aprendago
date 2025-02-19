package cap13

import (
	"fmt"
)

func ex01() {
	fmt.Println(retornaNum())
	ok, err := doisRetornos()
	fmt.Println(ok, err)
}

func retornaNum() int {
	return 2
}

func doisRetornos() (int, string) {
	return 0, "err"
}
