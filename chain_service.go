package blockchairSDK

type ChainService struct {
	chain  string
	client *Client
}

func (c *Client) Ethereum() *ChainService {
	cs := &ChainService{
		chain:  "/ethereum",
		client: c,
	}

	return cs
}

func (c *Client) EthereumTestnet() *ChainService {
	cs := &ChainService{
		chain:  "/ethereum/testnet",
		client: c,
	}

	return cs
}
