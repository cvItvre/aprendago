package cap09

import (
	"fmt"
)

func ex07() {
	data := [][]string{
		{"Pedro", "Araujo", "YouTube"},
		{"Karol", "Souza", "TikTok"},
		[]string{"Unknown", "Person", "None"},
	}

	for _, v := range data {
		for _, j := range v {
			fmt.Printf("%v\t", j)
		}
		fmt.Println("")
	}
}
