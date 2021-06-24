package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	symbol := "btc"
	url := fmt.Sprintf("https://data.messari.io/api/v1/assets/%s/metrics/market-data", symbol)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
