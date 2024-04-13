package smartapigo

import "net/http"

type HistoricalParams struct {
	Exchange    string `json:"exchange"`
	SymbolToken string `json:"symboltoken"`
	Interval    string `json:"interval"`
	FromDate    string `json:"fromdate"`
	ToDate      string `json:"todate"`
}

type HistoricalResponse struct {
}

func (c *Client) GetHistoricalData(historicalParams HistoricalParams) (HistoricalResponse, error) {
	var data interface{}
	params := structToMap(data, "json")
	err := c.doEnvelope(http.MethodPost, URILTP, params, nil, &data, true)
	return data.(HistoricalResponse), err
}
