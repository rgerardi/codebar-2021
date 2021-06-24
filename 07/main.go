package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rgerardi/codebar-2021/crypto"
)

func main() {
	symbols := []string{"btc", "eth", "ltc", "dot", "pol", "bch"}

	start := time.Now()
	for _, symbol := range symbols {
		go printCryptoPrice(symbol) // HL
	}

	fmt.Println()
	fmt.Println("Total time:", time.Since(start))
}

func printCryptoPrice(symbol string) {
	c, err := crypto.GetCrypto(symbol)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(c)
}
