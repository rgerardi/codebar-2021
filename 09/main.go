package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rgerardi/codebar-2021/crypto"
)

func main() {
	symbols := []string{"btc", "eth", "ltc", "dot", "pol", "bch"}

	crchan := make(chan *crypto.Crypto)

	start := time.Now()
	for _, symbol := range symbols {
		go fetchCryptoPrice(symbol, crchan)
	}

	for range symbols {
		printCryptoPrice(os.Stdout, <-crchan) // HL
	}

	fmt.Println()
	fmt.Println("Total time:", time.Since(start))
}

func fetchCryptoPrice(symbol string, crchan chan<- *crypto.Crypto) {
	c, err := crypto.GetCrypto(symbol)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	crchan <- c
}

func printCryptoPrice(out io.Writer, c *crypto.Crypto) { // HL
	fmt.Fprintln(out, c) // HL
}
