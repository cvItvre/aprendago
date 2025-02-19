package cap08

import (
	"fmt"
)

func class10() {

	countryCode := map[int]string{
		55: "Brasil",
		1:  "USA",
		81: "Pernmabuco",
		54: "China",
	}

	fmt.Println(countryCode)

	for k, v := range countryCode {
		fmt.Println(k, v)
	}

	delete(countryCode, 1)

	for k, v := range countryCode {
		fmt.Println(k, v)
	}

}
