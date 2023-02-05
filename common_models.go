package blockchairSDK

import "encoding/json"

type GeneralResponse struct {
	Data    *json.RawMessage `json:"data"`
	Context Context          `json:"context"`
}

type Cache struct {
	Live     bool        `json:"live"`
	Duration int         `json:"duration"`
	Since    string      `json:"since"`
	Until    string      `json:"until"`
	Time     interface{} `json:"time"`
}

type API struct {
	Version         string      `json:"version"`
	LastMajorUpdate string      `json:"last_major_update"`
	NextMajorUpdate interface{} `json:"next_major_update"`
	Documentation   string      `json:"documentation"`
	Notice          string      `json:"notice"`
}

type Context struct {
	Code        int         `json:"code"`
	Source      string      `json:"source"`
	Limit       int         `json:"limit"`
	Offset      int         `json:"offset"`
	Rows        int         `json:"rows"`
	PreRows     int         `json:"pre_rows"`
	TotalRows   interface{} `json:"total_rows"`
	Cache       Cache       `json:"cache"`
	API         API         `json:"api"`
	Servers     string      `json:"servers"`
	Time        float64     `json:"time"`
	RenderTime  float64     `json:"render_time"`
	FullTime    float64     `json:"full_time"`
	RequestCost float64     `json:"request_cost"`
}
