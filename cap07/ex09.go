package cap07

import (
	"fmt"
)

func ex09() {
	esporteFavorito := "legends"
	switch esporteFavorito {
	case "futebol":
		fmt.Println("heterotop")
	case "nfl":
		fmt.Println("caiu nas graças do capitalismo americano")
	case "xadrez":
		fmt.Println("nerd")
	case "legends":
		fmt.Println("esses transam")
	default:
		fmt.Println("krl viado")
	}
}
