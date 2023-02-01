package blockchairSDK

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type DashboardService struct {
	chain  string
	client *Client
}

type AddressOptions struct {
	Limit               int
	Offset              int
	ERC20               string
	Nonce               bool
	TypeOutputOnly      bool
	AssetsInUSD         bool
	LatestStateOnly     bool
	ShowContractDetails bool
}

type FormattedAddressResponse struct {
	Data    ChainAddress `json:"data"`
	Context Context      `json:"context"`
}
type AddressResponse struct {
	Data    *json.RawMessage `json:"data"`
	Context Context          `json:"context"`
}
type ContractDetails struct {
	Type                    interface{} `json:"type"`
	Data                    interface{} `json:"data"`
	CreatingTransactionHash string      `json:"creating_transaction_hash"`
	CreatingAddress         string      `json:"creating_address"`
	CreatingTime            string      `json:"creating_time"`
}
type Address struct {
	Type                string          `json:"type"`
	ContractCodeHex     string          `json:"contract_code_hex"`
	ContractCreated     bool            `json:"contract_created"`
	ContractDestroyed   bool            `json:"contract_destroyed"`
	Balance             string          `json:"balance"`
	BalanceUsd          int             `json:"balance_usd"`
	ReceivedApproximate string          `json:"received_approximate"`
	ReceivedUsd         int             `json:"received_usd"`
	SpentApproximate    string          `json:"spent_approximate"`
	SpentUsd            int             `json:"spent_usd"`
	FeesApproximate     string          `json:"fees_approximate"`
	FeesUsd             int             `json:"fees_usd"`
	ReceivingCallCount  int             `json:"receiving_call_count"`
	SpendingCallCount   int             `json:"spending_call_count"`
	CallCount           int             `json:"call_count"`
	TransactionCount    int             `json:"transaction_count"`
	FirstSeenReceiving  string          `json:"first_seen_receiving"`
	LastSeenReceiving   string          `json:"last_seen_receiving"`
	FirstSeenSpending   string          `json:"first_seen_spending"`
	LastSeenSpending    string          `json:"last_seen_spending"`
	Nonce               interface{}     `json:"nonce"`
	ContractDetails     ContractDetails `json:"contract_details"`
	AssetBalanceUsd     int             `json:"asset_balance_usd"`
}
type Calls struct {
	BlockID         int         `json:"block_id"`
	TransactionHash string      `json:"transaction_hash"`
	Index           string      `json:"index"`
	Time            string      `json:"time"`
	Sender          string      `json:"sender"`
	Recipient       string      `json:"recipient"`
	Value           int         `json:"value"`
	ValueUsd        interface{} `json:"value_usd"`
	Transferred     bool        `json:"transferred"`
}
type Erc20 struct {
	TokenAddress       string      `json:"token_address"`
	TokenName          string      `json:"token_name"`
	TokenSymbol        string      `json:"token_symbol"`
	TokenDecimals      int         `json:"token_decimals"`
	BalanceApproximate float64     `json:"balance_approximate"`
	Balance            string      `json:"balance"`
	BalanceUsd         interface{} `json:"balance_usd"`
}
type Layer2 struct {
	Erc20 []Erc20 `json:"erc_20"`
}
type ChainAddress struct {
	Address Address `json:"address"`
	Calls   []Calls `json:"calls"`
	Layer2  Layer2  `json:"layer_2"`
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
	Code           int     `json:"code"`
	Source         string  `json:"source"`
	Limit          string  `json:"limit"`
	Offset         string  `json:"offset"`
	Results        int     `json:"results"`
	State          int     `json:"state"`
	StateLayer2    int     `json:"state_layer_2"`
	MarketPriceUsd float64 `json:"market_price_usd"`
	Cache          Cache   `json:"cache"`
	API            API     `json:"api"`
	Servers        string  `json:"servers"`
	Time           float64 `json:"time"`
	RenderTime     float64 `json:"render_time"`
	FullTime       float64 `json:"full_time"`
	RequestCost    float64 `json:"request_cost"`
}

func (cs *ChainService) Dashboard() *DashboardService {
	return &DashboardService{
		chain:  cs.chain,
		client: cs.client,
	}
}

func (ds *DashboardService) Address(addr string, opts AddressOptions) (*FormattedAddressResponse, error) {
	path := ds.chain + "/dashboards/address/" + addr
	requestParams := url.Values{}

	if opts.ERC20 != "" {
		requestParams.Add("erc_20", opts.ERC20)
	}

	if opts.AssetsInUSD {
		requestParams.Add("assets_in_usd", "true")
	}

	if opts.LatestStateOnly {
		requestParams.Add("state", "latest")
	}

	if opts.Limit > 0 {
		requestParams.Add("limit", fmt.Sprint(opts.Limit))
	}

	if opts.Offset > 0 {
		requestParams.Add("offset", fmt.Sprint(opts.Offset))
	}

	if opts.Nonce {
		requestParams.Add("nonce", "true")
	}

	if opts.ShowContractDetails {
		requestParams.Add("contract_details", "true")
	}

	if opts.TypeOutputOnly {
		requestParams.Add("output", "type")
	}

	req, err := ds.client.newRequest(path, requestParams)
	if err != nil {
		return nil, err
	}

	var address AddressResponse
	_, err = ds.client.do(req, &address)
	if err != nil {
		return nil, err
	}

	var dataMap map[string]interface{}
	json.Unmarshal(*address.Data, &dataMap)

	addressJSON, err := json.Marshal(dataMap[addr])
	if err != nil {
		return nil, err
	}

	var addressStruct ChainAddress
	json.Unmarshal(addressJSON, &addressStruct)

	formattedResponse := &FormattedAddressResponse{
		Data:    addressStruct,
		Context: address.Context,
	}

	return formattedResponse, err
}
