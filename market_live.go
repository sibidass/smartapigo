package smartapigo

import "net/http"

type MarketLiveParams struct {
	Mode           string              `json:"mode"`
	ExchangeTokens map[string][]string `json:"exchangeTokens"`
}

type MarketLiveResponseData struct {
	Fetched []MarketLiveFetched `json:"fetched"`
}
type MarketLiveFetched struct {
	TradingSymbol string  `json:"tradingSymbol"`
	SymbolToken   string  `json:"symbolToken"`
	Ltp           float64 `json:"ltp"`
	TradeTime     string  `json:"exchTradeTime"`
	NetChange     float64 `json:"netChange"`
	TradeVolume   float64 `json:"tradeVolume"`
	LC            float64 `json:"lowerCircuit"`
	UC            float64 `json:"upperCircuit"`
	BuyQuan       float64 `json:"totBuyQuan"`
	SellQuan      float64 `json:"totSellQuan"`
	Depth         Depth   `json:"depth"`
	Open          float64 `json:"open"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Close         float64 `json:"close"`
}

type Depth struct {
	Buy  []StockOrder `json:"buy"`
	Sell []StockOrder `json:"sell"`
}

type StockOrder struct {
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Orders   int     `json:"orders"`
}

func (c *Client) LiveData(ldParams MarketLiveParams) (MarketLiveResponseData, error) {
	var ld MarketLiveResponseData
	params := structToMap(ldParams, "json")
	err := c.doEnvelope(http.MethodPost, URIMarketLive, params, nil, &ld, true)
	return ld, err
}
