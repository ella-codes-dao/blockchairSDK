package blochchairSDK

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Client struct {
	baseURL *url.URL
	apiKey  string

	httpClient *http.Client
}

func NewClient(httpClient *http.Client, apiKey string) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	//Create Base URL
	baseURL, err := url.Parse("https://api.blockchair.com")
	if err != nil {
		return nil, err
	}

	client := &Client{
		baseURL:    baseURL,
		apiKey:     apiKey,
		httpClient: httpClient,
	}

	return client, nil
}

func (c *Client) newRequest(path string, query url.Values) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	requestParams := url.Values{}

	if query != nil {
		requestParams = query
	}

	if c.apiKey != "" {
		requestParams.Add("key", c.apiKey)
	}

	req.URL.RawQuery = requestParams.Encode()

	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
