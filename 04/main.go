package main

import (
	"fmt"
	"os"

	"github.com/rgerardi/codebar-2021/crypto"
)

func main() {
	//	symbols := os.Args[1:] // For command line tools
	symbols := []string{"btc"} // To execute on slides

	c, err := crypto.GetCrypto(symbols[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(c)
}
