package smartapigo

import "net/http"

type MarketLiveParams struct {
	Mode           string   `json:"mode"`
	ExchangeTokens []string `json:"exchangeTokens"`
}

type MarketLiveResponseData struct {
	Fetched []MarketLiveFetched `json:"fetched"`
}
type MarketLiveFetched struct {
	TradingSymbol string `json:"tradingSymbol"`
	SymbolToken   string `json:"symbolToken"`
	Ltp           string `json:"ltp"`
	TradeTime     string `json:"exchTradeTime"`
	NetChange     string `json:"netChange"`
	TradeVolume   string `json:"tradeVolume"`
	LC            string `json:"lowerCircuit"`
	UC            string `json:"upperCircuit"`
	BuyQuan       string `json:"totBuyQuan"`
	SellQuan      string `json:"totSellQuan"`
	Depth         Depth  `json:"depth"`
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
