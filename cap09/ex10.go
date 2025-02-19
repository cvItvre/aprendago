package cap09

import (
	"fmt"
)

func ex10() {
	personToHobby := map[string][]string{
		"jr_neymar": {
			"jogar bola", "festas", "fazer filhos",
		},
		"senna_airton": {
			"ganhar", "comer loiras", "morrer",
		},
	}

	personToHobby["araujo_pedro"] = []string{
		"transar o cabelo", "comer", "jogar",
	}

	for k, v := range personToHobby {
		fmt.Printf("A pessoa %v tem o hobby:\n", k)
		for _, c := range v {
			fmt.Println("\t", c)
		}
	}

	delete(personToHobby, "jr_neymar")
	for k, v := range personToHobby {
		fmt.Printf("A pessoa %v tem o hobby:\n", k)
		for _, c := range v {
			fmt.Println("\t", c)
		}
	}
}
