package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(addr, nil))
}

type Ticker struct {
	ID               string      `json:"id"`
	Name             string      `json:"name"`
	Symbol           string      `json:"symbol"`
	Rank             string      `json:"rank"`
	PriceUsd         string      `json:"price_usd"`
	PriceBtc         string      `json:"price_btc"`
	Two4HVolumeUsd   string      `json:"24h_volume_usd"`
	MarketCapUsd     string      `json:"market_cap_usd"`
	AvailableSupply  string      `json:"available_supply"`
	TotalSupply      string      `json:"total_supply"`
	MaxSupply        interface{} `json:"max_supply"`
	PercentChange1H  string      `json:"percent_change_1h"`
	PercentChange24H string      `json:"percent_change_24h"`
	PercentChange7D  string      `json:"percent_change_7d"`
	LastUpdated      string      `json:"last_updated"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	// parse form body
	if err := r.ParseForm(); err != nil {
		log.Printf("error parsing form %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// grab values from form
	team := r.Form.Get("team_domain")
	user := r.Form.Get("user_name")
	cmd := r.Form.Get("command")
	args := strings.Fields(r.Form.Get("text"))

	if len(args) == 0 {
		fmt.Fprintln(w, "token is required")
		return
	}

	// log stuff
	fmt.Printf("%s %v from user %s (%s)\n", cmd, args, user, team)

	// perform the request
	url := "https://api.coinmarketcap.com/v1/ticker/" + args[0] + "/"
	res, err := http.Get(url)
	if err != nil {
		log.Printf("token required")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	tickers := []Ticker{}
	decodeErr := json.NewDecoder(res.Body).Decode(&tickers)

	if decodeErr != nil {
		log.Printf("can't parse API response")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// TODO: serialize into struct
	fmt.Fprintf(w, "%v", tickers[0].PriceUsd)
}
