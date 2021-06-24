package main

import (
	"fmt"
	"os"

	"github.com/rgerardi/codebar-2021/crypto"
)

func main() {
	symbols := []string{"btc", "eth", "ltc", "dot", "pol", "bch"}

	for _, symbol := range symbols { // HL
		c, err := crypto.GetCrypto(symbol)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(c)
	}
}
