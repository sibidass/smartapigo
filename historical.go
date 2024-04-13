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
	Data [][]string `json:"data"`
}

func (c *Client) GetHistoricalData(historicalParams HistoricalParams) ([][]string, error) {
	var data [][]string
	params := structToMap(historicalParams, "json")
	err := c.doEnvelope(http.MethodPost, URIHistorical, params, nil, &data, true)
	return data, err
}
