package cap09

import (
	"fmt"
)

func ex08() {
	personToHobby := map[string]string{
		"araujo_pedro":  "YouTube",
		"souza_karol":   "TikTok",
		"person_unknow": "none",
	}

	for k, v := range personToHobby {
		fmt.Printf("A pessoa %v tem o hobby %v\n", k, v)
	}

	data := map[string][]string{
		"jr_neymar": []string{
			"jogar bola", "festas", "fazer filhos",
		},
		"senna_airton": []string{
			"ganhar", "comer loiras", "morrer",
		},
	}

	for k, v := range data {
		fmt.Printf("A pessoa %v tem o hobby:\n", k)
		for _, c := range v {
			fmt.Println("\t", c)
		}
	}
}
