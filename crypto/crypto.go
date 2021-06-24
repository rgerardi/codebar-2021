package crypto // HL

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Crypto struct {
	Data struct {
		Symbol     string `json:"symbol"`
		Name       string `json:"name"`
		MarketData struct {
			PriceUSD float64 `json:"price_usd"` // HL
		} `json:"market_data"`
	} `json:"data"`
}

func (c *Crypto) String() string {
	return fmt.Sprintf("%s (%s): $ %.2f", c.Data.Name, c.Data.Symbol, c.Data.MarketData.PriceUSD)
}

func GetCrypto(symbol string) (*Crypto, error) {
	url := fmt.Sprintf("https://data.messari.io/api/v1/assets/%s/metrics/market-data", symbol)

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Cannot get crypto: %s", err)
	}

	defer res.Body.Close()

	c := &Crypto{}

	err = json.NewDecoder(res.Body).Decode(c) // HL

	return c, err
}
