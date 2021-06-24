package main

import (
	"fmt"
	"os"

	"github.com/rgerardi/codebar-2021/crypto" // HL
)

func main() {
	symbol := "btc"

	c, err := crypto.GetCrypto(symbol) // HL
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(c)
}
