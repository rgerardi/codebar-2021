package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rgerardi/codebar-2021/crypto"
)

func main() {
	http.HandleFunc("/prices", func(w http.ResponseWriter, r *http.Request) { // HL
		symbols := strings.Split(r.URL.Query().Get("symbols"), ",") // HL

		crchan := make(chan *crypto.Crypto)
		start := time.Now()
		for _, symbol := range symbols {
			go fetchCryptoPrice(symbol, crchan)
		}
		for range symbols {
			printCryptoPrice(w, <-crchan) // HL
		}

		fmt.Println()
		fmt.Println("Total time:", time.Since(start))
	})

	fmt.Println("Starting web server on port :4000")
	if err := http.ListenAndServe(":4000", nil); err != nil { // HL
		fmt.Println("Cannot start server:", err)
		os.Exit(1)
	}
}

func fetchCryptoPrice(symbol string, crchan chan<- *crypto.Crypto) {
	c, err := crypto.GetCrypto(symbol)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	crchan <- c
}

func printCryptoPrice(out io.Writer, c *crypto.Crypto) {
	fmt.Fprintln(out, c)
}
