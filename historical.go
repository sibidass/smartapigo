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
	Data [][]interface{} `json:"data"`
}

func (c *Client) GetHistoricalData(historicalParams HistoricalParams) ([][]interface{}, error) {
	var data [][]interface{}
	params := structToMap(historicalParams, "json")
	err := c.doEnvelope(http.MethodPost, URIHistorical, params, nil, &data, true)
	return data, err
}
