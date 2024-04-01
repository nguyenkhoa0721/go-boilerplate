package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-boilerplate/pkg/logger"
	"io"
	"math"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func GetPriceFromBinance(symbol string) (big.Float, error) {
	const reqUrl = "https://api.binance.com/api/v3/avgPrice?symbol="
	var result map[string]interface{}

	resp, err := http.Get(reqUrl + strings.ToUpper(symbol) + "USDT")
	if err != nil {
		return big.Float{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return big.Float{}, fmt.Errorf("error: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return big.Float{}, err
	}

	s := new(big.Float)
	if s, success := s.SetString(result["price"].(string)); success == true {
		return *s, nil
	} else {
		return big.Float{}, errors.New("can not parse float")
	}
}

func GetPriceChangePercent24hFromBinance(symbol string) (string, error) {
	const reqUrl = "https://api.binance.com/api/v3/ticker/24hr?symbol="
	var result map[string]interface{}

	resp, err := http.Get(reqUrl + strings.ToUpper(symbol) + "USDT")
	if err != nil {
		return "0", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "0", fmt.Errorf("error: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "0", err
	}

	priceChange := result["priceChangePercent"].(string)
	if priceChange == "" {
		return "0", nil
	}

	priceChangePercent, err := strconv.ParseFloat(priceChange, 64)
	if err != nil {
		return "0", err
	}

	// divide by 100 to get percent
	priceChangePercent = priceChangePercent / 100

	// round to 4 decimal places
	priceChangePercent = math.Ceil(priceChangePercent*10000) / 10000

	return fmt.Sprintf("%.4f", priceChangePercent), nil
}

func GetPriceFromCex(symbol string) (float64, error) {
	var result map[string]interface{}

	resp, err := http.Get(fmt.Sprintf("https://api.cryptowat.ctx./markets/cexio/%susd/price", strings.ToLower(symbol)))
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("error: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	return result["result"].(map[string]interface{})["price"].(float64), nil
}

type CoinMarketCapPriceConversionResponse struct {
	Data []struct {
		Symbol      string  `json:"symbol"`
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		Amount      float64 `json:"amount"`
		LastUpdated string  `json:"last_updated"`
		Quote       struct {
			USD struct {
				Price       float64 `json:"price"`
				LastUpdated string  `json:"last_updated"`
			} `json:"USD"`
		} `json:"quote"`
	} `json:"data"`
	Status struct {
		Timestamp    string `json:"timestamp"`
		ErrorCode    int    `json:"error_code"`
		ErrorMessage string `json:"error_message"`
		Elapsed      int    `json:"elapsed"`
		CreditCount  int    `json:"credit_count"`
		Notice       string `json:"notice"`
	} `json:"status"`
}

func GetPriceFromCoinMarketCap(symbol string, apiKey string) (big.Float, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v2/tools/price-conversion", nil)
	if err != nil {
		return big.Float{}, err
	}

	q := url.Values{}
	q.Add("amount", "1")
	q.Add("symbol", strings.ToUpper(symbol))
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return big.Float{}, err
	}
	defer resp.Body.Close()

	// Check the status code is 200 OK or not.
	if resp.StatusCode != http.StatusOK {
		logger.Error(context.Background()).Msgf("error: %s", resp.Status)
		return big.Float{}, errors.New("error: " + resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)

	var result CoinMarketCapPriceConversionResponse
	err = json.Unmarshal(body, &result)

	s := new(big.Float)

	return *s.SetFloat64(result.Data[0].Quote.USD.Price), nil

}
