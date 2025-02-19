package cap17

import (
	"encoding/json"
	"fmt"
)

type usr struct {
	First string
	Age   int
}

func ex01() {
	u1 := usr{
		First: "James",
		Age:   32,
	}

	u2 := usr{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := usr{
		First: "M",
		Age:   54,
	}

	users := []usr{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	jsonData, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonData))

}
