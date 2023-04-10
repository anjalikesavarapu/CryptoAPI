/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"
	"github.com/spf13/cobra"
	"io"
	"encoding/json"
	"strings"
	generated "github.com/anjalikesavarapu/CryptoAPI/generated"
)

// priceCmd represents the price command
var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("price called")
		//fmt.Println(args)
		currency := strings.ToLower(args[0])
		//getCurrencyprice(currency)
		showMarketData, _ := cmd.Flags().GetBool("markets")
		if showMarketData {
			getMarkets(currency)
		} else {
			getCurrencyprice(currency)
		}
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// priceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// priceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getCurrencyprice(currency string) {
	fmt.Println("Getting price for ", currency)
	coincapAPIurl := "https://api.coincap.io/v2/assets/" + currency
	client := http.Client{}

	req,err := http.NewRequest("GET", coincapAPIurl, nil)
	if err != nil {
		fmt.Println("Error creating request")
	}

	res,err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request")
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in reading response")
	}

	var data generated.Response
	json.Unmarshal(resBody, &data)

	if data.Data.Id == "" {
		fmt.Println("currency not found")
	} else {
		fmt.Printf("Currency: %s\n Symbol: %s\n Price: $%s\n Market Cap: $%s\n Volume: %s\n Change: %s%%\n Rank: %s\n Supply: %s\n Max Supply: %s\n Vwap: %s\n", data.Data.Name, data.Data.Symbol, data.Data.PriceUsd, data.Data.MarketCapUsd, data.Data.VolumeUsd24Hr, data.Data.ChangePercent24Hr, data.Data.Rank, data.Data.Supply, data.Data.MaxSupply, data.Data.Vwap24Hr)
	}

}

func getMarkets(currency string) {

	fmt.Println("Getting markets for the currency : ", currency)
	coincapAPIurl := "https://api.coincap.io/v2/assets/ " + currency + "/markets?limit=10"
	client := http.Client{}

	req, err := http.NewRequest("GET", coincapAPIurl, nil)
	if err != nil {
		fmt.Println("Error creating request")
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request")
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading request")
	}

	var data generated.MarketResponse
	json.Unmarshal(resBody, &data)
	for _, market := range data.Data {
        fmt.Printf("\n\nExchange: %s\nBase: %s\nQuote: %s\nPrice: $%s\nVolume: %s\nVolume Percent: %s\n", market.ExchangeId, market.BaseId, market.QuoteSymbol, market.PriceUsd, market.VolumeUsd24Hr, market.VolumePercent)
    }

}
