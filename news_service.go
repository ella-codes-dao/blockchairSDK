package blockchairSDK

import "net/url"

type NewsOptions struct {
	Title       string
	Timestamp   string
	Source      string
	Description string
	Language    string
	Limit       string
	Offset      string
}

type FormattedNewsResponse struct {
	Data    []NewsData `json:"data"`
	Context Context    `json:"context"`
}

type NewsData struct {
	Title         string      `json:"title"`
	Source        string      `json:"source"`
	Language      string      `json:"language"`
	Link          string      `json:"link"`
	LinkAmp       interface{} `json:"link_amp"`
	LinkIframable bool        `json:"link_iframable"`
	Time          string      `json:"time"`
	Tags          string      `json:"tags"`
	Description   string      `json:"description"`
	Hash          string      `json:"hash"`
	File          string      `json:"file"`
	Permalink     string      `json:"permalink"`
}

func (c *Client) News(opts NewsOptions) (*FormattedNewsResponse, error) {
	path := "/news"
	requestParams := url.Values{}

	var queryString string

	if opts.Language != "" {
		if queryString == "" {
			queryString = "language(" + opts.Language + ")"
		} else {
			queryString += "," + "language(" + opts.Language + ")"
		}
	}

	if opts.Timestamp != "" {
		if queryString == "" {
			queryString = "time(" + opts.Timestamp + ")"
		} else {
			queryString += "," + "time(" + opts.Timestamp + ")"
		}
	}

	if queryString != "" {
		requestParams.Add("q", queryString)
	}

	if opts.Limit != "" {
		requestParams.Add("limit", opts.Limit)
	}

	if opts.Offset != "" {
		requestParams.Add("offset", opts.Offset)
	}

	req, err := c.newRequest(path, requestParams)
	if err != nil {
		return nil, err
	}

	var news FormattedNewsResponse
	_, err = c.do(req, &news)
	if err != nil {
		return nil, err
	}

	return &news, nil
}
