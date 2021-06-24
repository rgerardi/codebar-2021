package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rgerardi/codebar-2021/crypto"
)

func main() {
	symbols := []string{"btc", "eth", "ltc", "dot", "pol", "bch"}

	crchan := make(chan *crypto.Crypto) // HL

	start := time.Now()
	for _, symbol := range symbols {
		go fetchCryptoPrice(symbol, crchan) // HL
	}

	for range symbols {
		printCryptoPrice(<-crchan) // HL
	}

	fmt.Println()
	fmt.Println("Total time:", time.Since(start))
}

func fetchCryptoPrice(symbol string, crchan chan<- *crypto.Crypto) { // HL
	c, err := crypto.GetCrypto(symbol)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	crchan <- c // HL
}

func printCryptoPrice(c *crypto.Crypto) {
	fmt.Println(c)
}
