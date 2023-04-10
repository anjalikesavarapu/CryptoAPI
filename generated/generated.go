package generated


type Response struct {
	Data CryptoData `json:"data"`
}

type CryptoData struct {
	Id string `json: "id"`
    Rank string `json: "rank"`
    Symbol string `json: "symbol"`
    Name string `json: "name"`
    Supply string `json: "supply"`
    MaxSupply string `json: "maxSupply"`
    MarketCapUsd string `json: "marketCapUsd"`
    VolumeUsd24Hr string `json:"volumeUsd24Hr"`
    PriceUsd string `json: "priceUsd"`
    ChangePercent24Hr string `json:"changePercent24Hr"`
    Vwap24Hr string `json:"vwap24Hr"`
}

type MarketResponse struct {
    Data []MarketData `json:"data"`
}

type MarketData struct {
    ExchangeId string `json: "exchangeId"`
    BaseId string `json:"baseId"`
    QuoteId string `json:"quoteId"`
    QuoteSymbol string `json: "quoteSymbol"`
    VolumeUsd24Hr string `json:"volumeUsd24Hr"`
    PriceUsd string `json: "priceUsd"`
    VolumePercent string `json: "volumePercent"`
}