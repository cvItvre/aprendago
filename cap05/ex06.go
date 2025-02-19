package cap05

import (
	"fmt"
)

const anoAtual = 2025

const (
	_ = anoAtual + iota
	um
	dois
	tres
	quatro
)

func ex06() {
	fmt.Printf("Ano atual: %v\n", anoAtual)
	fmt.Printf("Primeiro ano: %v\n", um)
	fmt.Printf("Segundo ano: %v\n", dois)
	fmt.Printf("Terceiro ano: %v\n", tres)
	fmt.Printf("Quarto ano: %v\n", quatro)
}
