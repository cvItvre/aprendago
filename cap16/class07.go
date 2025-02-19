package cap16

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func class07() {

	passwd := "28/07/1998"

	hashed_passwd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(passwd, string(hashed_passwd))

	err = bcrypt.CompareHashAndPassword(hashed_passwd, []byte("passwd"))
	if err != nil {
		fmt.Println(err)
		fmt.Println("Senha errada!")
	} else {
		fmt.Println("Senha correta!")
	}

}
